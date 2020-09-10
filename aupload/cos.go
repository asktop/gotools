package aupload

import (
    "context"
    "errors"
    "github.com/asktop/gotools/aclient"
    "github.com/asktop/gotools/afile"
    "github.com/asktop/gotools/astring"
    "github.com/tencentyun/cos-go-sdk-v5"
    "mime/multipart"
    "net/http"
    "net/url"
    "os"
    "path/filepath"
    "strings"
)

type CosClient struct {
    client *cos.Client
    site   string
    bucket string
}

type CosConfig struct {
    SiteUrl   string `json:"site_url"`
    Bucket    string `json:"bucket"` //文件基本路由，默认：upload
    SecretId  string `json:"secret_id"`
    SecretKey string `json:"secret_key"`
}

func NewCosClient(config CosConfig) (*CosClient, error) {
    if config.SiteUrl == "" {
        return nil, errors.New("cos:" + "bucket_url 不能为空")
    }
    if config.Bucket == "" {
        config.Bucket = "upload"
    }
    if config.SecretId == "" {
        return nil, errors.New("cos:" + "secret_id 不能为空")
    }
    if config.SecretKey == "" {
        return nil, errors.New("cos:" + "secret_key 不能为空")
    }

    u, _ := url.Parse(config.SiteUrl)
    b := &cos.BaseURL{BucketURL: u}
    client := cos.NewClient(b, &http.Client{
        Transport: &cos.AuthorizationTransport{
            SecretID:  config.SecretId,
            SecretKey: config.SecretKey,
        },
    })

    return &CosClient{client: client, site: config.SiteUrl, bucket: config.Bucket}, nil
}

func (c *CosClient) GetClient() *cos.Client {
    return c.client
}

func (c *CosClient) GetSite() string {
    return c.site
}

func (c *CosClient) GetBucket() string {
    return c.bucket
}

func (c *CosClient) GetUrl(uris ...string) string {
    if c.GetSite() != "" {
        return astring.JoinURL(c.GetSite(), c.GetBucket(), astring.JoinURL(uris...))
    } else {
        return c.GetUri(uris...)
    }
}

func (c *CosClient) GetUri(uris ...string) string {
    return astring.JoinURL(c.GetBucket(), astring.JoinURL(uris...))
}

func (c *CosClient) GetFilePath(uris ...string) string {
    return strings.TrimPrefix(c.GetUri(uris...), "/")
}

// 通过 文件 上传文件到cos
// @param file 文件
// @param filePathName cos文件存储路径
func (c *CosClient) UploadFromFile(file *os.File, filePathName string, checkSize ...int64) (fileInfo FileInfo, err error) {
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

    _, err = c.GetClient().Object.Put(context.Background(), fileInfo.Path, file, nil)
    if err != nil {
        return fileInfo, err
    }
    fileInfo.Url = c.GetUrl(filePathName)
    fileInfo.Uri = c.GetUri(filePathName)
    return fileInfo, nil
}

// 通过 文件FileHeader 上传文件到cos
// @param header 文件FileHeader
// @param filePathName cos文件存储路径
func (c *CosClient) UploadFromFileHeader(header *multipart.FileHeader, filePathName string, checkSize ...int64) (fileInfo FileInfo, err error) {
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

    file, err := header.Open()
    if err != nil {
        return fileInfo, err
    }
    defer file.Close()
    _, err = c.GetClient().Object.Put(context.Background(), fileInfo.Path, file, nil)
    if err != nil {
        return fileInfo, err
    }
    fileInfo.Url = c.GetUrl(filePathName)
    fileInfo.Uri = c.GetUri(filePathName)
    return fileInfo, nil
}

// 通过 文件绝对路径 上传文件到cos
// @param Path 文件绝对路径
// @param filePathName cos文件存储路径
func (c *CosClient) UploadFromPath(Path string, filePathName string, checkSize ...int64) (fileInfo FileInfo, err error) {
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

    _, err = c.GetClient().Object.PutFromFile(context.Background(), fileInfo.Path, Path, nil)
    if err != nil {
        return fileInfo, err
    }
    fileInfo.Url = c.GetUrl(filePathName)
    fileInfo.Path = strings.TrimPrefix(fileInfo.Url, c.GetSite())
    fileInfo.Uri = c.GetUri(filePathName)
    return fileInfo, nil
}

// 通过 url或文件存储路径 删除文件
// @param url 文件存储url
// @param filePathName 文件存储路径
func (c *CosClient) DeleteFile(url_filePathName string) (err error) {
    if url_filePathName == "" {
        return nil
    }
    filePathName := strings.TrimPrefix(url_filePathName, c.GetSite())
    filePathName = strings.TrimPrefix(filePathName, "/")

    _, err = c.GetClient().Object.Delete(context.Background(), filePathName)
    if err != nil {
        return err
    }
    return nil
}

//复制文件
func (c *CosClient) CopyFile(source_url_filePathName string, filePathName string) (fileInfo FileInfo, err error) {
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
            info, err := defaultClient.LocalClient.UploadFromByte(res, filepath.Join("temp", filePathName))
            if err != nil {
                return fileInfo, err
            } else {
                tempFilePath = info.Path
            }
        }
        fileInfo, err = c.UploadFromPath(defaultClient.LocalClient.GetUploadPath(tempFilePath), filePathName)
        defaultClient.LocalClient.DeleteFile(tempFilePath)
        return fileInfo, err
    } else {
        tempFilePath = strings.TrimPrefix(source_url_filePathName, "/")
        tempFilePath = strings.TrimPrefix(tempFilePath, defaultClient.LocalClient.GetBucket())
        tempFilePath = strings.TrimPrefix(tempFilePath, "/")
        fileInfo, err = c.UploadFromPath(defaultClient.LocalClient.GetUploadPath(tempFilePath), filePathName)
        return fileInfo, err
    }
}
