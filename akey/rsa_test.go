package akey

import (
	"fmt"
	"testing"
)

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

func TestRsaPKCS1(t *testing.T) {
	src := "Hello World"
	fmt.Println(src)

	s1, err := RsaEncrypt(src, publicKey2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s1)

	s2, err := RsaDecryptPKCS1(s1, privateKey2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s2)
}

var publicKey = `
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC5dbx/ZY0RbrDgMJFdwNOQ5iLa
S7m7JjGCWYwB5RHPSp9emKNA4GvCXhNrGoKxA7FsruQPhYqvSCJ2hU9DyXclXJ0d
Eyj1o9ZBaKf6hF+OL8I//ldjTsWGieAm9aOcf/8oydmwWv/DNJ1cflMVdaHO2WDn
j9OwTv+rDBIRaCdh/wIDAQAB
-----END PUBLIC KEY-----
`

var privateKey = `
-----BEGIN PRIVATE KEY-----
MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBALl1vH9ljRFusOAw
kV3A05DmItpLubsmMYJZjAHlEc9Kn16Yo0Dga8JeE2sagrEDsWyu5A+Fiq9IInaF
T0PJdyVcnR0TKPWj1kFop/qEX44vwj/+V2NOxYaJ4Cb1o5x//yjJ2bBa/8M0nVx+
UxV1oc7ZYOeP07BO/6sMEhFoJ2H/AgMBAAECgYB7BmD+WY0UrUrjzRQBDzLJAgDI
skcIoLNi9qfrcds4mRXTGInjNXwGOYXEHJfpeLuvjux2Z22yDLXfzVrharl/jDNf
K/bJGVhfiwV5x+yS+u1TXF86aSB/thcqleFXvRGIV3WuX7um+q7J8oA/OUPGxOyB
FjC8HltCWFRPm1az8QJBANsei0poDiq0vY0b3WrLetxO9Gp00D/fzBgIfbsw874s
UUWc3LHF/kLGSOvkvue93cnggCMeuG+i8Y2QalVz7/sCQQDYrN6BJhBNxUJnXBTr
F4bZ3Hms3GYLN7E+RuFpNB/wgzVN0LOkdKyogVLCt1inSa8szrLOXTcA3zMVfFdt
HMLNAkEAj6JmDFBJeRUha+5oJilcUC4xaddI65X4Y4itYpekL3U9kTRSNvZixcLU
6kz4F1EOodbYKC1rGULmtLWF/p4RIQJBAJ/KJLEjrAReg9ELxFV3XTiPcp/7Tbna
EXk29ocKLL/HU2kWj1Spwqbl8G2uns+H9IrbyFuNvMGE2PxwXV0XR8UCQAkKuKJL
TR0SqEpT1VKqpIauusMP5Jqfysvf/iR+nyS3zSId07+JZE9+uLgJOnoHVHRZNLB1
NWlpbP7I2GirvPo=
-----END PRIVATE KEY-----`

var publicKey2 = `
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDfw1/P15GQzGGYvNwVmXIGGxea
8Pb2wJcF7ZW7tmFdLSjOItn9kvUsbQgS5yxx+f2sAv1ocxbPTsFdRc6yUTJdeQol
DOkEzNP0B8XKm+Lxy4giwwR5LJQTANkqe4w/d9u129bRhTu/SUzSUIr65zZ/s6TU
GQD6QzKY1Y8xS+FoQQIDAQAB
-----END PUBLIC KEY-----
`

var privateKey2 = `
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDfw1/P15GQzGGYvNwVmXIGGxea8Pb2wJcF7ZW7tmFdLSjOItn9
kvUsbQgS5yxx+f2sAv1ocxbPTsFdRc6yUTJdeQolDOkEzNP0B8XKm+Lxy4giwwR5
LJQTANkqe4w/d9u129bRhTu/SUzSUIr65zZ/s6TUGQD6QzKY1Y8xS+FoQQIDAQAB
AoGAbSNg7wHomORm0dWDzvEpwTqjl8nh2tZyksyf1I+PC6BEH8613k04UfPYFUg1
0F2rUaOfr7s6q+BwxaqPtz+NPUotMjeVrEmmYM4rrYkrnd0lRiAxmkQUBlLrCBiF
u+bluDkHXF7+TUfJm4AZAvbtR2wO5DUAOZ244FfJueYyZHECQQD+V5/WrgKkBlYy
XhioQBXff7TLCrmMlUziJcQ295kIn8n1GaKzunJkhreoMbiRe0hpIIgPYb9E57tT
/mP/MoYtAkEA4Ti6XiOXgxzV5gcB+fhJyb8PJCVkgP2wg0OQp2DKPp+5xsmRuUXv
720oExv92jv6X65x631VGjDmfJNb99wq5QJBAMSHUKrBqqizfMdOjh7z5fLc6wY5
M0a91rqoFAWlLErNrXAGbwIRf3LN5fvA76z6ZelViczY6sKDjOxKFVqL38ECQG0S
pxdOT2M9BM45GJjxyPJ+qBuOTGU391Mq1pRpCKlZe4QtPHioyTGAAMd4Z/FX2MKb
3in48c0UX5t3VjPsmY0CQQCc1jmEoB83JmTHYByvDpc8kzsD8+GmiPVrausrjj4p
y2DQpGmUic2zqCxl6qXMpBGtFEhrUbKhOiVOJbRNGvWW
-----END RSA PRIVATE KEY-----
`
