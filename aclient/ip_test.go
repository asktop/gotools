package aclient

import (
    "fmt"
    "testing"
)

func TestGetIp(t *testing.T) {
    fmt.Println(GetLocalIp())
}

