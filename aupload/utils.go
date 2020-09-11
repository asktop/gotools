package aupload

import (
    "github.com/asktop/gotools/acast"
    "github.com/asktop/gotools/afile"
    "path/filepath"
    "strings"
)

//获取服务器本地文件上传目录绝对路径的默认方法
func defaultGetUploadPath(path ...string) string {
    localPath := filepath.Join(path...)
    localPath = filepath.Join("upload", localPath)
    localPath, _ = filepath.Abs(localPath)
    afile.CreateDir(localPath)
    return localPath
}

//生成不同角色文件存放路径（不包括扩展名）
func NewFilePathName(role string, roleId int64, tableName string, tableColumn string, fileName ...string) string {
    var newFilePathName string
    if role != "" {
        newFilePathName = role
        if roleId < 0 {
            roleId = 0
        }
        newFilePathName = filepath.Join(newFilePathName, acast.ToString(roleId))
    }
    if tableName != "" {
        newFilePathName = filepath.Join(newFilePathName, tableName)
    }
    if tableColumn != "" {
        newFilePathName = filepath.Join(newFilePathName, tableColumn)
    }
    if len(fileName) > 0 {
        fName := fileName[0]
        ext := filepath.Ext(fName)
        if ext != "" {
            fName = strings.TrimSuffix(fName, ext)
        }
        newFilePathName = filepath.Join(newFilePathName, fName)
    } else {
        newFilePathName = filepath.Join(newFilePathName, afile.NewFileName())
    }
    newFilePathName = strings.ReplaceAll(newFilePathName, `\`, `/`)
    return newFilePathName
}

func NewOtherFilePathName(fileName ...string) string {
    var newFilePathName string
    if len(fileName) > 0 {
        newFilePathName = filepath.Join(fileName...)
        ext := filepath.Ext(newFilePathName)
        if ext != "" {
            newFilePathName = strings.TrimSuffix(newFilePathName, ext)
        }
    } else {
        newFilePathName = filepath.Join(newFilePathName, afile.NewFileName())
    }
    newFilePathName = strings.ReplaceAll(newFilePathName, `\`, `/`)
    return newFilePathName
}
