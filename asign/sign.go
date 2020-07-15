package asign

import (
    "bytes"
    "crypto/hmac"
    "crypto/sha1"
    "encoding/base64"
    "encoding/json"
    "errors"
    "fmt"
    "github.com/asktop/gotools/asort"
    "github.com/asktop/gotools/atime"
    "strconv"
    "strings"
)

type SignParam struct {
    Method      string
    Host        string
    Router      string
    Query       map[string]string
    ContentType string
    RequestBody []byte
}

type SignConfig struct {
    Open    bool
    Timeout int64
    Scheme  string
}

//用户签名验证
//open 是否进行用户签名验证
//timeout
func CheckSignature(appSecret string, appSignature string, appTimestamp string, signParam SignParam, config SignConfig) error {
    //是否进行用户签名验证
    if config.Open {
        //签名信息非空验证
        if appSignature == "" || appTimestamp == "" {
            return errors.New("APP-SIGNATURE 或 APP-TIMESTAMP 不能为空")
        }

        //当前系统时间不能与Timestamp差距过大
        serverTime := atime.Now().UnixNano() / 1e6 // 获取当前系统时间（毫秒）
        clientTime, _ := strconv.ParseInt(appTimestamp, 10, 64)
        timeout := config.Timeout
        if timeout <= 0 {
            timeout = 30000
        }
        if serverTime-clientTime >= timeout {
            return errors.New("APP-TIMESTAMP 验证超时")
        }

        //获取请求方法
        httpMethod := signParam.Method

        //获取请求URL
        scheme := config.Scheme
        if strings.HasPrefix(scheme, "https") {
            scheme = "https://"
        } else {
            scheme = "http://"
        }
        host := signParam.Host
        router := signParam.Router
        rawquerys := ""
        querys := signParam.Query
        if len(querys) > 0 {
            rawquerys = "?" + asort.SortParamString(querys, "&")
        }
        fullUrl := scheme + host + router + rawquerys
        //获取POST请求参数
        jsonparams := map[string]interface{}{}
        params := map[string]string{}
        if httpMethod == "POST" {
            if signParam.ContentType != "application/json" {
                return errors.New("Content-Type must be application/json")
            } else if len(signParam.RequestBody) > 0 {
                //json反序列化处理
                decoder := json.NewDecoder(bytes.NewBuffer(signParam.RequestBody))
                decoder.UseNumber()
                decoder.Decode(&jsonparams)
                for k, v := range jsonparams {
                    switch t := v.(type) {
                    case string:
                        params[k] = v.(string)
                    case json.Number:
                        params[k] = v.(json.Number).String()
                    default:
                        _ = t
                        return errors.New("POST param[" + k + "] is unknown type")
                    }
                }
            } else {
                params = nil
            }
        } else {
            params = nil
        }

        //制作签名
        signatureMessage := MakeSignatureMessage(httpMethod, fullUrl, clientTime, params)
        signature := MakeSignature(signatureMessage, appSecret)

        //验证签名
        if appSignature != signature {
            return errors.New("SIGNATURE验证不通过，请核对签名信息：" + signatureMessage)
        }
    }
    return nil
}

//制作签名信息
func MakeSignatureMessage(httpMethod string, reqUrl string, timestamp int64, params map[string]string) string {
    var (
        signatureMessage string
        method           string = strings.ToUpper(httpMethod)
    )
    switch method {
    case "GET":
        signatureMessage = fmt.Sprintf("%s%s%d", method, reqUrl, timestamp)
    case "POST":
        signatureMessage = fmt.Sprintf("%s%s%d%s", method, reqUrl, timestamp, asort.SortParamString(params, "&"))
    }
    return signatureMessage
}

//加密生成签名
func MakeSignature(signatureMessage string, appSecret string) string {
    bmsg := base64.StdEncoding.EncodeToString([]byte(signatureMessage))
    sh := hmac.New(sha1.New, []byte(appSecret))
    sh.Write([]byte(bmsg))
    signature := base64.StdEncoding.EncodeToString(sh.Sum(nil))
    return signature
}
