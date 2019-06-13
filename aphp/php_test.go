package aphp

import (
	"testing"
)

//php序列化和反序列化
func TestPhpSerialize(t *testing.T) {
	str := `a:1:{s:3:"php";s:24:"世界上最好的语言";}`

	// unserialize() in php
	out, _ := UnMarshal([]byte(str))

	t.Log(out) //map[php:世界上最好的语言]

	// serialize() in php
	jsonbyte, _ := Marshal(out)

	t.Log(string(jsonbyte)) // a:1:{s:3:"php";s:24:"世界上最好的语言";}
}
