package aclient

import (
    "fmt"
    "testing"
)

//测试WebSocket客户端连接
//import "github.com/gorilla/websocket"
func TestWebSocket(t *testing.T) {
    ws, err := NewWebSocketClientCase("ws", "127.0.0.1:8883", "/ws", true)
    if err != nil {
        fmt.Println(err)
        return
    }
    ws.Start()

    ws.SendMsg(`{"sub":"active.activings"}`)
    //ws.SendMsg(`{"sub":"active.detail.4"}`)
    //ws.Sleep(time.Second)
    //ws.SendMsg(`{"req":"active.active","info":{"pi_id":1,"username":"187423","active_id":4}}`)

    ws.Wait()
}
