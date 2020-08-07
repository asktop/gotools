package aswagger

import (
    "errors"
    "fmt"
    "github.com/asktop/gotools/afile"
    "github.com/asktop/gotools/amap"
    "github.com/asktop/gotools/aswagger/scan"
    "github.com/asktop/gotools/atime"
    "github.com/astaxie/beego"
    "github.com/go-openapi/spec"
    "go/ast"
    "os"
    "path"
    "path/filepath"
    "strings"
)

type Apidoc struct {
    Appname string //项目包名 示例：asktop
    Apppath string //项目根目录（绝路路径） 示例：D:\gospace\src\asktop
    Docpath string //文档存放目录（绝路路径） 示例：D:\gospace\src\asktop\upload\doc

    Title     string //文档标题
    Host      string //正式域名 示例：asktop.wang
    LocalHost string //测试域名 示例：127.0.0.1:8881
    BasePath  string //接口公共路由 示例：/api/v1
}

//Spec指定项目
//onlyMakeApiDoc 是否只生成接口文档；默认即生成接口文档，又更新路由
func Spec(apidoc Apidoc, onlyMakeApiDoc ...bool) (err error, msgs []string) {
    if apidoc.Appname == "" {
        err = errors.New("Appname不能为空")
        return
    }
    if apidoc.Apppath == "" {
        err = errors.New("Apppath不能为空")
        return
    }
    if apidoc.Docpath == "" {
        err = errors.New("Docpath不能为空")
        return
    }

    var swg = spec.Swagger{
        SwaggerProps: spec.SwaggerProps{
            Swagger: "2.0",
            Info: &spec.Info{
                InfoProps: spec.InfoProps{
                    Title:       apidoc.Title,                             //文档标题
                    Version:     atime.FormatDateTime(atime.Now().Unix()), //文档版本
                    Description: "",                                       //文档描述
                },
                VendorExtensible: spec.VendorExtensible{Extensions: map[string]interface{}{"x-framework": "go-swagger"}},
            },
            Host:     apidoc.Host,
            BasePath: apidoc.BasePath,
            Consumes: []string{"application/x-www-form-urlencoded", "multipart/form-data", "application/json"},
            Produces: []string{"application/json"},
            Schemes:  []string{"http", "https"},
            Paths: &spec.Paths{
                Paths:            map[string]spec.PathItem{},
                VendorExtensible: spec.VendorExtensible{Extensions: map[string]interface{}{"x-framework": "go-swagger"}},
            },
            Responses:   map[string]spec.Response{},
            Definitions: map[string]spec.Schema{},
        },
    }

    //删除 api 路由的 import 文件
    deleteApiRouterImport(apidoc.Apppath)
    //扫描所有文件并获取接口文件路径
    modules, filePaths, err := scanAllFileImport(apidoc.Apppath)
    //fmt.Println(filePaths)
    if err != nil {
        return
    }
    //生成 api 路由的 import 文件
    makeApiRouterImport(apidoc.Apppath, apidoc.Appname, modules)
    //扫描项目main调用引用的所有包和文件
    scaner, err := scan.NewAppScanner(&scan.Opts{
        BasePath: apidoc.Appname,
    })
    if err != nil {
        err = errors.New("NewAppScanner err:" + err.Error())
        return
    }
    //扫描所有文件并解析
    modules, msgs = scanAllFile(filePaths, &swg, scaner)
    //删除 api 路由的 import 文件
    deleteApiRouterImport(apidoc.Apppath)

    if len(onlyMakeApiDoc) > 0 && onlyMakeApiDoc[0] {
        //生成接口文档
        err = makeApiDoc(&swg, apidoc.Docpath, apidoc.LocalHost)
    } else {
        //生成api router文件
        makeApiRouter(apidoc.Apppath, apidoc.Appname, apidoc.BasePath, modules)
        //生成接口文档
        err = makeApiDoc(&swg, apidoc.Docpath, apidoc.LocalHost)
    }
    return
}

