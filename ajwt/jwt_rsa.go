package ajwt

import (
    "crypto/rsa"
    "errors"
    "github.com/asktop/gotools/acast"
    "github.com/asktop/gotools/afile"
    "github.com/asktop/gotools/atime"
    "github.com/dgrijalva/jwt-go"
    "io/ioutil"
    "time"
)

var (
    defaultPrivateKey = `-----BEGIN PRIVATE KEY-----
MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBALl1vH9ljRFusOAw
kV3A05DmItpLubsmMYJZjAHlEc9Kn16Yo0Dga8JeE2sagrEDsWyu5A+Fiq9IInaF
T0PJdyVcnR0TKPWj1kFop/qEX44vwj/+V2NOxYaJ4Cb1o5x//yjJ2bBa/8M0nVx+
UxV1oc7ZYOeP07BO/6sMEhFoJ2H/AgMBAAECgYB7BmD+WY0UrUrjzRQBDzLJAgDI
skcIoLNi9qfrcds4mRXTGInjNXwGOYXEHJfpeLuvjux2Z22yDLXfzVrharl/jDNf
K/bJGVhfiwV5x+yS+u1TXF86aSB/thcqleFXvRGIV3WuX7um+q7J8oA/OUPGxOyB
FjC8HltCWFRPm1az8QJBANsei0poDiq0vY0b3WrLetxO9Gp00D/fzBgIfbsw874s
UUWc3LHF/kLGSOvkvue93cnggCMeuG+i8Y2QalVz7/sCQQDYrN6BJhBNxUJnXBTr
F4bZ3Hms3GYLN7E+RuFpNB/wgzVN0LOkdKyogVLCt1inSa8szrLOXTcA3zMVfFdt
HMLNAkEAj6JmDFBJeRUha+5oJilcUC4xaddI65X4Y4itYpekL3U9kTRSNvZixcLU
6kz4F1EOodbYKC1rGULmtLWF/p4RIQJBAJ/KJLEjrAReg9ELxFV3XTiPcp/7Tbna
EXk29ocKLL/HU2kWj1Spwqbl8G2uns+H9IrbyFuNvMGE2PxwXV0XR8UCQAkKuKJL
TR0SqEpT1VKqpIauusMP5Jqfysvf/iR+nyS3zSId07+JZE9+uLgJOnoHVHRZNLB1
NWlpbP7I2GirvPo=
-----END PRIVATE KEY-----`
    defaultPublicKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC5dbx/ZY0RbrDgMJFdwNOQ5iLa
S7m7JjGCWYwB5RHPSp9emKNA4GvCXhNrGoKxA7FsruQPhYqvSCJ2hU9DyXclXJ0d
Eyj1o9ZBaKf6hF+OL8I//ldjTsWGieAm9aOcf/8oydmwWv/DNJ1cflMVdaHO2WDn
j9OwTv+rDBIRaCdh/wIDAQAB
-----END PUBLIC KEY-----`
    defaultRsaJwt, _ = NewRsaJwt([]byte(defaultPrivateKey), []byte(defaultPublicKey), 60*60*24)
)

type RsaJwt struct {
    privateKey *rsa.PrivateKey //私钥
    publicKey  *rsa.PublicKey  //公钥
    expiresAt  int             //签名过期时间
}

func NewRsaJwt(privateKeyByte []byte, publicKeyByte []byte, expiresAt int) (*RsaJwt, error) {
    privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyByte)
    if err != nil {
        return nil, err
    }
    publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyByte)
    if err != nil {
        return nil, err
    }
    return &RsaJwt{privateKey: privateKey, publicKey: publicKey, expiresAt: expiresAt}, nil
}

func NewRsaJwtPath(privateKeyPath string, publicKeyPath string, expiresAt int) (*RsaJwt, error) {
    if !afile.IsExist(privateKeyPath) {
        return nil, errors.New("privateKeyPath:" + privateKeyPath + " not exist")
    }
    if !afile.IsExist(publicKeyPath) {
        return nil, errors.New("publicKeyPath:" + publicKeyPath + " not exist")
    }
    privateByte, err := ioutil.ReadFile(privateKeyPath)
    if err != nil {
        return nil, err
    }
    publicByte, err := ioutil.ReadFile(publicKeyPath)
    if err != nil {
        return nil, err
    }
    return NewRsaJwt(privateByte, publicByte, expiresAt)
}

//将 info 信息生成 token 字符串
func (j *RsaJwt) NewToken(info map[string]interface{}) (token string, expiresAt int64, err error) {
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
    token, err = tk.SignedString(j.privateKey)
    return
}

//解析token字符串，获取info
func (j *RsaJwt) GetInfo(token string) (info map[string]interface{}, err error) {
    //将token字符串解析成token对象，会自动校验有效性，超时会报错
    tk, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
        return j.publicKey, nil
    })
    if err != nil {
        return
    }
    return tk.Claims.(jwt.MapClaims), nil
}

//解析token字符串，获取info
func (j *RsaJwt) GetInfoObj(token string, info interface{}) error {
    data, err := j.GetInfo(token)
    if err != nil {
        return err
    } else {
        return acast.MapToStruct(data, info)
    }
}

func DefaultRsaJwt(privateKeyByte []byte, publicKeyByte []byte, expiresAt int) {
    defaultRsaJwt, _ = NewRsaJwt(privateKeyByte, publicKeyByte, expiresAt)
}

func DefaultRsaJwtPath(privateKeyPath string, publicKeyPath string, expiresAt int) {
    defaultRsaJwt, _ = NewRsaJwtPath(privateKeyPath, publicKeyPath, expiresAt)
}

//将 info 信息生成 token 字符串
func NewRsaToken(info map[string]interface{}) (token string, expiresAt int64, err error) {
    return defaultRsaJwt.NewToken(info)
}

//解析token字符串，获取info
func GetRsaInfo(token string) (info map[string]interface{}, err error) {
    return defaultRsaJwt.GetInfo(token)
}

//解析token字符串，获取info
func GetRsaInfoObj(token string, info interface{}) error {
    return defaultRsaJwt.GetInfoObj(token, info)
}
