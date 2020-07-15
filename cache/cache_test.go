package cache

import (
    "fmt"
    "testing"
    "time"
)

func TestCache(t *testing.T) {
    err := NewCache().Put("test1", "testv1", time.Second*300)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(NewCache().Get("test1"))
    fmt.Println(NewCache().GetString("test1"))
}
