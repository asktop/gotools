package akey

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestDes(t *testing.T) {
	src := []byte("Hello World")               // 待加密的数据
	key := []byte("ABCDEFGH")                  // 加密的密钥
	key3 := []byte("ABCDEFGHABCDEFGH12345678") // 加密的密钥
	fmt.Println("原文：", string(src))

	fmt.Println("------------------ ECB模式 --------------------")
	encrypted, err := DesEncryptECB(src, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("密文(hex)：", hex.EncodeToString(encrypted))
	fmt.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted, err := DesDecryptECB(encrypted, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("解密结果：", string(decrypted))

	fmt.Println("------------------ ECB模式3DES --------------------")
	encrypted, err = DesEncryptECBTriple(src, key3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("密文(hex)：", hex.EncodeToString(encrypted))
	fmt.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted, err = DesDecryptECBTriple(encrypted, key3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("解密结果：", string(decrypted))

	fmt.Println("------------------ CBC模式 --------------------")
	encrypted, err = DesEncryptCBC(src, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("密文(hex)：", hex.EncodeToString(encrypted))
	fmt.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted, err = DesDecryptCBC(encrypted, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("解密结果：", string(decrypted))

	fmt.Println("------------------ CBC模式3DES --------------------")
	encrypted, err = DesEncryptCBCTriple(src, key3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("密文(hex)：", hex.EncodeToString(encrypted))
	fmt.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted, err = DesDecryptCBCTriple(encrypted, key3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("解密结果：", string(decrypted))
}
