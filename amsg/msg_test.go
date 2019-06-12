package amsg

import (
	"fmt"
	"testing"
)

func TestMsg(t *testing.T) {
	msg1 := testMsg1()
	t.Log(msg1)
	t.Log(msg1.Error())
	t.Log(msg1.Format())

	t.Log("---------")

	msg2 := testMsg2()
	t.Log(msg2)
	t.Log(msg2.Error())
	t.Log(msg2.Format())

	t.Log("---------")

	msg3 := testMsg3()
	t.Log(msg3)
	t.Log(msg3.Error())
	t.Log(msg3.Format())
}

func TestMsgStatus(t *testing.T) {
	msg4 := testMsg4()
	t.Log(msg4)
	t.Log(msg4.Error())
	t.Log(msg4.Format())
	t.Log(msg4.Status())
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
