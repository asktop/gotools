package acaptcha

//
////获取验证码
//func (c *Base) GetCaptcha() {
//    typ := c.GetString("type")
//    height, _ := c.GetInt("height")
//    width, _ := c.GetInt("width")
//    length, _ := c.GetInt("length")
//    c.Send200("验证码刷新成功", GetCaptcha(CaptchaConfig{Type: typ, Height: height, Width: width, Length: length}))
//}
//
////校验验证码
//func (c *Base) VerifyCaptcha() {
//    captcha_id := c.GetString("captcha_id")
//    captcha_code := c.GetString("captcha_code")
//    if !VerifyCaptcha(captcha_id, captcha_code) {
//        c.Send500("验证码不正确")
//    }
//}
