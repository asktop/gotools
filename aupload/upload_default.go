package aupload

import (
    "mime/multipart"
    "os"
)

var defaultClient, _ = NewClient(DriverLocal, LocalConfig{}, Config{})

//设置默认全局文件上传位置客户端（本地或cos等）
//defaultDriver 默认文件上传位置
//localConfig 本地文件上传配置
//config 文件上传位置相关配置
func SetDefaultClient(defaultDriver driver, localConfig LocalConfig, config Config) error {
    client, err := NewClient(defaultDriver, localConfig, config)
    if err != nil {
        return err
    } else {
        defaultClient = client
        return nil
    }
}

//获取默认客户端
func GetDefaultClient() *Client {
    return defaultClient
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
