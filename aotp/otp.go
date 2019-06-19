package aotp

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"github.com/asktop/gotools/arand"
	"github.com/skip2/go-qrcode"
	"time"
)

//OTP（One-Time Password）：OTP动态口令，又称一次性密码（阿里云身份宝、Google Authenticator）

//生成基于时间的OTP密钥、OTP扫描用字符串、OTP扫描二维码的base64加密字符串
func NewOtpSecret(account string) (otpSecret string, otpBody string, otpQrcode string) {
	otpSecret = arand.RandBase32(10)
	//OTP扫描用字符串格式：otpauth://totp/[客户端显示的账户信息]?secret=[secretBase32]
	otpBody = "otpauth://totp/" + account + "?secret=" + otpSecret
	//OTP扫描二维码的base64加密字符串
	qrdata, _ := qrcode.Encode(otpBody, qrcode.Medium, 256)
	otpQrcode = "data:image/png;base64," + base64.StdEncoding.EncodeToString(qrdata)
	return
}

// GetOtpCode 根据OTP密钥 和 当前时间戳timestamp 生成基于时间OTP验证码
func GetOtpCode(otpSecret string) string {
	timestamp := time.Now().Unix()
	hs, err := hmacSha1(otpSecret, timestamp/30)
	if err != nil {
		fmt.Println("GetOtpCode err", err.Error())
		return ""
	}
	d := lastBit4byte(hs)
	return fmt.Sprintf("%06d", d)
}

// GetOtpCode 根据OTP密钥 和 指定时间戳timestamp 生成基于时间OTP验证码
func GetOtpCodeFrom(otpSecret string, timestamp int64) string {
	hs, err := hmacSha1(otpSecret, timestamp/30)
	if err != nil {
		fmt.Println("GetOtpCode err", err.Error())
		return ""
	}
	d := lastBit4byte(hs)
	return fmt.Sprintf("%06d", d)
}

// CheckOtpCode 校验 根据OTP密钥 和 当前时间戳timestamp 生成的基于时间OTP验证码
func CheckOtpCode(otpSecret string, code string) bool {
	otpCode := GetOtpCode(otpSecret)
	return otpCode == code
}

// CheckOtpCodeFrom 校验 根据OTP密钥 和 指定时间戳timestamp 生成的基于时间OTP验证码
func CheckOtpCodeFrom(otpSecret string, timestamp int64, code string) bool {
	otpCode := GetOtpCodeFrom(otpSecret, timestamp)
	return otpCode == code
}

func hmacSha1(secret string, timestamp int64) ([]byte, error) {
	decodeKey, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		return nil, err
	}

	cData := make([]byte, 8)
	binary.BigEndian.PutUint64(cData, uint64(timestamp))

	h1 := hmac.New(sha1.New, decodeKey)
	_, e := h1.Write(cData)
	if e != nil {
		return nil, e
	}
	return h1.Sum(nil), nil
}

func lastBit4byte(hmacSha1 []byte) int32 {
	if len(hmacSha1) != sha1.Size {
		return 0
	}
	offsetBits := int8(hmacSha1[len(hmacSha1)-1]) & 0x0f
	p := (int32(hmacSha1[offsetBits]) << 24) | (int32(hmacSha1[offsetBits+1]) << 16) | (int32(hmacSha1[offsetBits+2]) << 8) | (int32(hmacSha1[offsetBits+3]) << 0)
	snum := p & 0x7fffffff
	return snum % 1000000
}
