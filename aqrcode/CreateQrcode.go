package aqrcode

import (
    "encoding/base64"
    "github.com/skip2/go-qrcode"
    "html/template"
)

type QrcodeInfo struct {
    QrcodeData   []byte       //二维码内容（生成图片用）
    QrcodeBase64 template.URL //二维码base64（页面直接显示用）
}

//创建二维码
func CreateQrcode(content string, size ...int) (*QrcodeInfo, error) {
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
    return info, err
}
