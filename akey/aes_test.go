package akey

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestAes(t *testing.T) {
	str := []byte("Hello World")      // 待加密的数据
	key := []byte("ABCDEFGHIJKLMNOP") // 加密的密钥
	fmt.Println("原文：", string(str))

	fmt.Println("------------------ ECB模式 --------------------")
	encrypted, err := AesEncryptECB(str, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("密文(hex)：", hex.EncodeToString(encrypted))
	fmt.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted, err := AesDecryptECB(encrypted, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("解密结果：", string(decrypted))

	fmt.Println("------------------ CBC模式 --------------------")
	encrypted, err = AesEncryptCBC(str, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("密文(hex)：", hex.EncodeToString(encrypted))
	fmt.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted, err = AesDecryptCBC(encrypted, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("解密结果：", string(decrypted))

	fmt.Println("------------------ CFB模式 --------------------")
	encrypted, err = AesEncryptCFB(str, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("密文(hex)：", hex.EncodeToString(encrypted))
	fmt.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted, err = AesDecryptCFB(encrypted, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("解密结果：", string(decrypted))
}
