package akey

import (
	"encoding/base64"
	"fmt"
)

//Base64加密
func Base64Encode(src []byte) []byte {
	var dst []byte
	base64.StdEncoding.Encode(dst, src)
	return dst
}

//Base64解密
func Base64Decode(src []byte) []byte {
	var dst []byte
	_, err := base64.StdEncoding.Decode(dst, src)
	if err != nil {
		fmt.Println(err)
	}
	return dst
}

//Base64加密
func Base64EncodeToString(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

//Base64解密
func Base64DecodeString(src string) []byte {
	dst, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		fmt.Println(err)
	}
	return dst
}
