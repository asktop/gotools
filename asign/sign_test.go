package asign

import (
    "fmt"
    "strconv"
    "testing"
    "time"
)

func TestCheckSignature(t *testing.T) {
    //var appTimestamp int64 = 1564586623
    var appTimestamp int64 = time.Now().UnixNano() / 1e6
    appSecret := "abc"
    signatureMessage := MakeSignatureMessage("GET", "http://127.0.0.1:80/api/v1/systemTime", appTimestamp, nil)
    fmt.Println(signatureMessage) //GEThttp://127.0.0.1:80/api/v1/systemTime1564586623
    appSignature := MakeSignature(signatureMessage, appSecret)
    fmt.Println(appSignature) //79LISK+q8dGf+QXpwBrafnCI8X0=

    err := CheckSignature(appSecret, appSignature, strconv.FormatInt(appTimestamp, 10),
        SignParam{Method: "GET", Host: "127.0.0.1:80", Router: "/api/v1/systemTime", Query: map[string]string{}, ContentType: "application/json", RequestBody: nil},
        SignConfig{Open: true, Timeout: 30000, Scheme: "http"})
    if err != nil {
        fmt.Println("签名验证不通过", err)
    } else {
        fmt.Println("签名验证通过")
    }
}
