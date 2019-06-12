package akey

import (
	"testing"
)

func TestSha1(t *testing.T) {
	t.Log(len(Sha1("abc")))
	t.Log(Sha1("abc"))
}
