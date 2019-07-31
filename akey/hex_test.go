package akey

import (
	"testing"
)

func TestHex(t *testing.T) {
	hello := "你好，世界！ hello world"
	//5L2g5aW977yM5LiW55WM77yBIGhlbGxvIHdvcmxk

	//hello := `{"exp":1546565110,"iat":1546478710,"uid":123,"username":"abc"}`
	//eyJleHAiOjE1NDY1NjUxMTAsImlhdCI6MTU0NjQ3ODcxMCwidWlkIjoxMjMsInVzZXJuYW1lIjoiYWJjIn0=

	//hello := `{"exp":1546565110,"iat":1546478710,"username":"abc"}`
	//eyJleHAiOjE1NDY1NjU4MTQsImlhdCI6MTU0NjQ3OTQxNCwidXNlcm5hbWUiOiJhYmMifQ==

	t.Log(hello)

	desrc := HexEncodeToString([]byte(hello))
	t.Log(desrc)

	ensrc := HexDecodeString(desrc)
	t.Log(string(ensrc))
}
