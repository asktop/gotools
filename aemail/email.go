package aemail

import (
    "gopkg.in/gomail.v2"
)

type Config struct {
    Host     string `json:"host"`     //SMTP服务器地址:smtp.163.com
    Port     int    `json:"port"`     //SMTP服务器端口
    Email    string `json:"email"`    //发送用户用户名或邮箱
    Password string `json:"password"` //发送用户密码
    Username string `json:"username"` //发送用户显示名
}

type EmailClient struct {
    Host     string //SMTP服务器地址:smtp.163.com
    Port     int    //SMTP服务器端口
    Email    string //发送用户用户名或邮箱
    Password string //发送用户密码
    Username string //发送用户显示名
}

func NewEmailClient(config Config) *EmailClient {
    return &EmailClient{
        Host:     config.Host,
        Port:     config.Port,
        Email:    config.Email,
        Password: config.Password,
        Username: config.Username,
    }
}

func New163EmailClient(email string, password string, username string) *EmailClient {
    return &EmailClient{
        Host:     "smtp.163.com",
        Port:     465,
        Email:    email,
        Password: password,
        Username: username,
    }
}

//发送邮件
func (e *EmailClient) Send(subject string, body string, to ...string) error {
    m := gomail.NewMessage()
    m.SetHeader("From", e.Username+"<"+e.Email+">")
    m.SetHeader("To", to...)
    m.SetHeader("Subject", subject) //设置邮件主题
    m.SetBody("text/html", body)    //设置邮件正文

    d := gomail.NewDialer(e.Host, e.Port, e.Email, e.Password)
    return d.DialAndSend(m)
}
