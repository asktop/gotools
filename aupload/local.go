package aupload

import (
    "errors"
    "github.com/asktop/gotools/aclient"
    "github.com/asktop/gotools/afile"
    "github.com/asktop/gotools/astring"
    "io"
    "io/ioutil"
    "mime/multipart"
    "os"
    "path"
    "path/filepath"
    "strings"
)

type LocalClient struct {
    site          string
    bucket        string
    GetUploadPath func(path ...string) string
}

type LocalConfig struct {
    Site          string                      `json:"site"`   //网址
    Bucket        string                      `json:"bucket"` //文件基本路由，默认：upload
    GetUploadPath func(path ...string) string `json:"-"`      //获取本地文件上传本地绝对路径
}

func NewLocalClient(config LocalConfig) *LocalClient {
    if config.Bucket == "" {
        config.Bucket = "upload"
    }
    if config.GetUploadPath == nil {
        config.GetUploadPath = defaultGetUploadPath
    }
    return &LocalClient{site: config.Site, bucket: config.Bucket, GetUploadPath: config.GetUploadPath}
}

func (c *LocalClient) GetSite() string {
    return c.site
}

func (c *LocalClient) GetBucket() string {
    return c.bucket
}

func (c *LocalClient) GetUrl(uris ...string) string {
    if c.GetSite() != "" {
        return astring.JoinURL(c.GetSite(), c.GetBucket(), astring.JoinURL(uris...))
    } else {
        return c.GetUri(uris...)
    }
}

func (c *LocalClient) GetUri(uris ...string) string {
    return astring.JoinURL(c.GetBucket(), astring.JoinURL(uris...))
}

func (c *LocalClient) GetFilePath(uris ...string) string {
    return strings.TrimPrefix(astring.JoinURL(uris...), "/")
}

//保存到本地
func (c *LocalClient) UploadFromByte(file []byte, filePathName string) (fileInfo FileInfo, err error) {
    if file == nil {
        return fileInfo, errors.New("file 不能为空")
    }
    filePathName = strings.Trim(strings.TrimSpace(filePathName), "/")
    if filepath.Ext(filePathName) == "" {
        return fileInfo, errors.New("filePathName 扩展名不能为空")
    }
    fileInfo.Path = c.GetFilePath(filePathName)

    filePath, fileName := filepath.Split(fileInfo.Path)
    fileInfo.OldName = afile.NameNoExt(fileName)

    //获取存储路径并创建文件夹
    localFilePathName := filepath.Join(c.GetUploadPath(filePath), fileName)
    //获取文件存储流
    err = ioutil.WriteFile(localFilePathName, file, 0777)
    if err != nil {
        return fileInfo, err
    }
    fileInfo.Url = c.GetUrl(filePathName)
    fileInfo.Uri = c.GetUri(filePathName)
    return fileInfo, nil
}

