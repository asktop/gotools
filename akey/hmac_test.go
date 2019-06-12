package akey

import (
	"testing"
)

func TestHmacMd5(t *testing.T) {
	t.Log(len(HmacMd5("abc", "abc")))
	t.Log(HmacMd5("abc", "abc"))
}

func TestHmacSha1(t *testing.T) {
	t.Log(len(HmacSha1("abc", "abc")))
	t.Log(HmacSha1("abc", "abc"))
}

func TestHmacSha256(t *testing.T) {
	t.Log(len(HmacSha256("abc", "abc")))
	t.Log(HmacSha256("abc", "abc"))
}
