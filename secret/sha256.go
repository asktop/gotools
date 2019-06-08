package secret

import (
	"crypto/sha256"
	"fmt"
)

//sha256单向加密 64位
func Sha256(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

//sha256单向加密
func Sha256Byte(str string) []byte {
	hash := sha256.New()
	hash.Write([]byte(str))
	return hash.Sum(nil)
}
