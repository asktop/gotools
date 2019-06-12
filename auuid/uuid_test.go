package auuid

import "testing"

func TestNew(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Log(New().String())
	}
}

func TestNewV1(t *testing.T) {
	for i := 0; i < 5; i++ {
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
