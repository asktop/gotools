package akey

import (
	"encoding/hex"
	"fmt"
)

//Hex加密
func HexEncode(src []byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return dst
}

//Hex解密
func HexDecode(src []byte) []byte {
	n, err := hex.Decode(src, src)
	if err != nil {
		fmt.Println(err)
	}
	return src[:n]
}

//Hex加密
func HexEncodeToString(src []byte) string {
	return hex.EncodeToString(src)
}

//Hex解密
func HexDecodeString(src string) []byte {
	dst, err := hex.DecodeString(src)
	if err != nil {
		fmt.Println(err)
	}
	return dst
}
