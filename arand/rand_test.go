package arand

import (
	"math/rand"
	"testing"
)

func TestRand(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(rand.Intn(10))
	}
	for i := 0; i < 10; i++ {
		t.Log(Rand(1, 3))
	}
	for i := 0; i < 10; i++ {
		t.Log(Rand(-1, 1))
	}
}

func TestRandStr(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(RandStr(1, "0"))
	}
	for i := 0; i < 10; i++ {
		t.Log(RandStr(6))
	}
}

func TestRandMd5(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(RandMd5())
	}
}
