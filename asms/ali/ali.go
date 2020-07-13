package ali

import (
    "crypto/hmac"
    "crypto/sha1"
    "encoding/base64"
    "encoding/json"
    "errors"
    "github.com/asktop/gotools/aclient"
    "github.com/asktop/gotools/ajson"
    "github.com/asktop/gotools/arand"
    "github.com/asktop/gotools/atime"
    "net/url"
    "sort"
    "strings"
    "time"
)

type AliClient struct {
    gatewayUrl      string
    accessKeyId     string
    accessKeySecret string
}

func NewClient(id string, key string) *AliClient {
    aliClient := &AliClient{}
    aliClient.gatewayUrl = "http://dysmsapi.aliyuncs.com"
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
    if len(tplParams) == 0 {
        err = errors.New("TemplateParam required")
        return
    }
    //参数
    params := make(map[string]string)
    params["AccessKeyId"] = c.accessKeyId
    params["Timestamp"] = atime.Now().In(time.FixedZone("GMT", 0)).Format("2006-01-02T15:04:05Z")
    params["SignatureNonce"] = strings.TrimRight(arand.RandBase32(), "=")
    params["SignatureMethod"] = "HMAC-SHA1"
    params["SignatureVersion"] = "1.0"
    params["Format"] = "json"

    params["Action"] = "SendSms"
    params["Version"] = "2017-05-25"
    params["RegionId"] = "cn-hangzhou"
    params["PhoneNumbers"] = phoneNumber
    params["SignName"] = signName
    params["TemplateCode"] = tplCode
    params["TemplateParam"] = ajson.Encode(tplParams)
    params["SmsUpExtendCode"] = "90999"
    params["OutId"] = "abcdefg"

    //参数排序并加密生成签名
    sortedQueryStr, signature := makeQueryStrAndSignature(params, c.accessKeySecret)
    smsUrl := c.gatewayUrl + "?Signature=" + signature + sortedQueryStr
    //fmt.Println("sortedQueryStr:", sortedQueryStr)
    //fmt.Println("signature:", signature)
    //fmt.Println("smsUrl:", smsUrl)

    //发送短信
    data, _, err := aclient.NewClient().Get(smsUrl, nil)
    if err != nil {
        return
    }
    rs := smsAliResponse{}
    err = json.Unmarshal(data, &rs)
    if err != nil {
        return
    }
    if rs.Code == "OK" {
        return rs.Message, nil
    } else {
        return "", errors.New(rs.Message)
    }
}

type smsAliResponse struct {
    Code      string `json:"Code"`
    Message   string `json:"Message"`
    RequestId string `json:"RequestId"`
    BizId     string `json:"BizId"`
}

//参数排序并加密生成签名
func makeQueryStrAndSignature(params map[string]string, accessKeySecret string) (string, string) {
    //排序参数
    var keys []string
    for key, _ := range params {
        keys = append(keys, key)
    }
    sort.Strings(keys)
    var sortedQueryStr string
    for _, key := range keys {
        paramName := specialUrlEncode(key)
        paramValue := specialUrlEncode(params[key])
        sortedQueryStr = sortedQueryStr + "&" + paramName + "=" + paramValue
    }
    //对排序参数加密生成签名
    signQueryStr := strings.TrimPrefix(sortedQueryStr, "&")
    signQueryStr = "GET" + "&" + specialUrlEncode("/") + "&" + specialUrlEncode(signQueryStr)
    signature := sign(accessKeySecret+"&", signQueryStr)
    signature = specialUrlEncode(signature)
    return sortedQueryStr, signature
}

func specialUrlEncode(value string) string {
    rstValue := url.QueryEscape(value)
    rstValue = strings.Replace(rstValue, "+", "%20", -1)
    rstValue = strings.Replace(rstValue, "*", "%2A", -1)
    rstValue = strings.Replace(rstValue, "%7E", "~", -1)
    return rstValue
}

func sign(accessKeySecret, sortedQueryStr string) string {
    h := hmac.New(sha1.New, []byte(accessKeySecret))
    h.Write([]byte(sortedQueryStr))
    return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
