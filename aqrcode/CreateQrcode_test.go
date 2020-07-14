package aqrcode

import (
	"fmt"
	"github.com/asktop/gotools/aotp"
	"testing"
)

func TestCreateQrcode(t *testing.T) {
	info, err := CreateQrcode("otpauth://totp/123456789?secret=WLQ52NSJ363HYLGX", "qrcode")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(info.QrcodeBase64)
		fmt.Println(info.Url)
		fmt.Println(info.FilePathName)
	}
}

func TestCreateQrcode2(t *testing.T) {
	_, otp := aotp.NewOtpSecret("123456789")
	fmt.Println(otp)
	info, err := CreateQrcode(otp, "qrcode", 250)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(info.QrcodeBase64)
		fmt.Println(info.Url)
		fmt.Println(info.FilePathName)
	}
}
