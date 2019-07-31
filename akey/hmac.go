package akey

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
)

//hmac-md5单向秘钥key加密 32位
func HmacMd5(src string, key string) string {
	hash := hmac.New(md5.New, []byte(key))
	hash.Write([]byte(src))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

//hmac-md5单向秘钥key加密
func HmacMd5Byte(src []byte, key []byte) []byte {
	hash := hmac.New(md5.New, key)
	hash.Write(src)
	return hash.Sum(nil)
}

//hmac-sha1单向秘钥key加密 40位
func HmacSha1(src string, key string) string {
	hash := hmac.New(sha1.New, []byte(key))
	hash.Write([]byte(src))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

//hmac-sha1单向秘钥key加密
func HmacSha1Byte(src []byte, key []byte) []byte {
	hash := hmac.New(sha1.New, key)
	hash.Write(src)
	return hash.Sum(nil)
}

//hmac-sha256单向秘钥key加密 64位
func HmacSha256(src string, key string) string {
	hash := hmac.New(sha256.New, []byte(key))
	hash.Write([]byte(src))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

//hmac-sha256单向秘钥key加密
func HmacSha256Byte(src []byte, key []byte) []byte {
	hash := hmac.New(sha256.New, key)
	hash.Write(src)
	return hash.Sum(nil)
}
