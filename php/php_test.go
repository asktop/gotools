package php

import (
	"fmt"
	"testing"
)

//php序列化和反序列化
//安装 go get -u github.com/techleeone/gophp/serialize
func TestPhpSerialize(t *testing.T) {
	str := `a:1:{s:3:"php";s:24:"世界上最好的语言";}`

	// unserialize() in php
	out, _ := UnMarshal([]byte(str))

	fmt.Println(out) //map[php:世界上最好的语言]

	// serialize() in php
	jsonbyte, _ := Marshal(out)

	fmt.Println(string(jsonbyte)) // a:1:{s:3:"php";s:24:"世界上最好的语言";}
}
