package ali

import (
    "errors"
    "github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
    "github.com/asktop/gotools/ajson"
)

type AliClient struct {
    accessKeyId     string
    accessKeySecret string
}

func NewClient(id string, key string) *AliClient {
    aliClient := &AliClient{}
    aliClient.accessKeyId = id
    aliClient.accessKeySecret = key
    return aliClient
}

func (c *AliClient) SendSms(phoneNumber, signName, tplCode string, tplParams map[string]string) (msg string, err error) {
    if len(c.accessKeyId) == 0 {
        err = errors.New("AccessKeyId required")
        return
    }
    if len(c.accessKeySecret) == 0 {
        err = errors.New("AccessKeySecret required")
        return
    }
    if len(phoneNumber) == 0 {
        err = errors.New("PhoneNumbers required")
        return
    }
    if len(signName) == 0 {
        err = errors.New("SignName required")
        return
    }
    if len(tplCode) == 0 {
        err = errors.New("TemplateCode required")
        return
    }

    client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", c.accessKeyId, c.accessKeySecret)
    request := dysmsapi.CreateSendSmsRequest()
    request.Scheme = "https"
    request.PhoneNumbers = phoneNumber
    request.SignName = signName
    request.TemplateCode = tplCode
    if tplParams != nil {
        request.TemplateParam = ajson.Encode(tplParams)
    }

    //发送短信
    rs, e := client.SendSms(request)
    if e != nil {
        return "", e
    }
    if rs.Code == "OK" {
        return rs.Message, nil
    } else {
        return "", errors.New(rs.Message)
    }
}
