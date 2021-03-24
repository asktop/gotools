package context

import (
    "encoding/json"
    "fmt"
    "time"

    "github.com/asktop/gotools/aqywx/util"
)

const (
    // accessTokenURL 获取access_token的接口
    accessTokenURL = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s"
)

// AccessTokenResp struct
type AccessTokenResp struct {
    util.CommonError

    AccessToken string `json:"access_token"`
    ExpiresIn   int64  `json:"expires_in"`
}

func (ctx *Context) GetAccessTokenCacheKey() string {
    return fmt.Sprintf("qywx:access_token:%s:%s", ctx.CorpID, ctx.CorpSecret)
}

//缓存微信access_token
func (ctx *Context) SetAccessToken(accessToken string, timeout ...int64) error {
    ctx.AccessTokenLock.Lock()
    defer ctx.AccessTokenLock.Unlock()

    var expires int64
    if len(timeout) > 0 && timeout[0] > 0 {
        expires = timeout[0]
    } else {
        expires = 7000
    }

    return ctx.Cache.Set(ctx.GetAccessTokenCacheKey(), accessToken, time.Duration(expires)*time.Second)
}

// GetAccessToken 获取access_token
func (ctx *Context) GetAccessToken(force ...bool) (accessToken string, err error) {
    ctx.AccessTokenLock.Lock()
    defer ctx.AccessTokenLock.Unlock()

    if !(len(force) > 0 && force[0]) {
        val := ctx.Cache.Get(ctx.GetAccessTokenCacheKey())
        if val != nil {
            accessToken = val.(string)
            return
        }
    }

    //从微信服务器获取
    var accessTokenResp AccessTokenResp
    accessTokenResp, err = ctx.GetAccessTokenFromServer()
    if err != nil {
        return
    }

    accessToken = accessTokenResp.AccessToken
    expires := accessTokenResp.ExpiresIn - 1500
    err = ctx.Cache.Set(ctx.GetAccessTokenCacheKey(), accessTokenResp.AccessToken, time.Duration(expires)*time.Second)
    return
}

// GetAccessTokenFromServer 强制从企业微信服务器获取token
func (ctx *Context) GetAccessTokenFromServer() (accessTokenResp AccessTokenResp, err error) {
    //log.Printf("GetAccessTokenFromServer")
    url := fmt.Sprintf(accessTokenURL, ctx.CorpID, ctx.CorpSecret)
    var body []byte
    body, err = util.HTTPGet(url)
    if err != nil {
        return
    }
    err = json.Unmarshal(body, &accessTokenResp)
    if err != nil {
        return
    }
    if accessTokenResp.ErrCode != 0 {
        err = fmt.Errorf("get work wechat access_token error : errcode=%v , errormsg=%v", accessTokenResp.ErrCode, accessTokenResp.ErrMsg)
        return
    }
    return
}
