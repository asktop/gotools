package add

import (
    "fmt"
    "testing"
)

func TestSendRobotMsg(t *testing.T) {
    accessToken := "813c8d887e18fc4713863fbc4574b6b743a57737de6439509b2595164bdba327"
    sign := "SEC18d408c8b2a4d73e87d9117a9d566b7dde3b520188317f178f68815a8980e209"
    title := "订单成交"
    text := "eos_4.01"
    msg, err := NewRobotClient(accessToken, sign).SendMarkdownMsg(title, text, false)
    if err != nil {
        fmt.Println(msg, err)
    } else {
        fmt.Println("ok")
    }
}
