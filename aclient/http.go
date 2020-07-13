package aclient

import (
    "bytes"
    "crypto/tls"
    "crypto/x509"
    "encoding/json"
    "fmt"
    "github.com/asktop/gotools/acast"
    "io"
    "io/ioutil"
    "net/http"
    "net/url"
    "strings"
    "time"
)

type Client struct {
    certPath string
    client   http.Client
}

func NewClient(certPath ...string) *Client {
    c := new(Client)

    if len(certPath) > 0 {
        c.certPath = certPath[0]
    }

    var client http.Client
    if c.certPath == "" {
        //创建HttpClient并发起请求
        client = http.Client{
            Transport: &http.Transport{
                DisableKeepAlives:   true, //true:不同HTTP请求之间TCP连接的重用将被阻止（http1.1默认为长连接，此处改为短连接）
                MaxIdleConnsPerHost: 512,  //控制每个主机下的最大闲置连接数目
            },
            Timeout: time.Second * 60, //Client请求的时间限制,该超时限制包括连接时间、重定向和读取response body时间;Timeout为零值表示不设置超时
        }
    } else {
        //获取客户端证书
        crts := x509.NewCertPool()
        crt, err := ioutil.ReadFile(c.certPath)
        if err != nil {
            fmt.Println("获取客户端https证书出错，err:", err)
        }
        crts.AppendCertsFromPEM(crt)

        //创建HttpClient并发起请求
        client = http.Client{
            Transport: &http.Transport{
                DisableKeepAlives:   true,                       //true:不同HTTP请求之间TCP连接的重用将被阻止（http1.1默认为长连接，此处改为短连接）
                MaxIdleConnsPerHost: 512,                        //控制每个主机下的最大闲置连接数目
                TLSClientConfig:     &tls.Config{RootCAs: crts}, //添加证书
            },
            Timeout: time.Second * 60, //Client请求的时间限制,该超时限制包括连接时间、重定向和读取response body时间;Timeout为零值表示不设置超时
        }
    }
    c.client = client
    return c
}

func (c *Client) Request(method string, url string, body io.Reader, headers ...map[string]string) (respBody []byte, statusCode int, err error) {
    //创建HttpRequest
    req, err := http.NewRequest(strings.ToUpper(method), url, body)
    if err != nil {
        return nil, 0, err
    }
    //req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    //req.Header.Set("Content-Type", "application/json")
    if len(headers) > 0 {
        for k, v := range headers[0] {
            req.Header.Set(k, v)
        }
    }

    resp, err := c.client.Do(req)
    if resp != nil {
        statusCode = resp.StatusCode
    }
    if err != nil {
        return nil, statusCode, err
    }
    defer resp.Body.Close()

    //解析响应信息
    reply, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, statusCode, err
    }
    return reply, statusCode, nil
}

func (c *Client) Get(URL string, querys map[string]interface{}, headers ...map[string]string) (respBody []byte, statusCode int, err error) {
    var rawQuery string
    for key, value := range querys {
        rawQuery += fmt.Sprintf("&%s=%v", key, value)
    }
    if rawQuery != "" {
        if strings.Contains(URL, "?") {
            URL += rawQuery
        } else {
            URL += strings.Replace(rawQuery, "&", "?", 1)
        }
    }
    return c.Request(http.MethodGet, URL, nil, headers...)
}

func (c *Client) PostForm(URL string, params map[string]interface{}, headers ...map[string]string) (respBody []byte, statusCode int, err error) {
    header := make(map[string]string)
    if len(headers) > 0 && headers[0] != nil {
        header = headers[0]
    }
    if _, ok := header["Content-Type"]; !ok {
        header["Content-Type"] = "application/x-www-form-urlencoded"
    }
    data := url.Values{}
    for k, v := range params {
        nv, err := acast.ToStringSliceE(v)
        if err != nil {
            return nil, 0, err
        }
        data[k] = nv
    }
    body := strings.NewReader(data.Encode())
    return c.Request(http.MethodPost, URL, body, header)
}

func (c *Client) PostJson(URL string, params map[string]interface{}, headers ...map[string]string) (respBody []byte, statusCode int, err error) {
    header := make(map[string]string)
    if len(headers) > 0 && headers[0] != nil {
        header = headers[0]
    }
    if _, ok := header["Content-Type"]; !ok {
        header["Content-Type"] = "application/json"
    }
    //组装请求信息
    data, _ := json.Marshal(params)
    body := bytes.NewReader(data)
    return c.Request(http.MethodPost, URL, body, header)
}
