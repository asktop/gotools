package tx

import (
    "errors"
    "github.com/qichengzx/qcloudsms_go"
    "strconv"
)

type TxClient struct {
    appId  string
    appKey string
}

func NewClient(id string, key string) *TxClient {
    txClient := &TxClient{}
    txClient.appId = id
    txClient.appKey = key
    return txClient
}

func (c *TxClient) SendSms(phoneNumber, signName, tplCode string, tplParams []string) (msg string, err error) {
    if len(c.appId) == 0 {
        err = errors.New("AccessKeyId required")
        return
    }
    if len(c.appKey) == 0 {
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
    tplId, err := strconv.Atoi(tplCode)
    if err != nil {
        return
    }
    //参数
    var client = qcloudsms.NewClient(qcloudsms.NewOptions(c.appId, c.appKey, ""))
    req := qcloudsms.SMSSingleReq{
        Tel:   qcloudsms.SMSTel{Nationcode: "86", Mobile: phoneNumber},
        Sign:  signName,
        TplID: tplId,
    }
    if tplParams != nil {
        req.Params = tplParams
    }
    ok, err := client.SendSMSSingle(req)
    if ok {
        return "", nil
    } else {
        return "", err
    }
}
