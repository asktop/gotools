package akey

import (
	"crypto/sha1"
	"fmt"
)

//sha1单向加密 40位
func Sha1(src string) string {
	hash := sha1.New()
	hash.Write([]byte(src))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

//sha1单向加密
func Sha1Byte(src []byte) []byte {
	hash := sha1.New()
	hash.Write(src)
	return hash.Sum(nil)
}
