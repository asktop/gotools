package aupload

//import (
//    "github.com/asktop/gotools/astring"
//    "github.com/minio/minio-go"
//    "mime/multipart"
//    "os"
//    "strings"
//)
//
//type MinioClient struct {
//    client  *minio.Client
//    site    string
//    bucket  string
//    baseUrl string
//}
//
//type MinioConfig struct {
//    Scheme    string
//    Endpoint  string
//    AccessKey string
//    SecretKey string
//    Bucket    string
//    Site      string
//}
//
////连接局域网oss
////文档：https://docs.minio.io/cn/
//func NewMinioClient(config MinioConfig) (*MinioClient, error) {
//    scheme := config.Scheme
//    if scheme == "" {
//        scheme = "http"
//    }
//    endpoint := config.Endpoint
//    if endpoint == "" {
//        endpoint = "127.0.0.1:9000"
//    }
//    accesskey := config.AccessKey
//    secretkey := config.SecretKey
//    bucket := config.Bucket
//    if bucket == "" {
//        bucket = "upload"
//    }
//    site := config.Site
//
//    secure := false
//    if scheme == "https" {
//        secure = true
//    } else {
//        scheme = "http"
//    }
//    if site == "" {
//        site = scheme + "://" + endpoint
//    }
//    site = strings.TrimRight(site, "/")
//
//    client, err := minio.New(endpoint, accesskey, secretkey, secure)
//    if err != nil {
//        return nil, err
//    }
//    ok, err := client.BucketExists(bucket)
//    if err != nil {
//        return nil, err
//    } else if !ok {
//        err = client.MakeBucket(bucket, "")
//        if err != nil {
//            return nil, err
//        }
//    }
//
//    minioClient := &MinioClient{
//        client:  client,
//        site:    site,
//        bucket:  bucket,
//        baseUrl: astring.JoinURL(site, bucket),
//    }
//    return minioClient, nil
//}
//
//func (c *MinioClient) GetClient() *minio.Client {
//    return c.client
//}
//
//func (c *MinioClient) GetSite() string {
//    return c.site
//}
//
//func (c *MinioClient) GetBucket() string {
//    return c.bucket
//}
//
//func (c *MinioClient) GetBaseUrl() string {
//    return c.baseUrl
//}
//
//func (c *MinioClient) GetAllUrl(uris ...string) string {
//    return astring.JoinURL(c.GetBaseUrl(), astring.JoinURL(uris...))
//}
//
//// 通过 文件 上传文件到minio
//// @param file 文件
//// @param filePathName minio文件存储路径
//func (c *MinioClient) UploadFromFile(file *os.File, filePathName string) (url string, err error) {
//    filePathName = strings.TrimPrefix(filePathName, "/")
//
//    fInfo, _ := file.Stat()
//    _, err = c.GetClient().PutObject(c.GetBucket(), filePathName, file, fInfo.Size(), minio.PutObjectOptions{ContentType: GetContentType(fInfo.Name())})
//    if err != nil {
//        return url, err
//    }
//    url = c.GetAllUrl(filePathName)
//    return
//}
//
//// 通过 文件FileHeader 上传文件到minio
//// @param header 文件FileHeader
//// @param filePathName minio文件存储路径
//func (c *MinioClient) UploadFromFileHeader(header *multipart.FileHeader, filePathName string) (url string, err error) {
//    filePathName = strings.TrimPrefix(filePathName, "/")
//
//    file, err := header.Open()
//    if err != nil {
//        return url, err
//    }
//    defer file.Close()
//    _, err = c.GetClient().PutObject(c.GetBucket(), filePathName, file, header.Size, minio.PutObjectOptions{ContentType: GetContentType(header.Filename)})
//    if err != nil {
//        return url, err
//    }
//    url = c.GetAllUrl(filePathName)
//    return
//}
//
//// 通过 文件绝对路径 上传文件到minio
//// @param Path 文件绝对路径
//// @param filePathName minio文件存储路径
//func (c *MinioClient) UploadFromPath(Path string, filePathName string) (url string, err error) {
//    filePathName = strings.TrimPrefix(filePathName, "/")
//
//    file, err := os.Open(Path)
//    if err != nil {
//        return
//    }
//    defer file.Close()
//    fInfo, _ := file.Stat()
//    _, err = c.GetClient().PutObject(c.GetBucket(), filePathName, file, fInfo.Size(), minio.PutObjectOptions{ContentType: GetContentType(fInfo.Name())})
//    if err != nil {
//        return url, err
//    }
//    url = c.GetAllUrl(filePathName)
//    return
//}
//
//// 通过 url或文件存储路径 删除文件
//// @param url 文件存储url
//// @param filePathName 文件存储路径
//func (c *MinioClient) DeleteFile(url_filePathName string) (err error) {
//    filePathName := strings.TrimPrefix(url_filePathName, c.GetBaseUrl())
//    filePathName = strings.TrimPrefix(filePathName, "/")
//
//    return c.GetClient().RemoveObject(c.GetBucket(), filePathName)
//}
