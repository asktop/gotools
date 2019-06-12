package akey

import (
	"encoding/hex"
	"github.com/asktop/gotools/atest"
	"testing"
)

var (
	errKey     = []byte("1111111111111234123456789")
	errIv      = []byte("123456789")
	errPadding = 5
)

func TestDesECB(t *testing.T) {
	atest.Case(t, func() {
		key := []byte("11111111")
		text := []byte("12345678")
		padding := NOPADDING
		result := "858b176da8b12503"
		// encrypt test
		cipherText, err := DesECBEncrypt(key, text, padding)
		atest.AssertEQ(err, nil)
		atest.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := DesECBDecrypt(key, cipherText, padding)
		atest.AssertEQ(err, nil)
		atest.AssertEQ(string(clearText), "12345678")

		// encrypt err test. when throw exception,the err is not equal nil and the string is nil
		errEncrypt, err := DesECBEncrypt(key, text, errPadding)
		atest.AssertNE(err, nil)
		atest.AssertEQ(errEncrypt, nil)
		errEncrypt, err = DesECBEncrypt(errKey, text, padding)
		atest.AssertNE(err, nil)
		atest.AssertEQ(errEncrypt, nil)
		// err decrypt test.
		errDecrypt, err := DesECBDecrypt(errKey, cipherText, padding)
		atest.AssertNE(err, nil)
		atest.AssertEQ(errDecrypt, nil)
		errDecrypt, err = DesECBDecrypt(key, cipherText, errPadding)
		atest.AssertNE(err, nil)
		atest.AssertEQ(errDecrypt, nil)
	})

	atest.Case(t, func() {
		key := []byte("11111111")
		text := []byte("12345678")
		padding := PKCS5PADDING
		errPadding := 5
		result := "858b176da8b12503ad6a88b4fa37833d"
		cipherText, err := DesECBEncrypt(key, text, padding)
		atest.AssertEQ(err, nil)
		atest.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := DesECBDecrypt(key, cipherText, padding)
		atest.AssertEQ(err, nil)
		atest.AssertEQ(string(clearText), "12345678")

		// err test
		errEncrypt, err := DesECBEncrypt(key, text, errPadding)
		atest.AssertNE(err, nil)
		atest.AssertEQ(errEncrypt, nil)
		errDecrypt, err := DesECBDecrypt(errKey, cipherText, padding)
		atest.AssertNE(err, nil)
		atest.AssertEQ(errDecrypt, nil)
	})
}

func Test3DesECB(t *testing.T) {
	atest.Case(t, func() {
		key := []byte("1111111111111234")
		text := []byte("1234567812345678")
		padding := NOPADDING
		result := "a23ee24b98c26263a23ee24b98c26263"
		// encrypt test
		cipherText, err := TripleDesECBEncrypt(key, text, padding)
		atest.AssertEQ(err, nil)
		atest.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := TripleDesECBDecrypt(key, cipherText, padding)
		atest.AssertEQ(err, nil)
		atest.AssertEQ(string(clearText), "1234567812345678")
		// err test
		errEncrypt, err := DesECBEncrypt(key, text, errPadding)
		atest.AssertNE(err, nil)
		atest.AssertEQ(errEncrypt, nil)
	})

	atest.Case(t, func() {
		key := []byte("111111111111123412345678")
		text := []byte("123456789")
		padding := PKCS5PADDING
		errPadding := 5
		result := "37989b1effc07a6d00ff89a7d052e79f"
		// encrypt test
		cipherText, err := TripleDesECBEncrypt(key, text, padding)
		atest.AssertEQ(err, nil)
		atest.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := TripleDesECBDecrypt(key, cipherText, padding)
		atest.AssertEQ(err, nil)
		atest.AssertEQ(string(clearText), "123456789")
		// err test, when key is err, but text and padding is right
		errEncrypt, err := TripleDesECBEncrypt(errKey, text, padding)
		atest.AssertNE(err, nil)
		atest.AssertEQ(errEncrypt, nil)
		// when padding is err,but key and text is right
		errEncrypt, err = TripleDesECBEncrypt(key, text, errPadding)
		atest.AssertNE(err, nil)
		atest.AssertEQ(errEncrypt, nil)
		// decrypt err test,when key is err
		errEncrypt, err = TripleDesECBDecrypt(errKey, text, padding)
		atest.AssertNE(err, nil)
		atest.AssertEQ(errEncrypt, nil)
	})
}

