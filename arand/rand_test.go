package arand

import (
	"testing"
)

func TestIntn(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(Intn(10))
	}
}

func TestRandInt64(t *testing.T) {
	//for i := 0; i < 10; i++ {
	//	t.Log(rand.Intn(10))
	//}

	for i := 0; i < 10; i++ {
		t.Log(RandInt64(1, 3))
	}

	//for i := 0; i < 10; i++ {
	//	t.Log(RandInt64(-1, 1))
	//}
}

func TestRandStr(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(RandStr(1, "0"))
	}

	for i := 0; i < 10; i++ {
		t.Log(RandStr(6))
	}
}

func TestRandUUID(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(RandUUID())
	}
}

func TestRandMd5(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(RandMd5())
	}
}

func TestRandBase32(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(RandBase32())
	}
}
