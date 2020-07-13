package ajwt

import (
    "github.com/asktop/gotools/acast"
    "github.com/asktop/gotools/atime"
    "github.com/dgrijalva/jwt-go"
    "time"
)

//JWT : Json Web Token
//规则：36位base64( header的json串 )+"."+base64( claims的json串 )+"."+43位加密算法秘钥加密生成的signature签名( 36位base64( header的json串 )+"."+base64( claims的json串 ) )
//示例：eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjAwODg4MjEsInVzZXJpZCI6MTIzfQ.N5HT1gpwA2tXip9V9-47iwd9fWwHAY5waUZVKleMIkQ

//Header 组成部分：头
//type Header struct {
//	typ string //JWT规范：JWT
//	alg string //签名加密算法：HS256
//}

//Claims 组成部分：有效载荷
//type Claims struct {
//	Id        string `json:"jti,omitempty"` //id
//	Subject   string `json:"sub,omitempty"` //主题
//	Audience  string `json:"aud,omitempty"` //用户
//	Issuer    string `json:"iss,omitempty"` //发行者
//	IssuedAt  int64  `json:"iat,omitempty"` //发行时间
//	NotBefore int64  `json:"nbf,omitempty"` //生效时间
//	ExpiresAt int64  `json:"exp,omitempty"` //过期时间
//}

var defaultJwt = NewJwt("asktop", 60*60*24)

type Jwt struct {
    secretKey string //签名加密秘钥
    expiresAt int    //签名过期时间
}

func NewJwt(secretkey string, expiresAt int) *Jwt {
    return &Jwt{secretKey: secretkey, expiresAt: expiresAt}
}

//将 info 信息生成 token 字符串
func (j *Jwt) NewToken(info map[string]interface{}) (token string, expiresAt int64, err error) {
    expiresAt = atime.Now().Add(time.Second * time.Duration(j.expiresAt)).Unix() //超时时间
    //创建 hs256 类型的 token 对象
    tk := jwt.New(jwt.SigningMethodHS256)
    //赋值，并设置超时时间
    mapClaims := make(jwt.MapClaims)
    mapClaims["exp"] = expiresAt //超时时间
    for k, v := range info {
        mapClaims[k] = v
    }
    tk.Claims = mapClaims
    token, err = tk.SignedString([]byte(j.secretKey))
    return
}

//解析token字符串，获取info
func (j *Jwt) GetInfo(token string) (info map[string]interface{}, err error) {
    //将token字符串解析成token对象，会自动校验有效性，超时会报错
    tk, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
        return []byte(j.secretKey), nil
    })
    if err != nil {
        return
    }
    return tk.Claims.(jwt.MapClaims), nil
}

//解析token字符串，获取info
func (j *Jwt) GetInfoObj(token string, info interface{}) error {
    data, err := j.GetInfo(token)
    if err != nil {
        return err
    } else {
        return acast.MapToStruct(data, info)
    }
}

func SetDefaultJwt(secretkey string, expiresAt int) {
    defaultJwt = NewJwt(secretkey, expiresAt)
}

//将 info 信息生成 token 字符串
func NewToken(info map[string]interface{}) (token string, expiresAt int64, err error) {
    return defaultJwt.NewToken(info)
}

//解析token字符串，获取info
func GetInfo(token string) (info map[string]interface{}, err error) {
    return defaultJwt.GetInfo(token)
}

//解析token字符串，获取info
func GetInfoObj(token string, info interface{}) error {
    return defaultJwt.GetInfoObj(token, info)
}
