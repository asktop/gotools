package akey

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestAes(t *testing.T) {
	src := []byte("Hello World")      // 待加密的数据
	key := []byte("ABCDEFGHABCDEFGH12345678") // 加密的密钥
	fmt.Println("原文：", string(src))

	fmt.Println("------------------ ECB模式 --------------------")
	encrypted, err := AesEncryptECB(src, key)
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
	encrypted, err = AesEncryptCBC(src, key)
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
	encrypted, err = AesEncryptCFB(src, key)
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
