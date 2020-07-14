package aqrcode

import (
    "encoding/base64"
    "github.com/asktop/gotools/acast"
    "github.com/asktop/gotools/akey"
    "github.com/asktop/gotools/aupload"
    "github.com/skip2/go-qrcode"
    "html/template"
    "path"
    "strings"
)

type QrcodeInfo struct {
    QrcodeData   []byte       //二维码内容
    QrcodeBase64 template.URL //二维码base64
    Url          string       //二维码url
    FilePathName string       //二维码本地存储相对路径
}

//创建二维码
func CreateQrcode(content string, filePath string, size ...int) (*QrcodeInfo, error) {
    filePath = strings.Trim(strings.TrimSpace(filePath), "/")
    if filePath == "" {
        filePath = "qrcode"
    }
    var siz int
    if len(size) > 0 {
        siz = size[0]
    } else {
        siz = 200
    }
    qrcodeData, err := qrcode.Encode(content, 2, siz)
    if err != nil {
        return nil, err
    }

    info := new(QrcodeInfo)
    info.QrcodeData = qrcodeData
    info.QrcodeBase64 = template.URL("data:image/png;base64," + base64.StdEncoding.EncodeToString(qrcodeData))
    fileName := akey.Md5(content) + "_" + acast.ToString(siz) + ".png"
    filePathName := path.Join(filePath, fileName)

    data, err := aupload.UploadFromByte(qrcodeData, filePathName)
    if err != nil {
        return nil, err
    }
    info.Url = data.Url
    info.FilePathName = filePathName
    return info, err
}
