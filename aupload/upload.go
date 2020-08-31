package aupload

import (
    "errors"
    "github.com/asktop/gotools/afile"
    "mime/multipart"
    "os"
    "path/filepath"
    "strings"
)

const (
    DriverLocal driver = "local"
    DriverCos   driver = "cos"
    //DriverMinio driver = "minio"
)

//文件上传位置
type driver string

type Client struct {
    GetUploadPath func(path ...string) string //获取该路径的本地缓存或保存的绝对路径
    driver        driver                      //保存方式
    localClient   *LocalClient
    cosClient     *CosClient
    //minioClient   *MinioClient
}

type Config struct {
    Local LocalConfig
    Cos   CosConfig
    //Minio MinioConfig
}

type FileInfo struct {
    Url     string `json:"url"`      //文件访问url
    Path    string `json:"path"`     //文件在存储库的位置
    OldName string `json:"old_name"` //旧文件名（不包括扩展名）
}

//创建一个文件上传位置客户端（本地或cos等）
//getUploadPath 获取服务器本地文件上传根目录绝对路径的方法
//driver 文件上传位置
//config 文件上传位置相关配置
func NewClient(getUploadPath func(path ...string) string, driver driver, config Config) (*Client, error) {
    if getUploadPath == nil {
        getUploadPath = defaultGetUploadPath
    }

    client := &Client{
        GetUploadPath: getUploadPath,
        driver:        driver,
    }

    localClient := NewLocalClient(config.Local, getUploadPath)
    client.localClient = localClient

    switch driver {
    //case DriverMinio:
    //    minioClient, err := NewMinioClient(config.Minio)
    //    client.minioClient = minioClient
    //    return client, err
    case DriverCos:
        cosClient, err := NewCosClient(config.Cos)
        client.cosClient = cosClient
        return client, err
    default:
        return client, nil
    }
}

//上传单个文件
func (c *Client) UploadFromByte(file []byte, filePathName string) (fileInfo FileInfo, err error) {
    filePathName = strings.Trim(strings.TrimSpace(filePathName), "/")

    if file == nil {
        err = errors.New("file 不能为空")
        return
    }
    if filepath.Ext(filePathName) == "" {
        err = errors.New("filePathName 扩展名不能为空")
        return
    }
    fileInfo.Path = filePathName
    filePath, fileName := filepath.Split(filePathName)
    fileInfo.OldName = afile.NameNoExt(fileName)

    fileInfo.Url, err = c.localClient.UploadFromByte(file, filePathName)
    switch c.driver {
    //case DriverMinio:
    //    localFilePathName := filepath.Join(c.GetUploadPath(filePath), fileName)
    //    fileInfo.Url, err = c.minioClient.UploadFromPath(localFilePathName, filePathName)
    case DriverCos:
        localFilePathName := filepath.Join(c.GetUploadPath(filePath), fileName)
        fileInfo.Url, err = c.cosClient.UploadFromPath(localFilePathName, filePathName)
    default:

    }
    return
}

//上传单个文件
func (c *Client) UploadFromFile(file *os.File, filePathName string, checkSize ...int64) (fileInfo FileInfo, err error) {
    filePathName = strings.Trim(strings.TrimSpace(filePathName), "/")

    if file == nil {
        err = errors.New("file 不能为空")
        return
    }
    fInfo, _ := file.Stat()
    if len(checkSize) > 0 {
        if fInfo.Size() > checkSize[0] {
            err = errors.New("文件过大")
            return
        }
    }
    oldName := fInfo.Name()
    fileInfo.OldName = afile.NameNoExt(oldName)
    if filepath.Ext(filePathName) == "" {
        filePathName += filepath.Ext(oldName)
    }
    fileInfo.Path = filePathName

    switch c.driver {
    //case DriverMinio:
    //   fileInfo.Url, err = c.minioClient.UploadFromFile(file, filePathName)
    case DriverCos:
        fileInfo.Url, err = c.cosClient.UploadFromFile(file, filePathName)
    default:
        fileInfo.Url, err = c.localClient.UploadFromFile(file, filePathName)
    }
    return
}

//上传单个文件
func (c *Client) UploadFromFileHeader(header *multipart.FileHeader, filePathName string, checkSize ...int64) (fileInfo FileInfo, err error) {
    filePathName = strings.Trim(strings.TrimSpace(filePathName), "/")

    if header == nil {
        err = errors.New("header 不能为空")
        return
    }
    if len(checkSize) > 0 {
        if header.Size > checkSize[0] {
            err = errors.New("文件过大")
            return
        }
    }
    oldName := header.Filename
    fileInfo.OldName = afile.NameNoExt(oldName)
    if filepath.Ext(filePathName) == "" {
        filePathName += filepath.Ext(oldName)
    }
    fileInfo.Path = filePathName

    switch c.driver {
    //case DriverMinio:
    //    fileInfo.Url, err = c.minioClient.UploadFromFileHeader(header, filePathName)
    case DriverCos:
        fileInfo.Url, err = c.cosClient.UploadFromFileHeader(header, filePathName)
    default:
        fileInfo.Url, err = c.localClient.UploadFromFileHeader(header, filePathName)
    }
    return
}

//上传单个文件
func (c *Client) UploadFromPath(Path string, filePathName string, checkSize ...int64) (fileInfo FileInfo, err error) {
    filePathName = strings.Trim(strings.TrimSpace(filePathName), "/")

    Path = strings.Trim(strings.TrimSpace(Path), "/")
    if Path == "" {
        err = errors.New("Path 不能为空")
        return
    }
    file, err := os.Open(Path)
    if err != nil {
        return
    }
    defer file.Close()
    fInfo, _ := file.Stat()
    if len(checkSize) > 0 {
        if fInfo.Size() > checkSize[0] {
            err = errors.New("文件过大")
            return
        }
    }
    oldName := fInfo.Name()
    fileInfo.OldName = afile.NameNoExt(oldName)
    if filepath.Ext(filePathName) == "" {
        filePathName += filepath.Ext(oldName)
    }
    fileInfo.Path = filePathName

    switch c.driver {
    //case DriverMinio:
    //    fileInfo.Url, err = c.minioClient.UploadFromFile(file, filePathName)
    case DriverCos:
        fileInfo.Url, err = c.cosClient.UploadFromFile(file, filePathName)
        fileInfo.Path = strings.TrimPrefix(fileInfo.Url, c.cosClient.GetSite())
    default:
        fileInfo.Url, err = c.localClient.UploadFromFile(file, filePathName)
    }
    return
}

//删除文件
func (c *Client) DeleteFile(url_filePathName string) (err error) {
    if url_filePathName == "" {
        return nil
    }
    switch c.driver {
    //case DriverMinio:
    //    err = c.minioClient.DeleteFile(url_filePathName)
    case DriverCos:
        err = c.cosClient.DeleteFile(url_filePathName)
    default:
        err = c.localClient.DeleteFile(url_filePathName)
    }
    return
}
