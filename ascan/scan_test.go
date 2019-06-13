package ascan

import (
	"testing"
)

func TestConvertAssign(t *testing.T) {
	soure := "abc"
	var target string
	err := ConvertAssign(&target, soure)
	if err != nil {
		t.Error(err)
	}
	t.Log(target)
}
