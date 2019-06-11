package astring

import (
	"errors"
	"fmt"
	"testing"
)

func TestJoin(t *testing.T) {
	a := map[string]interface{}{}
	a["a"] = "abc"
	a["b"] = 123
	e := errors.New("err")
	fmt.Println(Join("uid:", 111, "data:", a, e))
}

func TestTrimSpaceToOne(t *testing.T) {
	fmt.Println("---" + TrimSpaceToOne("\ta	b	  c  d	") + "---")
	fmt.Println("---" + TrimSpaceToOne("a\t b	  c  d ") + "---")
}

func TestHidePwd(t *testing.T) {
	fmt.Println(HidePwd(""))
	fmt.Println(HidePwd("as"))
	fmt.Println(HidePwd("asdfjkhksdfkj"))
	fmt.Println(HidePwd("123456789"))
}

func TestHidePhone(t *testing.T) {
	fmt.Println(HidePhone(""))
	fmt.Println(HidePhone("13412345678"))
	fmt.Println(HidePhone("+8613412345678"))
	fmt.Println(HidePhone("8941123"))
	fmt.Println(HidePhone("89414567"))
	fmt.Println(HidePhone("0539-89414567"))
}

func TestHideEmail(t *testing.T) {
	fmt.Println(HideEmail(""))
	fmt.Println(HideEmail("as@163.com"))
	fmt.Println(HideEmail("asdfjkhksdfkj@163.com"))
	fmt.Println(HideEmail("123456789@163.com"))
	fmt.Println(HideEmail("13412345678@163.com"))
}
