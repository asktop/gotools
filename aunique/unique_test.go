package aunique

import (
	"testing"
	"time"
)

func TestUniqueNo(t *testing.T) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 300)
		t.Log(UniqueNo(15))
	}

	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 300)
		t.Log(UniqueNo(15, "USER"))
	}
}
