package scan

import (
	"fmt"
	"testing"
)

func TestConvertAssign(t *testing.T) {
	soure := "abc"
	var target string
	err := ConvertAssign(&target, soure)
	fmt.Println(err)
	fmt.Println(target)
}
