package acaptcha

import (
    "fmt"
    "github.com/mojocn/base64Captcha"
    "html/template"
    "strings"
)

var store = base64Captcha.DefaultMemStore

type CaptchaConfig struct {
    Type   string //类型：digit:数字；alpha:字母；string:字符；math:算数
    Height int    //验证码图像高度
    Width  int    //验证码图像宽度
    Length int    //验证码长度
}

//生成验证码
func GetCaptcha(configs ...CaptchaConfig) map[string]interface{} {
    rs := map[string]interface{}{}
    captchaType := "digit"
    height := 40
    width := 100
    length := 4
    if len(configs) > 0 {
        config := configs[0]
        captchaType = config.Type
        if config.Height > 0 {
            height = config.Height
        }
        if config.Width > 0 {
            width = config.Width
        }
        if config.Length > 0 {
            length = config.Length
        }
    }

    var driver base64Captcha.Driver
    switch captchaType {
    case "alpha": //字母
        //noiseCount:验证码复杂度，越大越复杂
        //showLineOptions:背景复杂度，越大越复杂
        driver = base64Captcha.NewDriverString(height, width, 1, 5, length, "abcdefghigklmnopqrstuvwxyz", nil, []string{})
    case "string": //字符
        //noiseCount:验证码复杂度，越大越复杂
        //showLineOptions:背景复杂度，越大越复杂
        driver = base64Captcha.NewDriverString(height, width, 1, 5, length, "acdefhkmnpqrstuvwxyz234578", nil, []string{})
    case "math": //算数
        //noiseCount:验证码复杂度，越大越复杂
        //showLineOptions:背景复杂度，越大越复杂
        driver = base64Captcha.NewDriverMath(height, width, 1, 5, nil, []string{})
    default:
        //maxSkew:验证码复杂度，越大越复杂
        //dotCount:背景复杂度，越大越复杂
        driver = base64Captcha.NewDriverDigit(height, width, length, 0.2, 100)
    }

    id, b64s, err := base64Captcha.NewCaptcha(driver, store).Generate()
    if err != nil {
        fmt.Println(err)
    } else {
        rs["captcha_id"] = id
        rs["captcha_src"] = template.URL(b64s)
    }
    return rs
}

//校验验证码
func VerifyCaptcha(captcha_id string, captcha_code string) bool {
    captcha_code = strings.ToLower(captcha_code)
    return store.Verify(captcha_id, captcha_code, true)
}
