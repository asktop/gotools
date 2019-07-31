package akey

import (
	"crypto/sha256"
	"fmt"
)

//sha256单向加密 64位
func Sha256(src string) string {
	hash := sha256.New()
	hash.Write([]byte(src))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

//sha256单向加密
func Sha256Byte(src []byte) []byte {
	hash := sha256.New()
	hash.Write(src)
	return hash.Sum(nil)
}
