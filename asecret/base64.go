package asecret

import "encoding/base64"

//Base64加密
func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

//Base64解密
func Base64Decode(str string) string {
	enbyte, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	}
	return string(enbyte)
}
