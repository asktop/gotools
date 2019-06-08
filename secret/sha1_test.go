package secret

import (
	"fmt"
	"testing"
)

func TestSha1(t *testing.T) {
	fmt.Println(len(Sha1("abc")))
	fmt.Println(Sha1("abc"))
}
