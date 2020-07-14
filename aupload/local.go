package aupload

import (
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
    baseUrl       string
    getUploadPath func(path ...string) string
}

type LocalConfig struct {
    Site   string //网址
    Bucket string //文件基本路由，默认：upload
}

func NewLocalClient(config LocalConfig, getUploadPath func(path ...string) string) *LocalClient {
    if config.Bucket == "" {
        config.Bucket = "upload"
    }
    if getUploadPath == nil {
        getUploadPath = defaultGetUploadPath
    }
    return &LocalClient{site: config.Site, bucket: config.Bucket, baseUrl: astring.JoinURL(config.Site, config.Bucket), getUploadPath: getUploadPath}
}

func (c *LocalClient) GetSite() string {
    return c.site
}

func (c *LocalClient) GetBucket() string {
    return c.bucket
}

func (c *LocalClient) GetBaseUrl() string {
    return c.baseUrl
}

func (c *LocalClient) GetAllUrl(uris ...string) string {
    return astring.JoinURL(c.GetBaseUrl(), astring.JoinURL(uris...))
}

//保存到本地
func (c *LocalClient) UploadFromByte(file []byte, filePathName string) (url string, err error) {
    filePathName = strings.TrimPrefix(filePathName, "/")
    filePath, fileName := path.Split(filePathName)

    //获取存储路径并创建文件夹
    localFilePathName := filepath.Join(c.getUploadPath(filePath), fileName)
    //获取文件存储流
    err = ioutil.WriteFile(localFilePathName, file, 0777)
    if err != nil {
        return "", err
    }
    url = c.GetAllUrl(filePathName)
    return url, nil
}

//保存到本地
func (c *LocalClient) UploadFromFile(file *os.File, filePathName string) (url string, err error) {
    filePathName = strings.TrimPrefix(filePathName, "/")
    filePath, fileName := path.Split(filePathName)

    //获取存储路径并创建文件夹
    localFilePathName := filepath.Join(c.getUploadPath(filePath), fileName)
    //获取文件存储流
    f, err := os.OpenFile(localFilePathName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
    if err != nil {
        return "", err
    }
    defer f.Close()
    io.Copy(f, file)
    url = c.GetAllUrl(filePathName)
    return url, nil
}

//保存到本地
func (c *LocalClient) UploadFromFileHeader(header *multipart.FileHeader, filePathName string) (url string, err error) {
    filePathName = strings.TrimPrefix(filePathName, "/")
    filePath, fileName := path.Split(filePathName)

    //获取存储路径并创建文件夹
    localFilePathName := filepath.Join(c.getUploadPath(filePath), fileName)
    //获取文件存储流
    f, err := os.OpenFile(localFilePathName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
    if err != nil {
        return "", err
    }
    defer f.Close()
    //获取文件读取流
    file, err := header.Open()
    if err != nil {
        return "", err
    }
    defer file.Close()
    io.Copy(f, file)
    url = c.GetAllUrl(filePathName)
    return url, nil
}

//保存到本地
func (c *LocalClient) UploadFromPath(Path string, filePathName string) (url string, err error) {
    filePathName = strings.TrimPrefix(filePathName, "/")
    filePath, fileName := path.Split(filePathName)

    //获取存储路径并创建文件夹
    localFilePathName := filepath.Join(c.getUploadPath(filePath), fileName)
    //获取文件存储流
    f, err := os.OpenFile(localFilePathName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
    if err != nil {
        return "", err
    }
    defer f.Close()
    //获取文件读取流
    file, err := os.Open(Path)
    if err != nil {
        return "", err
    }
    defer file.Close()
    io.Copy(f, file)
    url = c.GetAllUrl(filePathName)
    return url, nil
}

//从本地删除
func (c *LocalClient) DeleteFile(url_filePathName string) (err error) {
    filePathName := strings.TrimPrefix(url_filePathName, c.GetBaseUrl())
    filePathName = strings.TrimPrefix(filePathName, "/")

    //获取存储路径
    localFilePathName := c.getUploadPath(filePathName)
    return afile.Delete(localFilePathName)
}
