package asms

import (
    "errors"
    "github.com/asktop/gotools/asms/ali"
    "github.com/asktop/gotools/asms/tx"
)

const (
    Driver_Ali Driver = "ali"
    Driver_Tx  Driver = "tx"
)

type Driver string

type Client struct {
    driver string
    id     string
    key    string
}

func NewClient(driver Driver, id string, key string) *Client {
    client := &Client{}
    client.driver = string(driver)
    client.id = id
    client.key = key
    return client
}

func (c *Client) SendSms(phoneNumber, signName, tplCode string, tplParams interface{}) (msg string, err error) {
    switch c.driver {
    case "ali":
        params, ok := tplParams.(map[string]string)
        if ok {
            return ali.NewClient(c.id, c.key).SendSms(phoneNumber, signName, tplCode, params)
        } else {
            return "", errors.New("sms_ali tplParams类型必须为map[string]string")
        }
    case "tx":
        params, ok := tplParams.([]string)
        if ok {
            return tx.NewClient(c.id, c.key).SendSms(phoneNumber, signName, tplCode, params)
        } else {
            return "", errors.New("sms_tx tplParams类型必须为[]string")
        }
    default:
        return "", errors.New("driver 不能为空")
    }
}
