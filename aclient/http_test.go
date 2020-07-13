package aclient

import (
	"fmt"
	"testing"
)

//测试Http客户端
func TestHttp(t *testing.T) {
	url := "http://127.0.0.1:8882/test/json?id=123"
	params := map[string]interface{}{}
	params["name"] = "abc"
	params["account"] = 200.005

	//发起 http json 请求
	//reply, _, err := NewClient().PostForm(url, params)
	reply, _, err := NewClient().PostJson(url, params)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(reply))
}
