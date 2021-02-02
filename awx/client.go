package awx

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/asktop/gotools/aclient"
    "github.com/asktop/gotools/ajson"
    "github.com/asktop/gotools/cache"
    "regexp"
    "strings"
    "sync"
    "time"
)

type Client struct {
    mu        sync.RWMutex
    appid     string
    appsecret string

    //token
    accessTokenDriver     string //token存储位置
    accessTokenCreateTime int64  //token获取时间
    accessTokenExpireTime int64  //token有效时长
    accessToken           string //token

    //消息模板
    tmplMu sync.RWMutex
    tmpls  map[string]*MsgTemplate
}

//消息模板
type MsgTemplate struct {
    TemplateId string `json:"template_id"`
    Content    string `json:"content"`
    timestamp  int64  `json:"-"`
}

func NewClient(appid string, appsecret string, accessTokenDriver ...string) *Client {
    c := new(Client)
    c.appid = appid
    c.appsecret = appsecret

    if len(accessTokenDriver) > 0 {
        c.accessTokenDriver = accessTokenDriver[0]
    }
    c.accessTokenExpireTime = 7000

    c.tmpls = map[string]*MsgTemplate{}
    return c
}

func (c *Client) SetAccessToken(accessToken string) {
    c.accessToken = accessToken
    c.accessTokenCreateTime = time.Now().Unix()
}

//获取微信access_token redis缓存的key
func (c *Client) GetAccessTokenRedisKey() string {
    return fmt.Sprintf("wx:access_token:%s", c.appid)
}

