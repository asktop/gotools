package afile

import (
    "errors"
    "github.com/asktop/gotools/arand"
    "github.com/asktop/gotools/atime"
    "io/ioutil"
    "os"
    "path"
    "path/filepath"
    "strings"
)

// 获取的当前运行路径
func GetWorkDir() string {
    workDir, _ := os.Getwd()
    return workDir
}

// 判断文件或文件夹否存在
func IsExist(path string) bool {
    _, err := os.Stat(path)
    if err == nil {
        return true
    }
    return os.IsExist(err)
}

// 删除文件或文件夹
// @param force 强制删除非空的文件夹
func Delete(path string, forces ...bool) error {
    if path == "" {
        return errors.New("path can not be empty")
    }
    ok := IsExist(path)
    if !ok {
        return nil
    }
    if len(forces) > 0 && forces[0] {
        return os.RemoveAll(path)
    } else {
        return os.Remove(path)
    }
}

// 创建文件夹（递归）
// @param perm 权限 0777（读4写2执行1）
func CreateDir(dirPath string, perms ...os.FileMode) error {
    if dirPath == "" {
        return errors.New("path can not be empty")
    }
    ok := IsExist(dirPath)
    if ok {
        return nil
    }
    var perm os.FileMode
    if len(perms) > 0 {
        perm = perms[0]
    } else {
        perm = 0777
    }
    err := os.MkdirAll(dirPath, perm)
    if err != nil {
        return err
    }
    return nil
}

//写入文本文件内容
// @param force 文件夹不存在时自动创建
func WriteFile(filePath string, body string, forces ...bool) error {
    if len(forces) > 0 && forces[0] {
        dir := strings.Replace(filePath, `\`, "/", -1)
        index := strings.LastIndex(dir, "/")
        dir = dir[:index]
        err := CreateDir(dir)
        if err != nil {
            return err
        }
    }
    return ioutil.WriteFile(filePath, []byte(body), 0777)
}

//读取文本文件内容
func ReadFile(filePath string) (string, error) {
    if filePath == "" {
        return "", errors.New("path can not be empty")
    }
    ok := IsExist(filePath)
    if !ok {
        return "", errors.New("file does not exist, path：" + filePath)
    }
    body, err := ioutil.ReadFile(filePath)
    if err != nil {
        return "", err
    }
    return string(body), nil
}

//获取所有文件和文件夹的文件名
// @param ext 过滤文件，只获取匹配后缀名的文件，示例：.go
func GetNames(dirPath string, exts ...string) (dirNames []string, fileNames []string, err error) {
    // 处理要过滤的后缀名
    var ext string
    if len(exts) > 0 {
        ext = path.Ext(exts[0])
        if ext == "" {
            err = errors.New("ext format incorrect, ext:" + exts[0])
            return
        }
    }

    // 读取文件和文件夹
    files, err := ioutil.ReadDir(dirPath)
    if err != nil {
        return
    }
    for _, file := range files {
        if file.IsDir() {
            dirNames = append(dirNames, file.Name())
        } else {
            if ext != "" && path.Ext(file.Name()) != ext {
                continue
            }
            fileNames = append(fileNames, file.Name())
        }
    }
    return
}

//获取所有文件和文件夹的路径
// @param ext 过滤文件，只获取匹配后缀名的文件，示例：.go
func GetPaths(dirPath string, exts ...string) (dirPaths []string, filePaths []string, err error) {
    // 处理要过滤的后缀名
    var ext string
    if len(exts) > 0 {
        ext = path.Ext(exts[0])
        if ext == "" {
            err = errors.New("ext format incorrect, ext:" + exts[0])
            return
        }
    }

    // 读取文件和文件夹
    files, err := ioutil.ReadDir(dirPath)
    if err != nil {
        return
    }
    for _, file := range files {
        if file.IsDir() {
            dirPaths = append(dirPaths, filepath.Join(dirPath, file.Name()))
        } else {
            if ext != "" && path.Ext(file.Name()) != ext {
                continue
            }
            filePaths = append(filePaths, filepath.Join(dirPath, file.Name()))
        }
    }
    return
}

//获取所有文件和文件夹的路径，包含子文件夹下的文件和文件夹
// @param ext 过滤文件，只获取匹配后缀名的文件，示例：.go
func GetAllPaths(dirPath string, exts ...string) (dirPaths []string, filePaths []string, err error) {
    dirPaths, filePaths, err = GetPaths(dirPath, exts...)
    if err != nil {
        return
    }

    // 读取子文件夹下文件和文件夹
    for _, dirPath2 := range dirPaths {
        dirPaths2, filePaths2, err := GetAllPaths(dirPath2, exts...)
        if err != nil {
            return nil, nil, err
        }
        dirPaths = append(dirPaths, dirPaths2...)
        filePaths = append(filePaths, filePaths2...)
    }
    return
}

//获取文件路径中的文件名（包括扩展名）
func Name(filePath string) string {
    _, fileName := filepath.Split(filePath)
    return fileName
}

//获取文件路径中的文件名（不包括扩展名）
func NameNoExt(filePath string) string {
    fileName := Name(filePath)
    ext := filepath.Ext(fileName)
    if ext != "" {
        fileName = strings.TrimSuffix(fileName, ext)
    }
    return fileName
}

//生成新文件名（不包括扩展名）
func NewFileName() string {
    //8+1+6+1+16=32
    return atime.Now().Format("20060102-150405") + "-" + strings.ToLower(arand.RandBase32(10))
}
