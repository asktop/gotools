package aupload

import (
    "context"
    "errors"
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
    client  *cos.Client
    site    string
    bucket  string
    baseUrl string
}

type CosConfig struct {
    BucketUrl string `json:"bucket_url"`
    SecretId  string `json:"secret_id"`
    SecretKey string `json:"secret_key"`
}

func NewCosClient(config CosConfig) (*CosClient, error) {
    bucketUrl := config.BucketUrl
    secretID := config.SecretId
    secretKey := config.SecretKey

    if bucketUrl == "" {
        return nil, errors.New("cos:" + "bucket_url 不能为空")
    }
    if secretID == "" {
        return nil, errors.New("cos:" + "secret_id 不能为空")
    }
    if secretKey == "" {
        return nil, errors.New("cos:" + "secret_key 不能为空")
    }

    bucketUrl = strings.TrimRight(bucketUrl, "/")
    var site, bucket string
    index := strings.LastIndex(bucketUrl, "/")
    if index > 0 && index != len(bucketUrl)-1 {
        site = bucketUrl[:index]
        bucket = bucketUrl[index+1:]
    }
    u, _ := url.Parse(bucketUrl)
    b := &cos.BaseURL{BucketURL: u}
    client := cos.NewClient(b, &http.Client{
        Transport: &cos.AuthorizationTransport{
            SecretID:  secretID,
            SecretKey: secretKey,
        },
    })

    return &CosClient{client: client, site: site, bucket: bucket, baseUrl: bucketUrl}, nil
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

func (c *CosClient) GetBaseUrl() string {
    return c.baseUrl
}

func (c *CosClient) GetAllUrl(uris ...string) string {
    return astring.JoinURL(c.GetBaseUrl(), astring.JoinURL(uris...))
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
    fileInfo.Path = filePathName

    _, err = c.GetClient().Object.Put(context.Background(), filePathName, file, nil)
    if err != nil {
        return fileInfo, err
    }
    fileInfo.Url = c.GetAllUrl(filePathName)
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
    fileInfo.Path = filePathName

    file, err := header.Open()
    if err != nil {
        return fileInfo, err
    }
    defer file.Close()
    _, err = c.GetClient().Object.Put(context.Background(), filePathName, file, nil)
    if err != nil {
        return fileInfo, err
    }
    fileInfo.Url = c.GetAllUrl(filePathName)
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
    fileInfo.Path = filePathName

    _, err = c.GetClient().Object.PutFromFile(context.Background(), filePathName, Path, nil)
    if err != nil {
        return fileInfo, err
    }
    fileInfo.Url = c.GetAllUrl(filePathName)
    fileInfo.Path = strings.TrimPrefix(fileInfo.Url, c.GetSite())
    return fileInfo, nil
}

// 通过 url或文件存储路径 删除文件
// @param url 文件存储url
// @param filePathName 文件存储路径
func (c *CosClient) DeleteFile(url_filePathName string) (err error) {
    if url_filePathName == "" {
        return nil
    }
    filePathName := strings.TrimPrefix(url_filePathName, c.GetBaseUrl())
    filePathName = strings.TrimPrefix(filePathName, "/")

    _, err = c.GetClient().Object.Delete(context.Background(), filePathName)
    if err != nil {
        return err
    }
    return nil
}