//获取微信access_token
func (c *Client) GetAccessToken() (accessToken string, errMsg string, err error) {
    if c.accessTokenDriver == "redis" {
        accessToken, _, err = cache.NewRedis().Get(c.GetAccessTokenRedisKey()).String()
        if err != nil {
            errMsg = "查询redis微信access_token出错"
            return "", errMsg, err
        }
    } else {
        if time.Now().Unix()-c.accessTokenCreateTime <= c.accessTokenExpireTime {
            accessToken = c.accessToken
        }
    }

    if accessToken == "" {
        if c.appid == "" {
            errMsg = "请先设置微信appid"
            return "", errMsg, errors.New(errMsg)
        }
        if c.appsecret == "" {
            errMsg = "请先设置微信appsecret"
            return "", errMsg, errors.New(errMsg)
        }
        res, _, err := aclient.NewClient().Get(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", c.appid, c.appsecret), nil)
        if err != nil {
            errMsg = "请求微信access_token出错"
            return "", errMsg, err
        }
        rs, err := ajson.DecodeToMapString(res)
        if err != nil {
            errMsg = "解析微信access_token出错"
            return "", errMsg, err
        }
        if rs["errmsg"] != "" {
            errMsg = "请求微信access_token出错"
            return "", errMsg, errors.New(rs["errmsg"])
        }
        accessToken = rs["access_token"]
        if accessToken == "" {
            errMsg = "请求微信access_token出错"
            return "", errMsg, errors.New("access_token为空")
        }
        //fmt.Println("access_token:" + accessToken)

        if c.accessTokenDriver == "redis" {
            _, err = cache.NewRedis().Set(c.GetAccessTokenRedisKey(), accessToken, c.accessTokenExpireTime)
            if err != nil {
                errMsg = fmt.Sprintf("缓存微信access_token出错, access_token:%s, err:%s", accessToken, err.Error())
            }
        } else {
            c.accessToken = accessToken
            c.accessTokenCreateTime = time.Now().Unix()
        }
    }
    return accessToken, errMsg, nil
}

//获取微信消息模板
func (c *Client) GetTemplate(tmplId string) (tmpl *MsgTemplate, errMsg string, err error) {
    c.tmplMu.RLock()
    tmpl = c.tmpls[tmplId]
    c.tmplMu.RUnlock()
    if tmpl == nil || time.Now().Unix()-tmpl.timestamp > 60 {
        var accessToken string
        accessToken, errMsg, err = c.GetAccessToken()
        if err != nil {
            return nil, errMsg, err
        }
        res, _, err := aclient.NewClient().Get(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/template/get_all_private_template?access_token=%s", accessToken), nil)
        if err != nil {
            errMsg = "请求微信消息模板列表出错"
            return nil, errMsg, err
        }
        type MsgTemplateTemp struct {
            TemplateList []MsgTemplate `json:"template_list"`
        }
        wxTmplTemp := MsgTemplateTemp{}
        err = json.Unmarshal(res, &wxTmplTemp)
        if err != nil {
            errMsg = "解析微信消息模板列表出错"
            return nil, errMsg, errors.New(string(res))
        }
        c.tmplMu.Lock()
        for _, wxTmpl := range wxTmplTemp.TemplateList {
            tmplTemp := &MsgTemplate{
                TemplateId: wxTmpl.TemplateId,
                Content:    wxTmpl.Content,
                timestamp:  time.Now().Unix(),
            }
            //fmt.Println(ajson.Encode(tmplTemp))
            c.tmpls[wxTmpl.TemplateId] = tmplTemp
            if wxTmpl.TemplateId == tmplId {
                tmpl = tmplTemp
            }
        }
        c.tmplMu.Unlock()
        if tmpl == nil {
            errMsg = "微信消息模板列表不存在"
            return nil, errMsg, errors.New(string(res))
        }
    }
    return tmpl, errMsg, nil
}

//获取微信消息模板数据
func (c *Client) GetTemplateParams(tmplContent string, dataParams map[string]string) (map[string]map[string]string, error) {
    params := map[string]map[string]string{}
    var paramKeys []string
    allMatches := regexp.MustCompile(`{{([\w|\d]+)\.DATA}}`).FindAllStringSubmatch(tmplContent, -1)
    for _, matchTemp := range allMatches {
        if len(matchTemp) > 1 {
            match := matchTemp[1]
            paramKeys = append(paramKeys, match)
        }
    }
    for _, key := range paramKeys {
        param := map[string]string{}
        param["value"] = dataParams[key]
        param["color"] = "#173177"
        params[key] = param
    }
    return params, nil
}

//替换数据库微信消息模板数据
func (c *Client) ReplaceTemplateParams(tmplContent string, dataParams map[string]string) string {
    for key, value := range dataParams {
        tmplContent = strings.ReplaceAll(tmplContent, fmt.Sprintf("{{%s.DATA}}", key), value)
    }
    return tmplContent
}

//发送微信模板消息
func (c *Client) SendTmplMsg(openid string, tmplId string, dataParams map[string]string) (errMsg string, err error) {
    if openid == "" {
        errMsg = "openid不能为空"
        return errMsg, errors.New(errMsg)
    }
    if tmplId == "" {
        errMsg = "tmplId不能为空"
        return errMsg, errors.New(errMsg)
    }

    tmpl, errMsg, err := c.GetTemplate(tmplId)
    if err != nil {
        return errMsg, err
    }

    data, err := c.GetTemplateParams(tmpl.Content, dataParams)
    if err != nil {
        errMsg = "获取模板参数失败"
        return errMsg, err
    }

    var accessToken string
    accessToken, errMsg, err = c.GetAccessToken()
    if err != nil {
        return errMsg, err
    }
    params := map[string]interface{}{}
    params["touser"] = openid
    params["template_id"] = tmplId
    params["topcolor"] = "#FF0000"
    params["url"] = ""
    params["data"] = data
    res, _, err := aclient.NewClient().PostJson(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", accessToken), params)
    if err != nil {
        errMsg = "微信模板消息发送失败"
        return errMsg, err
    }
    //fmt.Println(string(res))
    rs, err := ajson.DecodeToMapString(res)
    if err != nil {
        errMsg = "微信模板消息发送结果解析失败"
        return errMsg, err
    }
    if rs["errmsg"] == "" || rs["errmsg"] == "ok" {
        return "", nil
    } else {
        errMsg = rs["errmsg"]
        return errMsg, errors.New(rs["errmsg"])
    }
}

//发送微信客服消息
func (c *Client) SendCustomMsg(openid string, tmplContent string, dataParams map[string]string) (errMsg string, err error) {
    if openid == "" {
        errMsg = "openid不能为空"
        return errMsg, errors.New(errMsg)
    }
    if tmplContent == "" {
        errMsg = "tmplContent不能为空"
        return errMsg, errors.New(errMsg)
    }

    content := c.ReplaceTemplateParams(tmplContent, dataParams)

    var accessToken string
    accessToken, errMsg, err = c.GetAccessToken()
    if err != nil {
        return errMsg, err
    }
    params := map[string]interface{}{}
    params["touser"] = openid
    params["msgtype"] = "text"
    params["text"] = map[string]string{"content:": content}
    res, _, err := aclient.NewClient().PostJson(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s", accessToken), params)
    if err != nil {
        errMsg = "微信客服消息发送失败"
        return errMsg, err
    }
    //fmt.Println(string(res))
    rs, err := ajson.DecodeToMapString(res)
    if err != nil {
        errMsg = "微信客服消息发送结果解析失败"
        return errMsg, err
    }
    if rs["errmsg"] == "" || rs["errmsg"] == "ok" {
        return "", nil
    } else {
        errMsg = rs["errmsg"]
        return errMsg, errors.New(rs["errmsg"])
    }
}
