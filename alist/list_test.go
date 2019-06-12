package alist

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	l1 := New()
	l1.PushBack("a")
	l1.PushBack(2)
	l1.PushBack("c")
	fmt.Println(l1)

	l2 := New()
	s2 := []interface{}{"s1", 123, "s3"}
	rs, _ := json.Marshal(s2)
	err := json.Unmarshal(rs, l2)
	fmt.Println(err)
	fmt.Println(l2)
}
