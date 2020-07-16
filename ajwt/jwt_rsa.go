package ajwt

import (
    "errors"
    "github.com/asktop/gotools/acast"
    "github.com/asktop/gotools/afile"
    "github.com/asktop/gotools/atime"
    "github.com/dgrijalva/jwt-go"
    "io/ioutil"
    "time"
)

//Jwt加密生成token
//secretKey 签名加密秘钥
//expiresAt 签名过期时长
func RsaEncrypt(info map[string]interface{}, expiresAt int, privateKeyStr string) (token string, exp int64, err error) {
    privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKeyStr))
    if err != nil {
        return
    }
    //赋值
    mapClaims := make(jwt.MapClaims)
    if exp >= 0 {
        exp = atime.Now().Add(time.Second * time.Duration(expiresAt)).Unix() //超时时间
        mapClaims["exp"] = exp                                               //超时时间
    }
    for k, v := range info {
        mapClaims[k] = v
    }
    //创建 hs256 类型的 token 对象
    tk := jwt.New(jwt.SigningMethodHS256)
    tk.Claims = mapClaims
    token, err = tk.SignedString(privateKey)
    return
}

//Jwt解密token
//secretKey 签名加密秘钥
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
//secretKey 签名加密秘钥
func RsaDecryptObj(token string, info interface{}, publicKeyStr string) error {
    data, err := RsaDecrypt(token, publicKeyStr)
    if err != nil {
        return err
    } else {
        return acast.MapToStruct(data, info)
    }
}

//Jwt加密生成token
//secretKey 签名加密秘钥
//expiresAt 签名过期时长
func RsaPathEncrypt(info map[string]interface{}, expiresAt int, privateKeyPath string) (token string, exp int64, err error) {
    if !afile.IsExist(privateKeyPath) {
        err = errors.New("privateKeyPath:" + privateKeyPath + " not exist")
        return
    }
    privateByte, err := ioutil.ReadFile(privateKeyPath)
    if err != nil {
        return
    }
    privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateByte)
    if err != nil {
        return
    }
    //赋值
    mapClaims := make(jwt.MapClaims)
    if exp >= 0 {
        exp = atime.Now().Add(time.Second * time.Duration(expiresAt)).Unix() //超时时间
        mapClaims["exp"] = exp                                               //超时时间
    }
    for k, v := range info {
        mapClaims[k] = v
    }
    //创建 hs256 类型的 token 对象
    tk := jwt.New(jwt.SigningMethodHS256)
    tk.Claims = mapClaims
    token, err = tk.SignedString(privateKey)
    return
}

//Jwt解密token
//secretKey 签名加密秘钥
func RsaPathDecrypt(token string, publicKeyPath string) (info map[string]interface{}, err error) {
    if !afile.IsExist(publicKeyPath) {
        err = errors.New("publicKeyPath:" + publicKeyPath + " not exist")
        return
    }
    publicByte, err := ioutil.ReadFile(publicKeyPath)
    if err != nil {
        return
    }
    publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicByte)
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
//secretKey 签名加密秘钥
func RsaPathDecryptObj(token string, info interface{}, publicKeyPath string) error {
    data, err := RsaPathDecrypt(token, publicKeyPath)
    if err != nil {
        return err
    } else {
        return acast.MapToStruct(data, info)
    }
}
