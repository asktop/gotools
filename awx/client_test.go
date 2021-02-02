package awx

import (
    "fmt"
    "github.com/asktop/gotools/atime"
    "testing"
)

func TestWeChatSendTmpl(t *testing.T) {
    appid := "wx87846d66df38cf51"
    appsecret := ""
    openid := "oenWX09-JbxytZz0833PlFI7ym10"
    tmplId := "5VPlifIQDHFQGYIabV-95FJhYXAhT_5JueBFl98HtCw"
    params := map[string]string{}
    params["first"] = "恭喜"
    params["keyword1"] = "测试抽奖"
    params["keyword2"] = "恭喜中奖"
    params["keyword3"] = "尽快领取"
    params["keyword4"] = "过期作废"
    params["remark"] = atime.Now().Format(atime.DATETIME)

    client := NewClient(appid, appsecret)
    errMsg, err := client.SendTmplMsg(openid, tmplId, params)
    if err != nil {
        fmt.Println(err, errMsg)
    } else {
        fmt.Println("ok")
    }
}

func TestWeChatSendCustom(t *testing.T) {
    appid := "wx87846d66df38cf51"
    appsecret := ""
    openid := "oenWX09-JbxytZz0833PlFI7ym10"
    tmplContent := `hello world`
    //tmplContent := `开奖时间:{{date3.DATA}}\n活动名称{{thing5.DATA}}\n开奖结果{{thing6.DATA}}\n温馨提示{{thing7.DATA}}`
    params := map[string]string{}
    //params["date3"] = atime.Now().Format(atime.DATETIME)
    //params["thing5"] = "测试抽奖"
    //params["thing6"] = "恭喜中奖"
    //params["thing7"] = "尽快领取"
    client := NewClient(appid, appsecret)
    errMsg, err := client.SendCustomMsg(openid, tmplContent, params)
    if err != nil {
        fmt.Println(err, errMsg)
    } else {
        fmt.Println("ok")
    }
}
