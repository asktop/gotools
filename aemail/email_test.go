package aemail

import (
    "fmt"
    "testing"
)

func TestSendEmail(t *testing.T) {
    url := "http://www.baidu.com"
    body := `<a href="` + url + `"> 点我继续 </a><br><br>若以上链接无法点击，请复制下面内容到浏览器地址栏并访问<br>` + url
    err := New163EmailClient(Config{Email: "asktop2010@163.com", Password: "", Name: "asktop"}).Send("测试123", body, "asktop2010@163.com")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("发送成功")
    }
}
