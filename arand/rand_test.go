package arand

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestRand(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(10))
	}
}

func TestRandom(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(Random(1, 3))
	}
}

func TestRandMd5(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(RandMd5())
	}
}

func TestRandStr(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(RandStr(10))
	}
}
