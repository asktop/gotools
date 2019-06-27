package amsg

import (
	"encoding/json"
	"fmt"
	"testing"
)

func msg1() Msg {
	msg := New("abcdef")
	msg.Data("AAA")
	return msg
}

func TestMsg1(t *testing.T) {
	msg := msg1()
	fmt.Println(msg)
	fmt.Println("Status", msg.Status())
	fmt.Println("Error", msg.Error())
	fmt.Println("String", msg.String())
	fmt.Println(msg.Stringf())
	data, _ := json.Marshal(&msg)
	fmt.Println("json", string(data))
}

func msg2() Msg {
	msg := New("abc%sdef", "123")
	msg.Data("AAA")
	return msg
}

func TestMsg2(t *testing.T) {
	msg := msg2()
	fmt.Println(msg)
	fmt.Println("Status", msg.Status())
	fmt.Println("Error", msg.Error())
	fmt.Println("String", msg.String())
	fmt.Println(msg.Stringf())
	data, _ := json.Marshal(&msg)
	fmt.Println("json", string(data))
}

func msg3() Msg {
	msg := NewErr(fmt.Errorf("abc%sdef", "123"))
	msg.Data("AAA")
	return msg
}

func TestMsg3(t *testing.T) {
	msg := msg3()
	fmt.Println(msg)
	fmt.Println("Status", msg.Status())
	fmt.Println("Error", msg.Error())
	fmt.Println("String", msg.String())
	fmt.Println(msg.Stringf())
	data, _ := json.Marshal(&msg)
	fmt.Println("json", string(data))
}

func msg4() Msg {
	msg := NewErr(nil).Msg("abc%sdef", "123")
	msg.Data("AAA")
	return msg
}

func TestMsg4(t *testing.T) {
	msg := msg4()
	fmt.Println(msg)
	fmt.Println("Status", msg.Status())
	fmt.Println("Error", msg.Error())
	fmt.Println("String", msg.String())
	fmt.Println(msg.Stringf())
	data, _ := json.Marshal(&msg)
	fmt.Println("json", string(data))
}
