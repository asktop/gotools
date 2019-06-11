package unique

import (
	"fmt"
	"testing"
	"time"
)

func TestUniqueNo(t *testing.T) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 300)
		fmt.Println(UniqueNo(15))
	}

	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 300)
		fmt.Println(UniqueNo(15, "USER"))
	}
}
