package akey

import (
	"testing"
)

func TestMd5(t *testing.T) {
	t.Log(len(Md5("abc")))
	t.Log(Md5("abc"))
}
