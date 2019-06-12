package akey

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

//Rsa加密
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

//Rsa解密
func RsaDecrypt(str string, privateKey string) (string, error) {
	if len(str) == 0 {
		return "", errors.New("rsa param can not be empty")
	}
	b, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}

	privatePem, _ := pem.Decode([]byte(privateKey))
	if privatePem == nil {
		return "", errors.New("rsa private key error")
	}
	privateX509, err := x509.ParsePKCS8PrivateKey(privatePem.Bytes)
	if err != nil {
		return "", err
	}
	ret, err := rsa.DecryptPKCS1v15(nil, privateX509.(*rsa.PrivateKey), b)
	if err != nil {
		return "", err
	}
	return string(ret), nil
}
