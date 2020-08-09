package aupload

import (
    "mime/multipart"
    "os"
)

var defaultClient, _ = NewClient(nil, DriverLocal, Config{})

//创建一个默认全局文件上传位置客户端（本地或cos等）
//getUploadPath 获取服务器本地文件上传根目录绝对路径的方法
//driver 文件上传位置
//config 文件上传位置相关配置
func DefaultClient(getUploadPath func(path ...string) string, driver driver, config Config) error {
    client, err := NewClient(getUploadPath, driver, config)
    if err != nil {
        return err
    } else {
        defaultClient = client
        return nil
    }
}

//上传单个文件
func UploadFromByte(file []byte, filePathName string) (fileInfo FileInfo, err error) {
    return defaultClient.UploadFromByte(file, filePathName)
}

//上传单个文件
func UploadFromFile(file *os.File, filePathName string, checkSize ...int64) (fileInfo FileInfo, err error) {
    return defaultClient.UploadFromFile(file, filePathName, checkSize...)
}

//上传单个文件
func UploadFromFileHeader(header *multipart.FileHeader, filePathName string, checkSize ...int64) (fileInfo FileInfo, err error) {
    return defaultClient.UploadFromFileHeader(header, filePathName, checkSize...)
}

//上传单个文件
func UploadFromPath(Path string, filePathName string, checkSize ...int64) (fileInfo FileInfo, err error) {
    return defaultClient.UploadFromPath(Path, filePathName, checkSize...)
}

//删除文件
func DeleteFile(url_filePathName string) (err error) {
    return defaultClient.DeleteFile(url_filePathName)
}
