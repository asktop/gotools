package aclient

import (
    "encoding/json"
    "fmt"
    "github.com/asktop/gotools/agzip"
    "github.com/gorilla/websocket"
    "net/http"
    "net/url"
    "time"
)

//创建WebSocket客户端连接
//import "github.com/gorilla/websocket"
func NewWebSocketClient(scheme, host, path string) (*websocket.Conn, error) {
    //处理scheme
    hscheme := "http"
    if scheme == "wss" {
        hscheme = "https"
    } else {
        scheme = "ws"
    }
    //处理host
    if host == "" {
        return nil, fmt.Errorf("host不能为空")
    }
    //组装url
    Url := url.URL{Scheme: scheme, Host: host, Path: path}
    Origin := url.URL{Scheme: hscheme, Host: host, Path: ""}
    //创建webSocket连接
    header := http.Header{}
    header.Set("Origin", Origin.String())
    ws, _, err := websocket.DefaultDialer.Dial(Url.String(), header)
    if err != nil {
        return nil, fmt.Errorf("创建webSocket连接出错，url:%s，err: %v", Url.String(), err)
    }
    return ws, nil
}

type WebSocketClientCase struct {
    ws        *websocket.Conn
    closed    bool
    wsSendMsg chan []byte
    s2cGzip   bool
}

//创建WebSocket客户端连接实例
func NewWebSocketClientCase(scheme, host, path string, s2cGzip bool) (*WebSocketClientCase, error) {
    ws, err := NewWebSocketClient(scheme, host, path)
    if err != nil {
        return nil, err
    }
    return &WebSocketClientCase{ws: ws, wsSendMsg: make(chan []byte, 10000), s2cGzip: s2cGzip}, nil
}

//启动
func (c *WebSocketClientCase) Start() {
    //读取ws信息
    go func(ws *websocket.Conn) {
        defer func() {
            c.closed = true
        }()
        for {
            if c.closed {
                break
            }
            _, message, err := ws.ReadMessage()
            if err != nil {
                fmt.Println("读取ws失败", err)
                break
            }

            if c.s2cGzip {
                message = agzip.UnGzip(message)
            }

            //处理消息
            var infos map[string]interface{}
            err = json.Unmarshal([]byte(message), &infos)
            if err != nil {
                fmt.Println("解析读取的message出错", "message:", string(message), "err:", err)
                continue
            }

            if _, ok := infos["ping"]; ok {
                //处理ping(读取到ping信息，发送pong信息)
                obj := pingJson{}
                if err := json.Unmarshal(message, &obj); err == nil {
                    c.SendObj(pongJson{Pong: obj.Ping})
                } else {
                    fmt.Println("解析ping时出错", "message:", string(message), "err:", err)
                }
            } else if _, ok := infos["pong"]; ok {
                //处理pong(读取到ping信息，发送pong信息)
                //fmt.Println("读取pong信息", "message:", string(message))
            } else {
                //读取信息
                fmt.Println("读取信息", string(message))
            }
        }
    }(c.ws)

    //发送ws信息
    go func(ws *websocket.Conn) {
        defer func() {
            c.closed = true
        }()
        //ticker := time.NewTicker(time.Second * 5)
        for {
            select {
            //定时发送ping信息
            //case ti := <-ticker.C:
            //    c.SendObj(pingJson{Ping: ti.Unix()})

            //发送信息
            case message := <-c.wsSendMsg:
                if c.closed {
                    return
                }
                if err := ws.WriteMessage(websocket.BinaryMessage, message); err != nil {
                    fmt.Println("发送信息失败", string(message), err)
                    return
                } else {
                    fmt.Println("发送信息", string(message))
                }
            }
        }
    }(c.ws)

    go func() {
        for {
            if c.closed {
                c.ws.Close()
                fmt.Println("--- 关闭ws连接 ---")
                return
            }
            time.Sleep(time.Second)
        }
    }()
}

//发送
func (c *WebSocketClientCase) SendMsg(msg string) {
    c.wsSendMsg <- []byte(msg)
}

//发送
func (c *WebSocketClientCase) SendObj(obj interface{}) {
    message, _ := json.Marshal(obj)
    c.wsSendMsg <- message
}

//等待
func (c *WebSocketClientCase) Sleep(d time.Duration) {
    time.Sleep(d)
}

//等待
func (c *WebSocketClientCase) Wait() {
    select {}
}

type pingJson struct {
    Ping int64 `json:"ping"` //时间戳
}

type pongJson struct {
    Pong int64 `json:"pong"` //时间戳
}
