package add

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/base64"
    "encoding/json"
    "errors"
    "fmt"
    "github.com/asktop/gotools/aclient"
    "time"
)

type RobotClient struct {
    site        string
    accessToken string
    sign        string
}

//钉钉机器人
func NewRobotClient(accessToken string, sign ...string) *RobotClient {
    c := new(RobotClient)
    c.site = "https://oapi.dingtalk.com"
    c.accessToken = accessToken
    if len(sign) > 0 && sign[0] != "" {
        c.sign = sign[0]
    }
    return c
}

func (c *RobotClient) SendMsg(params map[string]interface{}) (errMsg string, err error) {
    url := fmt.Sprintf("%s/robot/send?access_token=%s", c.site, c.accessToken)
    if c.sign != "" {
        t := time.Now().UnixNano() / 1e6
        sign := sign(t, c.sign)
        url = fmt.Sprintf("%s&timestamp=%d&sign=%s", url, t, sign)
    }
    respBody, _, err := aclient.NewClient().PostJson(url, params)
    if err != nil {
        errMsg = "钉钉机器人消息发送失败"
        return errMsg, err
    }
    msg := robotMsg{}
    err = json.Unmarshal(respBody, &msg)
    if err != nil {
        errMsg = "钉钉机器人消息解析失败"
        return errMsg, err
    }
    if msg.Errcode != 0 || msg.Errmsg != "ok" {
        errMsg = "钉钉机器人消息发送失败"
        return errMsg, errors.New(msg.Errmsg)
    }
    return "", nil
}

func sign(t int64, secret string) string {
    strToHash := fmt.Sprintf("%d\n%s", t, secret)
    hmac256 := hmac.New(sha256.New, []byte(secret))
    hmac256.Write([]byte(strToHash))
    data := hmac256.Sum(nil)
    return base64.StdEncoding.EncodeToString(data)
}

type robotMsg struct {
    Errcode int    `json:"errcode"`
    Errmsg  string `json:"errmsg"`
}

func (c *RobotClient) SendTextMsg(content string, isAtAll bool, atMobiles ...string) (errMsg string, err error) {
    params := map[string]interface{}{}
    params["msgtype"] = "text"
    params["text"] = map[string]string{
        "content": content,
    }
    params["at"] = map[string]interface{}{
        "atMobiles": atMobiles,
        "isAtAll":   isAtAll,
    }
    return c.SendMsg(params)
}

func (c *RobotClient) SendMarkdownMsg(title string, text string, isAtAll bool, atMobiles ...string) (errMsg string, err error) {
    params := map[string]interface{}{}
    params["msgtype"] = "markdown"
    params["markdown"] = map[string]string{
        "title": title,
        "text":  text,
    }
    params["at"] = map[string]interface{}{
        "atMobiles": atMobiles,
        "isAtAll":   isAtAll,
    }
    return c.SendMsg(params)
}

func (c *RobotClient) SendLinkMsg(title string, text string, messageUrl string, picUrl string, isAtAll bool, atMobiles ...string) (errMsg string, err error) {
    params := map[string]interface{}{}
    params["msgtype"] = "link"
    params["link"] = map[string]string{
        "title":      title,
        "text":       text,
        "messageUrl": messageUrl,
        "picUrl":     picUrl,
    }
    params["at"] = map[string]interface{}{
        "atMobiles": atMobiles,
        "isAtAll":   isAtAll,
    }
    return c.SendMsg(params)
}
