package akey

import (
	"testing"
)

func TestBase64(t *testing.T) {
	hello := "你好，世界！ hello world"
	//5L2g5aW977yM5LiW55WM77yBIGhlbGxvIHdvcmxk

	//hello := `{"exp":1546565110,"iat":1546478710,"uid":123,"username":"abc"}`
	//eyJleHAiOjE1NDY1NjUxMTAsImlhdCI6MTU0NjQ3ODcxMCwidWlkIjoxMjMsInVzZXJuYW1lIjoiYWJjIn0=

	//hello := `{"exp":1546565110,"iat":1546478710,"username":"abc"}`
	//eyJleHAiOjE1NDY1NjU4MTQsImlhdCI6MTU0NjQ3OTQxNCwidXNlcm5hbWUiOiJhYmMifQ==

	t.Log(hello)

	desrc := Base64EncodeToString([]byte(hello))
	t.Log(desrc)

	ensrc := Base64DecodeString(desrc)
	t.Log(string(ensrc))
}

func TestBase64_2(t *testing.T) {
	hello := "你好，世界！ hello world"
	t.Log(hello)

	desrc := Base64Encode([]byte(hello))
	t.Log(string(desrc))

	ensrc := Base64Decode(desrc)
	t.Log(string(ensrc))
}
