package avalid

import (
	"testing"
)

func TestNew(t *testing.T) {
	msg, ok := New("username", "abcd", "用户名").Required().Check()
	t.Log(ok, msg)
	msg, ok = New("username", "abcd", "用户名").Required().Length(6, 20).Check()
	t.Log(ok, msg)
	msg, ok = New("username", "abcd", "用户名").Required().Same("abc").Check()
	t.Log(ok, msg)
	msg, ok = New("username", "abcd", "用户名").Required().InSlice([]string{"a", "b", "c", "abc"}).Check()
	t.Log(ok, msg)
	msg, ok = New("amount", "12.3456", "金额").Required().IsInt().Check()
	t.Log(ok, msg)
	msg, ok = New("amount", "12.3456", "金额").Required().IsDecimal(nil).Check()
	t.Log(ok, msg)
	msg, ok = New("amount", "12.3456", "金额").Required().IsNumber(nil).Check()
	t.Log(ok, msg)
	msg, ok = New("amount", "12.3456", "金额").Required().Between(12, "12.15").Check()
	t.Log(ok, msg)
}

func TestNews(t *testing.T) {
	valid := News().Valid("username", "abcd", "用户名").Required().Length(6, 20).
		Valid("amount", "12.3456").Required().Between(12, "12.15")

	msg, ok := valid.Check()
	t.Log(ok, msg)

	msgs, ok := valid.Checks()
	if !ok {
		for k, v := range msgs {
			t.Log(k, v)
		}
	}
}
