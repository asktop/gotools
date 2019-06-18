package aunique

import (
	"testing"
	"time"
)

//唯一编码生成
func TestUniqueNo(t *testing.T) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 300)
		t.Log(UniqueNo(16))
	}

	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 300)
		t.Log(UniqueNo(18, "USER_"))
	}
}

func TestUniqueNo2(t *testing.T) {
	for i := 0; i < 1010; i++ {
		if i > 995 {
			t.Log(UniqueNo(16))
		} else {
			UniqueNo(16)
		}
	}
}
