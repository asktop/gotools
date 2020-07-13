package sms

import (
    "fmt"
    "testing"
)

func TestSendSms(t *testing.T) {
    msg, err := NewClient(Driver_Tx, "", "").SendSms("18769910303", "一寸时光", "438595", []string{"9527", "1"})
    if err != nil {
        fmt.Println("短信发送失败", err)
    } else {
        fmt.Println("短信发送成功", msg)
    }
}
