package amsg

import (
	"fmt"
	"testing"
)

func TestMsg(t *testing.T) {
	msg1 := testMsg1()
	fmt.Println(msg1)
	fmt.Println(msg1.Error())
	fmt.Println(msg1.Format())

	fmt.Println("---------")

	msg2 := testMsg2()
	fmt.Println(msg2)
	fmt.Println(msg2.Error())
	fmt.Println(msg2.Format())

	fmt.Println("---------")

	msg3 := testMsg3()
	fmt.Println(msg3)
	fmt.Println(msg3.Error())
	fmt.Println(msg3.Format())
}

func TestMsgStatus(t *testing.T) {
	msg4 := testMsg4()
	fmt.Println(msg4)
	fmt.Println(msg4.Error())
	fmt.Println(msg4.Format())
	fmt.Println(msg4.Status())
}

func testMsg1() Msg {
	msg := New("abcdef")
	return msg
}

func testMsg2() Msg {
	msg := New("abc%sdef", "123")
	return msg
}

func testMsg3() Msg {
	msg := NewErr(fmt.Errorf("abc%sdef", "123"))
	return msg
}

func testMsg4() Msg {
	msg := New("abcdef").SetSuccess()
	return msg
}
