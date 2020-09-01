package aupload

import (
    "mime/multipart"
    "os"
)

const (
    DriverLocal driver = "local"
    DriverCos   driver = "cos"
)

//文件上传位置
type driver string

type Client struct {
    GetUploadPath func(path ...string) string //获取该路径的本地缓存或保存的绝对路径
    LocalClient   *LocalClient
    defaultDriver driver //保存方式
    cosClient     *CosClient
}

type Config struct {
    Cos CosConfig
}

type FileInfo struct {
    Url     string `json:"url"`      //文件访问url（网址若为空，则为uri）
    Path    string `json:"path"`     //文件在存储库的位置（不包括upload）
    OldName string `json:"old_name"` //旧文件名（不包括扩展名）
}

//创建一个文件上传位置客户端（本地或cos等）
//defaultDriver 默认文件上传位置
//localConfig 本地文件上传配置
//config 文件上传位置相关配置
func NewClient(defaultDriver driver, localConfig LocalConfig, config Config) (*Client, error) {
    client := &Client{
        defaultDriver: defaultDriver,
    }
    client.LocalClient = NewLocalClient(localConfig)
    client.GetUploadPath = client.LocalClient.GetUploadPath

    switch defaultDriver {
    case DriverCos:
        cosClient, err := NewCosClient(config.Cos)
        client.cosClient = cosClient
        return client, err
    default:
        return client, nil
    }
}

//上传单个文件
func (c *Client) UploadFromFile(file *os.File, filePathName string, checkSize ...int64) (fileInfo FileInfo, err error) {
    switch c.defaultDriver {
    case DriverCos:
        return c.cosClient.UploadFromFile(file, filePathName, checkSize ...)
    default:
        return c.LocalClient.UploadFromFile(file, filePathName, checkSize ...)
    }
}

//上传单个文件
func (c *Client) UploadFromFileHeader(header *multipart.FileHeader, filePathName string, checkSize ...int64) (fileInfo FileInfo, err error) {
    switch c.defaultDriver {
    case DriverCos:
        return c.cosClient.UploadFromFileHeader(header, filePathName, checkSize ...)
    default:
        return c.LocalClient.UploadFromFileHeader(header, filePathName, checkSize ...)
    }
}

//上传单个文件
func (c *Client) UploadFromPath(Path string, filePathName string, checkSize ...int64) (fileInfo FileInfo, err error) {
    switch c.defaultDriver {
    case DriverCos:
        return c.cosClient.UploadFromPath(Path, filePathName, checkSize ...)
    default:
        return c.LocalClient.UploadFromPath(Path, filePathName, checkSize ...)
    }
}

//删除文件
func (c *Client) DeleteFile(url_filePathName string) (err error) {
    switch c.defaultDriver {
    case DriverCos:
        return c.cosClient.DeleteFile(url_filePathName)
    default:
        return c.LocalClient.DeleteFile(url_filePathName)
    }
}
