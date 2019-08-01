package akey

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

// Rsa加密
/*
```javascript
<script src="/static/common/js/jsencrypt.min.js"></script>
<script>
    var passwd = 'abc';//原始密码
    var encrypt = new JSEncrypt();
    encrypt.setPublicKey($('#rsa_public_key').val());
    var lastpwd = encrypt.encrypt(passwd);//加密密码
</script>
```
*/

// Rsa加密，密钥格式 -----BEGIN PUBLIC KEY-----
func RsaEncrypt(src string, publicKey string) (string, error) {
	if len(src) == 0 {
		return "", errors.New("src can not be empty")
	}
	srcByte := []byte(src)

	// 解密pem格式的公钥
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return "", errors.New("public key error")
	}
	// 解析公钥
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	pubKey := key.(*rsa.PublicKey)

	// rsa加密
	dstByte, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, srcByte)
	if err != nil {
		return "", err
	}

	// 对rsa加密结果进行base64加密
	dst := base64.StdEncoding.EncodeToString(dstByte)
	return dst, nil
}

// Rsa解密，密钥格式 -----BEGIN PRIVATE KEY-----
func RsaDecrypt(src string, privateKey string) (string, error) {
	if len(src) == 0 {
		return "", errors.New("src can not be empty")
	}
	// 对rsa加密结果进行base64解密
	srcByte, _ := base64.StdEncoding.DecodeString(src)

	// 解密pem格式的私钥
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return "", errors.New("private key error")
	}
	// 解析私钥
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	privKey := key.(*rsa.PrivateKey)

	// rsa解密
	dstByte, err := rsa.DecryptPKCS1v15(rand.Reader, privKey, srcByte)
	if err != nil {
		return "", err
	}
	dst := string(dstByte)
	return dst, nil
}

// Rsa加密，密钥格式 -----BEGIN RSA PUBLIC KEY-----
func RsaEncryptPKCS1(src string, publicKey string) (string, error) {
	if len(src) == 0 {
		return "", errors.New("src can not be empty")
	}
	srcByte := []byte(src)

	// 解密pem格式的公钥
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return "", errors.New("public key error")
	}
	// 解析公钥
	key, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return "", err
	}

	// rsa加密
	dstByte, err := rsa.EncryptPKCS1v15(rand.Reader, key, srcByte)
	if err != nil {
		return "", err
	}

	// 对rsa加密结果进行base64加密
	dst := base64.StdEncoding.EncodeToString(dstByte)
	return dst, nil
}

// Rsa解密，密钥格式 -----BEGIN RSA PRIVATE KEY-----
func RsaDecryptPKCS1(src string, privateKey string) (string, error) {
	if len(src) == 0 {
		return "", errors.New("src can not be empty")
	}
	// 对rsa加密结果进行base64解密
	srcByte, _ := base64.StdEncoding.DecodeString(src)

	// 解密pem格式的私钥
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return "", errors.New("private key error")
	}
	// 解析私钥
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	// rsa解密
	dstByte, err := rsa.DecryptPKCS1v15(rand.Reader, key, srcByte)
	if err != nil {
		return "", err
	}
	dst := string(dstByte)
	return dst, nil
}
