package akey

import (
	"fmt"
	"testing"
)

var publicKey = `
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCpjjKwzQxDIGqVQhzPtISt+jj1
lQs8OJUbT7cTkqG1fQuYecvkRkhG0cKGbx157ZdfYrk8ATLvT12jQLBPlhb5KUkc
tCtD3ErENKfoZ6m7ph78+ngMZ1y1tY2QYv9xxO8XGVg0GR5NN4z1ZLNpTHbLtjxu
U1VLQKHKtk4RFA0/iwIDAQAB
-----END PUBLIC KEY-----
`

var privateKey = `
-----BEGIN PRIVATE KEY-----
MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAKmOMrDNDEMgapVC
HM+0hK36OPWVCzw4lRtPtxOSobV9C5h5y+RGSEbRwoZvHXntl19iuTwBMu9PXaNA
sE+WFvkpSRy0K0PcSsQ0p+hnqbumHvz6eAxnXLW1jZBi/3HE7xcZWDQZHk03jPVk
s2lMdsu2PG5TVUtAocq2ThEUDT+LAgMBAAECgYAQS+2NVrf1/7iezfLs98HE4wb0
e7XPvPR/4oKLLA3E3tbtec9iCmtJ+0FCII3puS9SaK+7F7Zoj+1FEqOfkqJ0eh+Z
Hbs3kSGrLTnUqhx6hfbbZ7KmOb5ER9YoXjMLiJa6GjrI3JSvTTX4Gq4gtVOmmgKo
Sk1SGw4hDeIRcjqf0QJBAN++0JIbsPfnhp9LC8oeyyh4sJlnXNRdgg35VF2pntLG
4w8MM+tulXz+hoh/yjipiynq+f5FMrUILxLj+Yg7M7kCQQDB/4m7Yx7OZB9E7uQD
tZZKJDsf6jOaITe27HVkVoaYfdjiJdWxgTqOu+XfUlSamDfS8T577ozYe+UJQbjC
J7djAkAdV/Yco1sTOB8Utw/lwyIbvbBTfhXTmCvdT0y8N+mndQQETjJk9wzN5seY
EUJhKgMhTTraGRMsYNVjodhxAGqpAkEAp2hGiKfua5hGy4uCxBitCmndg5rs0LKd
i4LdgqqQF4Nx5yVTFIw55fYLVf2L+KQmeCeqDfQ7Io03hhdqwlDXlQJBAISJqBbK
a/CqrACq9IsQgUaN+OSrVouJWieEYwRYCSHuHFN6nF3q4mOOeQ/qs03NKEL2y5Ad
ZbT6qUiSgz/zHdY=
-----END PRIVATE KEY-----
`

var privateKeyPKCS1 = `
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQCpjjKwzQxDIGqVQhzPtISt+jj1lQs8OJUbT7cTkqG1fQuYecvk
RkhG0cKGbx157ZdfYrk8ATLvT12jQLBPlhb5KUkctCtD3ErENKfoZ6m7ph78+ngM
Z1y1tY2QYv9xxO8XGVg0GR5NN4z1ZLNpTHbLtjxuU1VLQKHKtk4RFA0/iwIDAQAB
AoGAEEvtjVa39f+4ns3y7PfBxOMG9Hu1z7z0f+KCiywNxN7W7XnPYgprSftBQiCN
6bkvUmivuxe2aI/tRRKjn5KidHofmR27N5Ehqy051KoceoX222eypjm+REfWKF4z
C4iWuho6yNyUr001+BquILVTppoCqEpNUhsOIQ3iEXI6n9ECQQDfvtCSG7D354af
SwvKHssoeLCZZ1zUXYIN+VRdqZ7SxuMPDDPrbpV8/oaIf8o4qYsp6vn+RTK1CC8S
4/mIOzO5AkEAwf+Ju2MezmQfRO7kA7WWSiQ7H+ozmiE3tux1ZFaGmH3Y4iXVsYE6
jrvl31JUmpg30vE+e+6M2HvlCUG4wie3YwJAHVf2HKNbEzgfFLcP5cMiG72wU34V
05gr3U9MvDfpp3UEBE4yZPcMzebHmBFCYSoDIU062hkTLGDVY6HYcQBqqQJBAKdo
Roin7muYRsuLgsQYrQpp3YOa7NCynYuC3YKqkBeDceclUxSMOeX2C1X9i/ikJngn
qg30OyKNN4YXasJQ15UCQQCEiagWymvwqqwAqvSLEIFGjfjkq1aLiVonhGMEWAkh
7hxTepxd6uJjjnkP6rNNzShC9suQHWW0+qlIkoM/8x3W
-----END RSA PRIVATE KEY-----
`

func TestRsa(t *testing.T) {
	src := "Hello World"
	fmt.Println(src)

	s1, err := RsaEncrypt(src, publicKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s1)

	s2, err := RsaDecrypt(s1, privateKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s2)
}

func TestRsaPath(t *testing.T) {
	src := "Hello World"
	fmt.Println(src)

	s1, err := RsaEncryptPath(src, `rsa_public_key.pub`)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s1)

	s2, err := RsaDecryptPath(s1, `rsa_private_key.pem`)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s2)
}

func TestRsaPKCS1(t *testing.T) {
	src := "Hello World"
	fmt.Println(src)

	s1, err := RsaEncrypt(src, publicKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s1)

	s2, err := RsaDecryptPKCS1(s1, privateKeyPKCS1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s2)
}

func TestRsaPKCS1Path(t *testing.T) {
	src := "Hello World"
	fmt.Println(src)

	s1, err := RsaEncryptPath(src, `rsa_public_key.pub`)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s1)

	s2, err := RsaDecryptPKCS1Path(s1, `rsa_private_key_pkcs1.pem`)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s2)
}
