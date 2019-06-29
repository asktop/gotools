package aotp

import (
	"fmt"
	"testing"
	"time"
)

func TestNewOtpKey(t *testing.T) {
	email := "asktop@qq.com"
	otpSecret, otpBody := NewOtpSecret(email)
	fmt.Println(otpSecret)
	fmt.Println(otpBody)
	fmt.Println(GetOtpCode(otpSecret))
	fmt.Println(GetOtpCodeFrom("NERS2IMGVD2ZJPPG", time.Now().Unix()))
}

func TestGetOtpCode(t *testing.T) {
	fmt.Println(GetOtpCode("NERS2IMGVD2ZJPPG"))
	fmt.Println(GetOtpCodeFrom("NERS2IMGVD2ZJPPG", 1560666562))
}