func TestDesCBC(t *testing.T) {
	atest.Case(t, func() {
		key := []byte("11111111")
		text := []byte("1234567812345678")
		padding := NOPADDING
		iv := []byte("12345678")
		result := "40826a5800608c87585ca7c9efabee47"
		// encrypt test
		cipherText, err := DesCBCEncrypt(key, text, iv, padding)
		atest.AssertEQ(err, nil)
		atest.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := DesCBCDecrypt(key, cipherText, iv, padding)
		atest.AssertEQ(err, nil)
		atest.AssertEQ(string(clearText), "1234567812345678")
		// encrypt err test.
		errEncrypt, err := DesCBCEncrypt(errKey, text, iv, padding)
		atest.AssertNE(err, nil)
		atest.AssertEQ(errEncrypt, nil)
		// the iv is err
		errEncrypt, err = DesCBCEncrypt(key, text, errIv, padding)
		//atest.AssertNE(err,nil)
		atest.AssertEQ(errEncrypt, nil)
		// the padding is err
		errEncrypt, err = DesCBCEncrypt(key, text, iv, errPadding)
		atest.AssertNE(err, nil)
		atest.AssertEQ(errEncrypt, nil)
		// decrypt err test. the key is err
		errDecrypt, err := DesCBCDecrypt(errKey, cipherText, iv, padding)
		atest.AssertNE(err, nil)
		atest.AssertEQ(errDecrypt, nil)
		// the iv is err
		errDecrypt, err = DesCBCDecrypt(key, cipherText, errIv, padding)
		atest.AssertNE(err, nil)
		atest.AssertEQ(errDecrypt, nil)
		// the padding is err
		errDecrypt, err = DesCBCDecrypt(key, cipherText, iv, errPadding)
		atest.AssertNE(err, nil)
		atest.AssertEQ(errDecrypt, nil)
	})

	atest.Case(t, func() {
		key := []byte("11111111")
		text := []byte("12345678")
		padding := PKCS5PADDING
		iv := []byte("12345678")
		result := "40826a5800608c87100a25d86ac7c52c"
		// encrypt test
		cipherText, err := DesCBCEncrypt(key, text, iv, padding)
		atest.AssertEQ(err, nil)
		atest.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := DesCBCDecrypt(key, cipherText, iv, padding)
		atest.AssertEQ(err, nil)
		atest.AssertEQ(string(clearText), "12345678")
		// err test
		errEncrypt, err := DesCBCEncrypt(key, text, errIv, padding)
		atest.AssertNE(err, nil)
		atest.AssertEQ(errEncrypt, nil)
	})
}

func Test3DesCBC(t *testing.T) {
	atest.Case(t, func() {
		key := []byte("1111111112345678")
		text := []byte("1234567812345678")
		padding := NOPADDING
		iv := []byte("12345678")
		result := "bfde1394e265d5f738d5cab170c77c88"
		// encrypt test
		cipherText, err := TripleDesCBCEncrypt(key, text, iv, padding)
		atest.AssertEQ(err, nil)
		atest.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := TripleDesCBCDecrypt(key, cipherText, iv, padding)
		atest.AssertEQ(err, nil)
		atest.AssertEQ(string(clearText), "1234567812345678")
		// encrypt err test
		errEncrypt, err := TripleDesCBCEncrypt(errKey, text, iv, padding)
		atest.AssertNE(err, nil)
		atest.AssertEQ(errEncrypt, nil)
		// the iv is err
		errEncrypt, err = TripleDesCBCEncrypt(key, text, errIv, padding)
		atest.AssertNE(err, nil)
		atest.AssertEQ(errEncrypt, nil)
		// the padding is err
		errEncrypt, err = TripleDesCBCEncrypt(key, text, iv, errPadding)
		atest.AssertNE(err, nil)
		atest.AssertEQ(errEncrypt, nil)
		// decrypt err test
		errDecrypt, err := TripleDesCBCDecrypt(errKey, cipherText, iv, padding)
		atest.AssertNE(err, nil)
		atest.AssertEQ(errDecrypt, nil)
		// the iv is err
		errDecrypt, err = TripleDesCBCDecrypt(key, cipherText, errIv, padding)
		atest.AssertNE(err, nil)
		atest.AssertEQ(errDecrypt, nil)
		// the padding is err
		errDecrypt, err = TripleDesCBCDecrypt(key, cipherText, iv, errPadding)
		atest.AssertNE(err, nil)
		atest.AssertEQ(errDecrypt, nil)
	})
	atest.Case(t, func() {
		key := []byte("111111111234567812345678")
		text := []byte("12345678")
		padding := PKCS5PADDING
		iv := []byte("12345678")
		result := "40826a5800608c87100a25d86ac7c52c"
		// encrypt test
		cipherText, err := TripleDesCBCEncrypt(key, text, iv, padding)
		atest.AssertEQ(err, nil)
		atest.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := TripleDesCBCDecrypt(key, cipherText, iv, padding)
		atest.AssertEQ(err, nil)
		atest.AssertEQ(string(clearText), "12345678")
	})

}
