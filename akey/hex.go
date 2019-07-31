package akey

import (
	"encoding/hex"
	"fmt"
)

//Hex加密
func HexEncode(src []byte) []byte {
	var dst []byte
	hex.Encode(dst, src)
	return dst
}

//Hex加密
func HexEncodeToString(src []byte) string {
	return hex.EncodeToString(src)
}

//Hex解密
func HexDecode(src []byte) []byte {
	var dst []byte
	_, err := hex.Decode(dst, src)
	if err != nil {
		fmt.Println(err)
	}
	return dst
}

//Hex解密
func HexDecodeString(src string) []byte {
	dst, err := hex.DecodeString(src)
	if err != nil {
		fmt.Println(err)
	}
	return dst
}
