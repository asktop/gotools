package secret

import (
	"crypto/sha1"
	"fmt"
)

//sha1单向加密 40位
func Sha1(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

//sha1单向加密
func Sha1Byte(str string) []byte {
	hash := sha1.New()
	hash.Write([]byte(str))
	return hash.Sum(nil)
}
