package arsa

import (
    "encoding/base64"
    "io/ioutil"
)

// 公钥加密
func PublicEncrypt(data, publicKey string) (string, error) {

    grsa := RSASecurity{}
    grsa.SetPublicKey(publicKey)

    rsadata, err := grsa.PubKeyENCTYPT([]byte(data))
    if err != nil {
        return "", err
    }

    return base64.StdEncoding.EncodeToString(rsadata), nil
}

// 私钥加密
func PriKeyEncrypt(data, privateKey string) (string, error) {

    grsa := RSASecurity{}
    grsa.SetPrivateKey(privateKey)

    rsadata, err := grsa.PriKeyENCTYPT([]byte(data))
    if err != nil {
        return "", err
    }

    return base64.StdEncoding.EncodeToString(rsadata), nil
}

// 公钥解密
func PublicDecrypt(data, publicKey string) (string, error) {

    databs, _ := base64.StdEncoding.DecodeString(data)

    grsa := RSASecurity{}
    grsa.SetPublicKey(publicKey)

    rsadata, err := grsa.PubKeyDECRYPT([]byte(databs))
    if err != nil {
        return "", err
    }

    return string(rsadata), nil
}

// 私钥解密
func PriKeyDecrypt(data, privateKey string) (string, error) {

    databs, _ := base64.StdEncoding.DecodeString(data)

    grsa := RSASecurity{}
    grsa.SetPrivateKey(privateKey)

    rsadata, err := grsa.PriKeyDECRYPT([]byte(databs))
    if err != nil {
        return "", err
    }

    return string(rsadata), nil
}

// 公钥加密
func PublicEncryptPath(data, publicKeyPath string) (string, error) {
    keys, err := ioutil.ReadFile(publicKeyPath)
    if err != nil {
        return "", err
    }
    return PublicEncrypt(data, string(keys))
}

// 私钥加密
func PriKeyEncryptPath(data, privateKeyPath string) (string, error) {
    keys, err := ioutil.ReadFile(privateKeyPath)
    if err != nil {
        return "", err
    }
    return PriKeyEncrypt(data, string(keys))
}

// 公钥解密
func PublicDecryptPath(data, publicKeyPath string) (string, error) {
    keys, err := ioutil.ReadFile(publicKeyPath)
    if err != nil {
        return "", err
    }
    return PublicDecrypt(data, string(keys))
}

// 私钥解密
func PriKeyDecryptPath(data, privateKeyPath string) (string, error) {
    keys, err := ioutil.ReadFile(privateKeyPath)
    if err != nil {
        return "", err
    }
    return PriKeyDecrypt(data, string(keys))
}
