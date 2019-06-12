package akey

import (
	"testing"
)

func TestSha256(t *testing.T) {
	t.Log(len(Sha256("abc")))
	t.Log(Sha256("abc"))
}
