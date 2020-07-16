package aclient

import (
	"bytes"
	"crypto/tls"
	"fmt"
	rpcjson "github.com/gorilla/rpc/json"
	"net/http"
	"time"
)

type TLSCert struct {
	CertPath string
	KeyPath  string
}

type JsonRpcClient struct {
	url    string
	server string
	client http.Client
}

//JsonRpc客户端
//import rpcjson "github.com/gorilla/rpc/json"
func NewJsonRpcClient(url string, server string, tlsCert ...TLSCert) *JsonRpcClient {
	jsonRpcClient := new(JsonRpcClient)
	jsonRpcClient.url = url
	jsonRpcClient.server = server
	var tlsConfig *tls.Config
	if len(tlsCert) > 0 {
		// 加载双向认证证书
		cert, _ := tls.LoadX509KeyPair(tlsCert[0].CertPath, tlsCert[0].KeyPath)
		tlsConfig = &tls.Config{
			InsecureSkipVerify: true,
			Certificates:       []tls.Certificate{cert},
		}
	} else {
		// 不加载双向认证证书
		tlsConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}
	//创建HttpClient并发起请求
	client := http.Client{
		Transport: &http.Transport{
			DisableKeepAlives:   true, //true:不同HTTP请求之间TCP连接的重用将被阻止（http1.1默认为长连接，此处改为短连接）
			MaxIdleConnsPerHost: 512,  //控制每个主机下的最大闲置连接数目
			TLSClientConfig:     tlsConfig,
			TLSHandshakeTimeout: time.Second * 10,
		},
		Timeout: time.Minute * 10, //Client请求的时间限制,该超时限制包括连接时间、重定向和读取response body时间;Timeout为零值表示不设置超时
	}
	jsonRpcClient.client = client
	return jsonRpcClient
}

func (c *JsonRpcClient) Call(method string, in interface{}, out interface{}) error {
	//组装JsonRpc请求信息
	body, err := rpcjson.EncodeClientRequest(c.server+"."+method, in)
	if err != nil {
		return fmt.Errorf("组装JsonRpc请求信息失败:%s", err.Error())
	}
	//fmt.Println(string(body))
	//{"method":"JsonRpc.Demo","params":[{"id":123,"name":"abc","age":18}],"id":5577006791947779410}

	//创建HttpRequest
	req, err := http.NewRequest("POST", c.url, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("创建HttpRequest失败:%s", err.Error())
	}
	//req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//解析JsonRpc响应信息
	err = rpcjson.DecodeClientResponse(resp.Body, out)
	if err != nil {
		return err
	}
	return nil
}
