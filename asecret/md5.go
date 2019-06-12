package asecret

import (
	"crypto/md5"
	"fmt"
)

//md5单向加密 32位
func Md5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

//md5单向加密
func Md5Byte(str string) []byte {
	hash := md5.New()
	hash.Write([]byte(str))
	return hash.Sum(nil)
}