//扫描所有文件并解析
func scanAllFile(filePaths []string, swg *spec.Swagger, scaner *scan.AppScanner) (modules map[string]map[string]map[string]string, msgs []string) {
    modules = map[string]map[string]map[string]string{}
    for _, filePath := range filePaths {
        //获取接口路由等参数
        apiPath := strings.Replace(filePath, `\`, `/`, -1)
        apiPath = strings.Replace(apiPath, "IO.go", "", -1)
        apiPaths := strings.Split(apiPath, `/`)
        apiPaths = apiPaths[len(apiPaths)-3:]
        apiPath = "/" + strings.Join(apiPaths, "/") //接口路由
        moduleName := apiPaths[0]                   //模块名
        serviceName := apiPaths[1]                  //服务名
        apiName := apiPaths[2]                      //方法名

        //扫描匹配文件并获取文件
        //fmt.Println("寻找：", filePath)
        file := getFile(scaner, filePath)
        if file == nil {
            msgs = append(msgs, "未找到文件："+filePath)
            beego.Info("未找到文件：", filePath)
            continue
        }
        //解析文件
        err := scaner.ParseSchema(file)
        if err != nil {
            msgs = append(msgs, "解析文件出错："+filePath+" err:"+err.Error())
            beego.Info("解析文件出错：", filePath, " err:", err)
            continue
        }

        //创建API对象
        item := spec.PathItem{}
        item.Post = spec.NewOperation("")

        //解析 Interface 结构
        if _, find := scaner.Definitions[apiName+"Interface"]; !find {
            msgs = append(msgs, "未找到 "+filePath+" 的 "+apiName+"Interface 结构")
            beego.Info("未找到 " + filePath + " 的 " + apiName + "Interface 结构")
            continue
        }
        itf := scaner.Definitions[apiName+"Interface"]
        item.Post.Summary = itf.Description
        item.Post.Description = `<a href="#top_title">返回模块列表</a><br>`
        //模块标签
        item.Post.Tags = []string{moduleName}
        var hasTag bool
        for _, tag := range swg.Tags {
            if tag.Name == moduleName {
                hasTag = true
                break
            }
        }
        if !hasTag {
            tag := spec.Tag{}
            tag.Name = moduleName
            tag.Description = ` <a href="#top_title">返回模块列表</a>`
            swg.Tags = append(swg.Tags, tag)
        }

        //解析 Req 结构
        if _, find := scaner.Definitions[apiName+"Req"]; !find {
            msgs = append(msgs, "未找到 "+filePath+" 的 "+apiName+"Req 结构")
            beego.Info("未找到 " + filePath + " 的 " + apiName + "Req 结构")
            //continue
        } else {
            //req 请求内容解析
            req := scaner.Definitions[apiName+"Req"]
            ordermap := amap.NewOrderMapFrom(req.SchemaProps.Properties)
            ordermap.SortKey()
            for _, name := range ordermap.Keys() {
                prop := ordermap.Get(name).(spec.Schema)
                param := spec.Parameter{}
                param.Name = name
                pjson, _ := prop.MarshalJSON()
                param.UnmarshalJSON(pjson)
                //fmt.Println(string(pjson))
                //fmt.Println(param.Ref.String())
                //字段类型
                if strings.Contains(param.Description, "[header]") {
                    param.Description = strings.Replace(param.Description, "[header]", "", -1)
                    param.In = "header"
                } else if strings.Contains(param.Description, "[query]") {
                    param.Description = strings.Replace(param.Description, "[query]", "", -1)
                    param.In = "query"
                } else {
                    param.In = "formData"
                }
                //是否必须
                if strings.Contains(param.Description, "[must]") {
                    param.Description = strings.Replace(param.Description, "[must]", "", -1)
                    param.Required = true
                }
                if param.Ref.String() != "" {
                    refNames := strings.Split(param.Ref.String(), "/")
                    refName := refNames[len(refNames)-1]
                    param.Type = "string"
                    param.Format = refName
                    param.Ref = spec.Ref{}
                }
                if param.Items != nil && param.Items.Ref.String() != "" {
                    refNames := strings.Split(param.Items.Ref.String(), "/")
                    refName := refNames[len(refNames)-1]
                    param.Type = "array"
                    param.Format = refName
                    param.Items.Ref = spec.Ref{}
                }

                item.Post.AddParam(&param)
            }
        }

        //解析 Res 结构
        if _, find := scaner.Definitions[apiName+"Res"]; !find {
            msgs = append(msgs, "未找到 "+filePath+" 的 "+apiName+"Res 结构")
            beego.Info("未找到 " + filePath + " 的 " + apiName + "Res 结构")
            continue
        } else {
            //res 调用结构解析
            res := scaner.Definitions[apiName+"Res"]
            for _, prop := range res.SchemaProps.Properties {
                if prop.Items != nil && prop.Items.Schema.Ref.String() != "" {
                    refNames := strings.Split(prop.Items.Schema.Ref.String(), "/")
                    refName := refNames[len(refNames)-1]
                    swg.Definitions[refName] = scaner.Definitions[refName]
                }
            }
            //res 响应内容解析
            res200 := spec.NewResponse()
            res200.Description = "请求成功"

            //返回值：请求成功时返回的数据
            resJson, _ := scaner.Definitions[apiName+"Res"].MarshalJSON()
            resData := &spec.Schema{}
            resData.UnmarshalJSON(resJson)

            //方式1：直接返回
            //res200.Schema = resData

            //方式2：封装返回
            //返回值：状态
            status := spec.Schema{}
            status.Type = []string{"boolean"}
            status.Format = "bool"
            status.Description = "状态[true:成功, false:失败]"
            //返回值：提示信息
            msg := spec.Schema{}
            msg.Type = []string{"string"}
            msg.Format = "string"
            msg.Description = "提示信息：默认''，请求失败时为失败提示"
            //返回：请求成功
            schema200 := &spec.Schema{}
            schema200.Type = []string{"object"}
            schema200.Properties = map[string]spec.Schema{}
            schema200.Properties["status"] = status
            schema200.Properties["msg"] = msg
            resData.Description = "成功返回数据"
            schema200.Properties["data"] = *resData
            res200.Schema = schema200

            item.Post.RespondsWith(200, res200)
        }

        //添加接口
        swg.Paths.Paths[strings.ToLower(apiPath)] = item

        //添加路由
        services, ok := modules[moduleName]
        if !ok {
            services = map[string]map[string]string{}
        }
        methods, ok := services[serviceName]
        if !ok {
            methods = map[string]string{}
        }
        methods[apiName] = itf.Description
        services[serviceName] = methods
        modules[moduleName] = services
    }
    return
}

//获取单个文件
func getFile(scaner *scan.AppScanner, filepath string) *ast.File {
    for _, packageInfo := range scaner.Prog.AllPackages {
        for _, fileInfo := range packageInfo.Files {
            file := scaner.Prog.Fset.File(fileInfo.Pos())
            if file.Name() == filepath {
                return fileInfo
            }
        }
    }
    return nil
}

//扫描所有文件并获取接口文件路径
func scanAllFileImport(apppath string) (modules map[string]map[string]map[string]string, filePaths []string, err error) {
    modules = map[string]map[string]map[string]string{}
    apppath = filepath.Join(apppath, "service")
    //fmt.Println(apppath)
    //按名称顺序依次扫描 servicePath 下的文件和文件夹
    err = filepath.Walk(apppath, func(filePath string, fInfo os.FileInfo, err error) error {
        //fmt.Println(filePath)
        if fInfo.IsDir() {
            return nil
        }
        //筛选以 IO.go 结尾的文件
        //if ok, _ := path.Match("*IO.go", filePath); ok {
        if strings.HasSuffix(filePath, "IO.go") {
            //fmt.Println("ok:",filePath)
            //获取接口路由等参数
            apiPath := strings.Replace(filePath, `\`, `/`, -1)
            apiPath = strings.Replace(apiPath, "IO.go", "", -1)
            apiPaths := strings.Split(apiPath, `/`)
            apiPaths = apiPaths[len(apiPaths)-3:]
            apiPath = "/" + strings.Join(apiPaths, "/") //接口路由
            moduleName := apiPaths[0]                   //模块名
            serviceName := apiPaths[1]                  //服务名
            apiName := apiPaths[2]                      //方法名

            //添加路由
            services, ok := modules[moduleName]
            if !ok {
                services = map[string]map[string]string{}
            }
            methods, ok := services[serviceName]
            if !ok {
                methods = map[string]string{}
            }
            methods[apiName] = ""
            services[serviceName] = methods
            modules[moduleName] = services

            filePaths = append(filePaths, filePath)
            return nil
        }
        return nil
    })
    if err != nil {
        err = fmt.Errorf("filepath.Walk() err: %v\n", err)
    }
    return
}

//生成 api 路由的 import 文件
func makeApiRouterImport(apppath string, appname string, modules map[string]map[string]map[string]string) {
    apipath := filepath.Join(apppath, "routers")
    //遍历生成api路由文件
    for module, services := range modules {
        filePath := filepath.Join(apipath, "api_router_"+module+"_import.go")
        imports := []string{}
        for service, _ := range services {
            imports = append(imports, `_ "`+appname+"/service/"+module+"/"+service+`"`)
        }
        //组装api路由文件内容
        var body string
        body += "package routers\n\n"
        body += "import (\n"
        for _, imp := range imports {
            body += "\t" + imp + "\n"
        }
        body += ")\n\n"
        afile.WriteFile(filePath, body)
    }
}

//删除 api 路由的 import 文件
func deleteApiRouterImport(apppath string) {
    apipath := filepath.Join(apppath, "routers")
    //删除原api路由文件
    _, apifiles, _ := afile.GetNames(apipath)
    for _, apifile := range apifiles {
        if ok, _ := path.Match("api_router_*_import.go", apifile); ok {
            afile.Delete(filepath.Join(apipath, apifile))
        }
    }
}

//生成 api 路由文件
func makeApiRouter(apppath string, appname string, bathrouter string, modules map[string]map[string]map[string]string) {
    apipath := filepath.Join(apppath, "routers")
    //删除原api路由文件
    _, apifiles, _ := afile.GetNames(apipath)
    for _, apifile := range apifiles {
        if ok, _ := path.Match("api_router_*.go", apifile); ok {
            afile.Delete(filepath.Join(apipath, apifile))
        }
    }
    //遍历生成api路由文件
    for module, services := range modules {
        filePath := filepath.Join(apipath, "api_router_"+module+".go")
        imports := []string{}
        imports = append(imports, `"`+appname+`/controllers/api/service"`)
        apis := []string{}
        for service, methods := range services {
            if len(methods) != 0 {
                imports = append(imports, `"`+appname+"/service/"+module+"/"+service+`"`)
            }
            for method, des := range methods {
                //var _ atvert.ListsInterface = new(atvert.Atvert)
                //api.NewRouter("/api/v2/base/atvert/Lists", &atvert.Atvert{}, "Lists")
                apis = append(apis, "/*"+des+"*/")
                apis = append(apis, "var _ "+service+"."+method+"Interface = new("+service+"."+strings.ToUpper(service[0:1])+service[1:]+")")
                //apis = append(apis, `service.NewRouter("`+bathrouter+"/"+module+"/"+service+"/"+method+`", &`+service+"."+strings.ToUpper(service[0:1])+service[1:]+`{}, "`+method+`")`)
                apis = append(apis, `service.NewRouter("`+bathrouter+"/"+module+"/"+service+"/"+strings.ToLower(method)+`", &`+service+"."+strings.ToUpper(service[0:1])+service[1:]+`{}, "`+method+`")`)
            }
        }
        //组装api路由文件内容
        var body string
        body += "package routers\n\n"
        body += "import (\n"
        for _, imp := range imports {
            body += "\t" + imp + "\n"
        }
        body += ")\n\n"
        body += "func init() {\n"
        for _, api := range apis {
            body += "\t" + api + "\n"
        }
        body += "}"
        afile.WriteFile(filePath, body)
    }
}

//生成接口文档
func makeApiDoc(swg *spec.Swagger, docpath string, localHost string) error {
    //添加模块列表
    description := swg.Info.Description
    description += `<span id="top_title">【模块列表】：</span><br>`
    descGroups := [][]string{}
    descGroup := []string{}
    descGroupIndex := 0
    tdCount := 4
    for _, tag := range swg.Tags {
        if descGroupIndex >= tdCount {
            descGroups = append(descGroups, descGroup)
            descGroupIndex = 0
            descGroup = []string{}
        }
        desc := `<a href="#operations-tag-` + tag.Name + `">` + tag.Name + "</a> " + strings.Replace(tag.Description, `<a href="#top_title">返回模块列表</a>`, "", -1)
        descGroup = append(descGroup, desc)
        descGroupIndex += 1
    }
    descGroups = append(descGroups, descGroup)
    descTable := `<table width="100%" border="1" style="word-break: break-all;">`
    for _, descGroup := range descGroups {
        descTable += `<tr>`
        for i := 0; i < tdCount; i++ {
            desc := ""
            if i < len(descGroup) {
                desc = descGroup[i]
            }
            descTable += `<td width="25%" style="padding: 5px">` + desc + `</td>`
        }
        descTable += `</tr>`
    }
    descTable += `</table><br>`
    description += descTable
    swg.Info.Description = description
    //序列化 swg 生成 json 文档
    swginfo, _ := swg.MarshalJSON()
    if localHost == "" {
        return afile.WriteFile(getApidocPath(docpath), string(swginfo))
    } else {
        err := afile.WriteFile(getApidocPath(docpath), string(swginfo))
        if err != nil {
            return err
        } else {
            swg.Host = localHost
            swginfo, _ := swg.MarshalJSON()
            return afile.WriteFile(getApidocPath(docpath, "local"), string(swginfo))
        }
    }
}

//获取apidoc文件路径
func getApidocPath(docpath string, docname ...string) string {
    if len(docname) > 0 && docname[0] != "" {
        return filepath.Join(docpath, "apidoc."+docname[0]+".json")
    } else {
        return filepath.Join(docpath, "apidoc.json")
    }
}
