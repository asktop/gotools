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

func TestRandStr(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(RandStr(10))
	}
}

func TestRandMd5(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(RandMd5())
	}
}
