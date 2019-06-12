package ascan

import (
	"testing"
)

func TestConvertAssign(t *testing.T) {
	soure := "abc"
	var target string
	err := ConvertAssign(&target, soure)
	t.Log(err)
	t.Log(target)
}
