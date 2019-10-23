package aunique

import (
	"fmt"
	"testing"
	"time"
)

//唯一编码生成
func TestUniqueNo(t *testing.T) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 300)
		t.Log(UniqueNo(20))
	}

	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 300)
		t.Log(UniqueNo(20, "USER_"))
	}
}

func TestUniqueNo2(t *testing.T) {
	for i := 0; i < 1010; i++ {
		fmt.Println(UniqueNo(20))
		//if i > 995 {
		//	fmt.Println(UniqueNo(20))
		//} else {
		//	UniqueNo(20)
		//}
	}
}

func BenchmarkUniqueNo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println(UniqueNo(20))
		//for i := 0; i < 100; i++ {
		//	if i > 995 {
		//		fmt.Println(UniqueNo(20))
		//	} else {
		//		UniqueNo(20)
		//	}
		//}
	}
}
