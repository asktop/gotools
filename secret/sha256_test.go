package secret

import (
	"fmt"
	"testing"
)

func TestSha256(t *testing.T) {
	fmt.Println(len(Sha256("abc")))
	fmt.Println(Sha256("abc"))
}
