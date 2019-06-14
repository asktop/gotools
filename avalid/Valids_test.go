package avalid

import (
	"testing"
)

func TestNew(t *testing.T) {
	msg, ok := New("用户名", "abcd").Required().Check()
	t.Log(ok, msg)
	msg, ok = New("用户名", "abcd", ).Required().Length(6, 20).Check()
	t.Log(ok, msg)
	msg, ok = New("用户名", "abcd").Required().Same("abc").Check()
	t.Log(ok, msg)
	msg, ok = New("用户名", "abcd").Required().InSlice([]string{"a", "b", "c", "abc"}).Check()
	t.Log(ok, msg)
	msg, ok = New("金额", "12.3456").Required().IsInt().Check()
	t.Log(ok, msg)
	msg, ok = New("金额", "12.3456").Required().IsDecimal(nil).Check()
	t.Log(ok, msg)
	msg, ok = New("金额", "12.3456").Required().IsNumber(nil).Check()
	t.Log(ok, msg)
	msg, ok = New("金额", "12.3456").Required().Between(12, "12.15").Check()
	t.Log(ok, msg)
}

func TestNews(t *testing.T) {
	msg, ok := News().
		Valid("用户名", "abcd").Required().Length(4, 20).
		Valid("金额", "12.3456").Required().Between(12, "12.15").
		Check()
	t.Log(ok, msg)
}
