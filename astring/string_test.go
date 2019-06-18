package astring

import (
	"errors"
	"testing"
)

func TestSubstr(t *testing.T) {
	str := "0123456789"
	t.Log(Substr(str, 8))
	t.Log(Substr(str, 8, 3))
	t.Log(Substr(str, 0, 3))
	t.Log(Substr(str, 0, -3))
	t.Log(Substr(str, 1, -3))
}

func TestTrimSpaceToOne(t *testing.T) {
	t.Log("---" + TrimSpaceToOne("\ta	b	  c  d	") + "---")
	t.Log("---" + TrimSpaceToOne("a\t b	  c  d ") + "---")
}

func TestIntToStr(t *testing.T) {
	t.Log(IntToStr(123, 6))
	t.Log(IntToStr(123456789, 6))
	t.Log(IntToStr(123456789, 6, true))
}

func TestJoin(t *testing.T) {
	a := map[string]interface{}{}
	a["a"] = "abc"
	a["b"] = 123
	e := errors.New("err")
	t.Log(Join("uid:", 111, "data:", a, e))
}

func TestHideNo(t *testing.T) {
	t.Log(HideNo("", 1, 1))
	t.Log(HideNo("", 1, 1, 6))
	t.Log(HideNo("a", 2, 2, 6))
	t.Log(HideNo("as", 2, 2, 6))
	t.Log(HideNo("asw", 1, 1, 6))
	t.Log(HideNo("asdfjkhksdfkj", 3, 0, 6))
	t.Log(HideNo("123456789", 2, 2, 6))
}

func TestHidePhone(t *testing.T) {
	t.Log(HidePhone(""))
	t.Log(HidePhone("13412345678"))
	t.Log(HidePhone("+8613412345678"))
	t.Log(HidePhone("8941123"))
	t.Log(HidePhone("89414567"))
	t.Log(HidePhone("0539-89414567"))
}

func TestHideEmail(t *testing.T) {
	t.Log(HideEmail(""))
	t.Log(HideEmail("as@163.com"))
	t.Log(HideEmail("asdfjkhksdfkj@163.com"))
	t.Log(HideEmail("123456789@163.com"))
	t.Log(HideEmail("13412345678@163.com"))
}

func TestHidePwd(t *testing.T) {
	t.Log(HidePwd(""))
	t.Log(HidePwd("", true))
	t.Log(HidePwd("asw"))
	t.Log(HidePwd("123456789"))
}

func TestToFirstUpper(t *testing.T) {
	a := "abcdef"
	t.Log(ToFirstUpper(a))
	b := "ABCDEF"
	t.Log(ToFirstLower(b))
}

func TestToCamelCase(t *testing.T) {
	a := "user_id"
	b := "User_Id"
	c := "userId"
	d := "UserId"
	t.Log(ToCamelCase(a))
	t.Log(ToCamelCase(b))
	t.Log(ToCamelCase(c))
	t.Log(ToCamelCase(d))

	t.Log("----------")

	t.Log(TocamelCase(a))
	t.Log(TocamelCase(b))
	t.Log(TocamelCase(c))
	t.Log(TocamelCase(d))

	t.Log("----------")

	t.Log(ToUnderscoreCase(a))
	t.Log(ToUnderscoreCase(b))
	t.Log(ToUnderscoreCase(c))
	t.Log(ToUnderscoreCase(d))

	t.Log("----------")

	t.Log(TounderscoreCase(a))
	t.Log(TounderscoreCase(b))
	t.Log(TounderscoreCase(c))
	t.Log(TounderscoreCase(d))
}
