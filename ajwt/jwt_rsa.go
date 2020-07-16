package ajwt

import (
    "errors"
    "github.com/asktop/gotools/acast"
    "github.com/asktop/gotools/afile"
    "github.com/dgrijalva/jwt-go"
    "io/ioutil"
)

//Jwt加密生成token
//exp 签名过期时间戳，为0不过期
//privateKeyStr 签名加密秘钥
func RsaEncrypt(info map[string]interface{}, exp int64, privateKeyStr string) (token string, err error) {
    privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKeyStr))
    if err != nil {
        return
    }
    //赋值
    mapClaims := make(jwt.MapClaims)
    if exp >= 0 {
        //超时时间
        mapClaims["exp"] = exp
    }
    for k, v := range info {
        mapClaims[k] = v
    }
    //创建 rs256 类型的 token 对象
    tk := jwt.New(jwt.SigningMethodRS256)
    tk.Claims = mapClaims
    token, err = tk.SignedString(privateKey)
    return
}

//Jwt解密token
//publicKeyStr 签名解密公钥
func RsaDecrypt(token string, publicKeyStr string) (info map[string]interface{}, err error) {
    publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKeyStr))
    if err != nil {
        return
    }
    //将token字符串解析成token对象，会自动校验有效性，超时会报错
    tk, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
        return publicKey, nil
    })
    if err != nil {
        return
    }
    return tk.Claims.(jwt.MapClaims), nil
}

//Jwt解密token
//publicKeyStr 签名解密公钥
func RsaDecryptObj(token string, info interface{}, publicKeyStr string) error {
    data, err := RsaDecrypt(token, publicKeyStr)
    if err != nil {
        return err
    } else {
        return acast.MapToStruct(data, info)
    }
}

//Jwt加密生成token
//exp 签名过期时间戳，为0不过期
//privateKeyPath 签名加密秘钥
func RsaPathEncrypt(info map[string]interface{}, exp int64, privateKeyPath string) (token string, err error) {
    if !afile.IsExist(privateKeyPath) {
        err = errors.New("privateKeyPath:" + privateKeyPath + " not exist")
        return
    }
    privateByte, err := ioutil.ReadFile(privateKeyPath)
    if err != nil {
        return
    }
    return RsaEncrypt(info, exp, string(privateByte))
}

//Jwt解密token
//publicKeyPath 签名解密公钥
func RsaPathDecrypt(token string, publicKeyPath string) (info map[string]interface{}, err error) {
    if !afile.IsExist(publicKeyPath) {
        err = errors.New("publicKeyPath:" + publicKeyPath + " not exist")
        return
    }
    publicByte, err := ioutil.ReadFile(publicKeyPath)
    if err != nil {
        return
    }
    return RsaDecrypt(token, string(publicByte))
}

//Jwt解密token
//publicKeyPath 签名解密公钥
func RsaPathDecryptObj(token string, info interface{}, publicKeyPath string) error {
    data, err := RsaPathDecrypt(token, publicKeyPath)
    if err != nil {
        return err
    } else {
        return acast.MapToStruct(data, info)
    }
}
