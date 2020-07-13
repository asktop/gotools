package adb

import "fmt"

//mysql配置
type Config struct {
    Host              string `json:"host" yaml:"host"`
    Port              int    `json:"port" yaml:"port"`
    Username          string `json:"username" yaml:"username"`
    Password          string `json:"password" yaml:"password"`
    Database          string `json:"database" yaml:"database"`
    Charset           string `json:"charset" yaml:"charset"`
    AllowOldPasswords int    `json:"allowoldpasswords" yaml:"allowoldpasswords"`
    MaxIdleConns      int    `json:"maxidleconns" yaml:"maxidleconns"`
    MaxOpenConns      int    `json:"maxopenconns" yaml:"maxopenconns"`
    ConnMaxLifetime   int    `json:"connmaxlifetime" yaml:"connmaxlifetime"`
    SqlLogLevel       int    `json:"sqlloglevel" yaml:"sqlloglevel"` //SQL日志打印级别：0：不打印SQL；1：只打印err；2：打印全部
}

//获取mysql连接配置
func (c *Config) GetConfig() string {
    if c.Host == "" {
        c.Host = "127.0.0.1"
    }
    if c.Port == 0 {
        c.Port = 3306
    }
    if c.Charset == "" {
        c.Charset = "utf8mb4"
    }
    if c.AllowOldPasswords == 0 {
        c.AllowOldPasswords = 1
    }
    if c.MaxIdleConns == 0 {
        c.MaxIdleConns = 300
    }
    if c.MaxOpenConns == 0 {
        c.MaxOpenConns = 500
    }
    if c.ConnMaxLifetime == 0 {
        c.ConnMaxLifetime = 180
    }
    return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&allowOldPasswords=%d",
        c.Username,
        c.Password,
        c.Host,
        c.Port,
        c.Database,
        c.Charset,
        c.AllowOldPasswords,
    )
}
