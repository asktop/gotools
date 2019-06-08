package secret

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	fmt.Println(len(Md5("abc")))
	fmt.Println(Md5("abc"))
}
