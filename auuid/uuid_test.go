package auuid

import (
	"testing"
	"time"
)

//uuid生成
func TestNew(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Log(New().String())
		t.Log(New().Bytes())
		t.Log(New().Version())
	}
}

func TestNewV1(t *testing.T) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Nanosecond)
		u, _ := NewV1()
		t.Log(u.String())
	}
}

func TestNewV4(t *testing.T) {
	for i := 0; i < 5; i++ {
		u, _ := NewV4()
		t.Log(u.String())
	}
}