//保存到本地
func (c *LocalClient) UploadFromFile(file *os.File, filePathName string, checkSize ...int64) (fileInfo FileInfo, err error) {
    if file == nil {
        return fileInfo, errors.New("file 不能为空")
    }
    fInfo, _ := file.Stat()
    if len(checkSize) > 0 {
        if fInfo.Size() > checkSize[0] {
            return fileInfo, errors.New("文件过大")
        }
    }
    oldName := fInfo.Name()
    fileInfo.OldName = afile.NameNoExt(oldName)

    filePathName = strings.Trim(strings.TrimSpace(filePathName), "/")
    if filepath.Ext(filePathName) == "" {
        filePathName += filepath.Ext(oldName)
    }
    fileInfo.Path = c.GetFilePath(filePathName)

    filePath, fileName := path.Split(fileInfo.Path)

    //获取存储路径并创建文件夹
    localFilePathName := filepath.Join(c.GetUploadPath(filePath), fileName)
    //获取文件存储流
    f, err := os.OpenFile(localFilePathName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
    if err != nil {
        return fileInfo, err
    }
    defer f.Close()
    io.Copy(f, file)
    fileInfo.Url = c.GetUrl(filePathName)
    fileInfo.Uri = c.GetUri(filePathName)
    return fileInfo, nil
}

//保存到本地
func (c *LocalClient) UploadFromFileHeader(header *multipart.FileHeader, filePathName string, checkSize ...int64) (fileInfo FileInfo, err error) {
    if header == nil {
        return fileInfo, errors.New("header 不能为空")
    }
    if len(checkSize) > 0 {
        if header.Size > checkSize[0] {
            return fileInfo, errors.New("文件过大")
        }
    }
    oldName := header.Filename
    fileInfo.OldName = afile.NameNoExt(oldName)

    filePathName = strings.Trim(strings.TrimSpace(filePathName), "/")
    if filepath.Ext(filePathName) == "" {
        filePathName += filepath.Ext(oldName)
    }
    fileInfo.Path = c.GetFilePath(filePathName)

    filePath, fileName := path.Split(fileInfo.Path)

    //获取存储路径并创建文件夹
    localFilePathName := filepath.Join(c.GetUploadPath(filePath), fileName)
    //获取文件存储流
    f, err := os.OpenFile(localFilePathName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
    if err != nil {
        return fileInfo, err
    }
    defer f.Close()
    //获取文件读取流
    file, err := header.Open()
    if err != nil {
        return fileInfo, err
    }
    defer file.Close()
    io.Copy(f, file)
    fileInfo.Url = c.GetUrl(filePathName)
    fileInfo.Uri = c.GetUri(filePathName)
    return fileInfo, nil
}

//保存到本地
func (c *LocalClient) UploadFromPath(Path string, filePathName string, checkSize ...int64) (fileInfo FileInfo, err error) {
    Path = strings.Trim(strings.TrimSpace(Path), "/")
    if Path == "" {
        return fileInfo, errors.New("Path 不能为空")
    }
    file, err := os.Open(Path)
    if err != nil {
        return fileInfo, err
    }
    defer file.Close()
    fInfo, _ := file.Stat()
    if len(checkSize) > 0 {
        if fInfo.Size() > checkSize[0] {
            return fileInfo, errors.New("文件过大")
        }
    }
    oldName := fInfo.Name()
    fileInfo.OldName = afile.NameNoExt(oldName)

    filePathName = strings.Trim(strings.TrimSpace(filePathName), "/")
    if filepath.Ext(filePathName) == "" {
        filePathName += filepath.Ext(oldName)
    }
    fileInfo.Path = c.GetFilePath(filePathName)

    filePath, fileName := path.Split(fileInfo.Path)

    //获取存储路径并创建文件夹
    localFilePathName := filepath.Join(c.GetUploadPath(filePath), fileName)
    //获取文件存储流
    f, err := os.OpenFile(localFilePathName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
    if err != nil {
        return fileInfo, err
    }
    defer f.Close()
    //获取文件读取流
    newfile, err := os.Open(Path)
    if err != nil {
        return fileInfo, err
    }
    defer newfile.Close()
    io.Copy(f, newfile)
    fileInfo.Url = c.GetUrl(filePathName)
    fileInfo.Uri = c.GetUri(filePathName)
    return fileInfo, nil
}

//从本地删除
func (c *LocalClient) DeleteFile(url_filePathName string) (err error) {
    if url_filePathName == "" {
        return nil
    }
    filePathName := strings.TrimPrefix(url_filePathName, c.GetSite())
    filePathName = strings.TrimPrefix(filePathName, "/")
    filePathName = strings.TrimPrefix(filePathName, c.GetBucket())
    filePathName = strings.TrimPrefix(filePathName, "/")

    //获取存储路径
    localFilePathName := c.GetUploadPath(filePathName)
    return afile.Delete(localFilePathName)
}

//复制文件
func (c *LocalClient) CopyFile(source_url_filePathName string, filePathName string) (fileInfo FileInfo, err error) {
    if source_url_filePathName == "" || filePathName == "" {
        return fileInfo, errors.New("文件不能为空")
    }
    var tempFilePath string
    if strings.HasPrefix(source_url_filePathName, "http://") || strings.HasPrefix(source_url_filePathName, "https://") {
        res, _, err := aclient.NewClient().Get(source_url_filePathName, nil)
        if err != nil {
            return fileInfo, err
        } else {
            if filepath.Ext(filePathName) == "" {
                filePathName += filepath.Ext(source_url_filePathName)
            }
            info, err := c.UploadFromByte(res, filepath.Join("temp", filePathName))
            if err != nil {
                return fileInfo, err
            } else {
                tempFilePath = info.Path
            }
        }
        fileInfo, err = c.UploadFromPath(c.GetUploadPath(tempFilePath), filePathName)
        c.DeleteFile(tempFilePath)
        return fileInfo, err
    } else {
        tempFilePath = strings.TrimPrefix(source_url_filePathName, "/")
        tempFilePath = strings.TrimPrefix(tempFilePath, c.GetBucket())
        tempFilePath = strings.TrimPrefix(tempFilePath, "/")
        fileInfo, err = c.UploadFromPath(c.GetUploadPath(tempFilePath), filePathName)
        return fileInfo, err
    }
}

//获取文件列表
func (c *LocalClient) GetFiles(fileDir string) (fileInfos []FileInfo, err error) {
    dir := c.GetUploadPath(fileDir)
    _, filePaths, err := afile.GetAllPaths(dir)
    if err != nil {
        return fileInfos, err
    }
    udir := c.GetUploadPath()
    for _, f := range filePaths {
        fileInfo := FileInfo{}
        fileInfo.Path = strings.TrimPrefix(astring.JoinURL(strings.TrimPrefix(f, udir)), "/")
        fileInfo.Uri = astring.JoinURL(c.GetBucket(), fileInfo.Path)
        fileInfo.Url = astring.JoinURL(c.GetSite(), fileInfo.Uri)
        fileInfos = append(fileInfos, fileInfo)
    }
    return fileInfos, nil
}

//删除文件夹下所有文件
func (c *LocalClient) DeleteDir(fileDir string) (err error) {
    //获取存储路径
    localFilePathName := c.GetUploadPath(fileDir)
    return afile.Delete(localFilePathName, true)
}
