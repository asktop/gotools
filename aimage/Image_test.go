package aimage

import (
	"fmt"
	"strings"
	"testing"
)

func TestThumb(t *testing.T) {
	inPath := `D:\00Down\cache\问道.jpg`
	outPath := strings.Replace(inPath, ".", "_thumb.", 1)
	fmt.Println("inPath=", inPath, " outPath=", outPath)

	err := ThumbPath(inPath, outPath, 150, 150, 100)
	if err != nil {
		fmt.Println(err)
	}
}

func TestClip(t *testing.T) {
	inPath := `D:\00Down\cache\问道.jpg`
	outPath := strings.Replace(inPath, ".", "_clip.", 1)
	fmt.Println("inPath=", inPath, " outPath=", outPath)

	err := ClipPath(inPath, outPath, 50, 50, 150, 150, 100)
	if err != nil {
		fmt.Println(err)
	}
}
