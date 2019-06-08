package secret

import (
	"fmt"
	"testing"
)

func TestHmacMd5(t *testing.T) {
	fmt.Println(len(HmacMd5("abc", "abc")))
	fmt.Println(HmacMd5("abc", "abc"))
}

func TestHmacSha1(t *testing.T) {
	fmt.Println(len(HmacSha1("abc", "abc")))
	fmt.Println(HmacSha1("abc", "abc"))
}

func TestHmacSha256(t *testing.T) {
	fmt.Println(len(HmacSha256("abc", "abc")))
	fmt.Println(HmacSha256("abc", "abc"))
}
