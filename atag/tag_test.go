package atag

import (
	"fmt"
	"testing"
)

type demo struct {
	Id       int64
	Address  string //地址
	Currency string //币种
	Image    string //图片
}

func TestNew(t *testing.T) {
	tag := New(`
type demo struct {
	Id       int64
	Address  string //地址
	Currency string //币种
	Image    string //图片
}
	`, "json", "form")
	fmt.Println(tag)
}

func TestNew2(t *testing.T) {
	tag := New(`
	Id       int64
	Address  string //地址
	Currency string //币种
	Image    string //图片
	`, "json", "form")
	fmt.Println(tag)
}
