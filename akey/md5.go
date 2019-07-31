package akey

import (
	"crypto/md5"
	"fmt"
)

//md5单向加密 32位
func Md5(src string) string {
	hash := md5.New()
	hash.Write([]byte(src))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

//md5单向加密
func Md5Byte(src []byte) []byte {
	hash := md5.New()
	hash.Write(src)
	return hash.Sum(nil)
}
