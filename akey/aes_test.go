package akey

import (
	"github.com/asktop/gotools/atest"
	"testing"
)

var (
	content       = []byte("pibigstar")
	content_16    = Base64Decode("v1jqsGHId/H8onlVHR8Vaw==")
	content_24    = Base64Decode("0TXOaj5KMoLhNWmJ3lxY1A==")
	content_32    = Base64Decode("qM/Waw1kkWhrwzek24rCSA==")
	content_16_iv = Base64Decode("DqQUXiHgW/XFb6Qs98+hrA==")
	content_32_iv = Base64Decode("ZuLgAOii+lrD5KJoQ7yQ8Q==")
	// iv 长度必须等于blockSize，只能为16
	iv         = []byte("Hello My GoFrame")
	key_16     = []byte("1234567891234567")
	key_17     = []byte("12345678912345670")
	key_24     = []byte("123456789123456789123456")
	key_32     = []byte("12345678912345678912345678912345")
	keys       = []byte("12345678912345678912345678912346")
	key_err    = []byte("1234")
	key_32_err = []byte("1234567891234567891234567891234 ")
)

func TestEncrypt(t *testing.T) {
	atest.Case(t, func() {
		data, err := AesEncrypt(content, key_16)
		atest.Assert(err, nil)
		atest.Assert(data, []byte(content_16))
		data, err = AesEncrypt(content, key_24)
		atest.Assert(err, nil)
		atest.Assert(data, []byte(content_24))
		data, err = AesEncrypt(content, key_32)
		atest.Assert(err, nil)
		atest.Assert(data, []byte(content_32))
		data, err = AesEncrypt(content, key_16, iv)
		atest.Assert(err, nil)
		atest.Assert(data, []byte(content_16_iv))
		data, err = AesEncrypt(content, key_32, iv)
		atest.Assert(err, nil)
		atest.Assert(data, []byte(content_32_iv))
	})
}

func TestDecrypt(t *testing.T) {
	atest.Case(t, func() {
		decrypt, err := AesDecrypt([]byte(content_16), key_16)
		atest.Assert(err, nil)
		atest.Assert(decrypt, content)

		decrypt, err = AesDecrypt([]byte(content_24), key_24)
		atest.Assert(err, nil)
		atest.Assert(decrypt, content)

		decrypt, err = AesDecrypt([]byte(content_32), key_32)
		atest.Assert(err, nil)
		atest.Assert(decrypt, content)

		decrypt, err = AesDecrypt([]byte(content_16_iv), key_16, iv)
		atest.Assert(err, nil)
		atest.Assert(decrypt, content)

		decrypt, err = AesDecrypt([]byte(content_32_iv), key_32, iv)
		atest.Assert(err, nil)
		atest.Assert(decrypt, content)

		decrypt, err = AesDecrypt([]byte(content_32_iv), keys, iv)
		atest.Assert(err, "invalid padding")
	})
}

func TestEncryptErr(t *testing.T) {
	atest.Case(t, func() {
		// encrypt key error
		_, err := AesEncrypt(content, key_err)
		atest.AssertNE(err, nil)
	})
}

func TestDecryptErr(t *testing.T) {
	atest.Case(t, func() {
		// decrypt key error
		encrypt, err := AesEncrypt(content, key_16)
		_, err = AesDecrypt(encrypt, key_err)
		atest.AssertNE(err, nil)

		// decrypt content too short error
		_, err = AesDecrypt([]byte("test"), key_16)
		atest.AssertNE(err, nil)

		// decrypt content size error
		_, err = AesDecrypt(key_17, key_16)
		atest.AssertNE(err, nil)
	})
}

func TestPKCS5UnPaddingErr(t *testing.T) {
	atest.Case(t, func() {
		// AesPKCS5UnPadding blockSize zero
		_, err := AesPKCS5UnPadding(content, 0)
		atest.AssertNE(err, nil)

		// AesPKCS5UnPadding src len zero
		_, err = AesPKCS5UnPadding([]byte(""), 16)
		atest.AssertNE(err, nil)

		// AesPKCS5UnPadding src len > blockSize
		_, err = AesPKCS5UnPadding(key_17, 16)
		atest.AssertNE(err, nil)

		// AesPKCS5UnPadding src len > blockSize
		_, err = AesPKCS5UnPadding(key_32_err, 32)
		atest.AssertNE(err, nil)
	})
}
