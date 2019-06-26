package ajson

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
)

type Server struct {
	//若没有json标签，则json字段名和struct属性名不区分大小写匹配；若有json标签，则json字段名和struct属性标签完全匹配。
	ID         int         `json:"-"`                       // struct中的ID不会解析到json串中；同样，json串中若有ID也不会解析到struct中
	ServerName string      `json:"serverName"`              // struct和json通过tag匹配，仅解析能匹配到的字段，若匹配不到则取零值
	ServerIP   string      `json:"serverIP, omitempty"`     // tag中的omitempty表示，若ServerIP没有值，则不会解析到json串中；同样，若json串serverIP没有值，则不会解析到struct中
	Code       int         `json:"code, string, omitempty"` // tag中的string表示，json串中的数据类型；json串中必须为string类型才能解析到struct中，struct中int等类型解析到json中时会自动转为string类型
	Timestamp  json.Number `json:"timestamp, omitempty"`    // json.Number 是string类型，但在json串中显示为数值类型；json串中的数值类型会string类型都可以解析到struct中
}

type Serverslice struct {
	Servers []Server `json:"servers"`
}

//json序列化：将对象转成json字符串或字节码响应给客户端
func TestMarshal(t *testing.T) {
	var s Serverslice
	s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1", Code: 360})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "", Timestamp: "1542077370"})

	//json序列化
	//b, _ := json.Marshal(s)
	//json序列化并格式化
	b, _ := json.MarshalIndent(s, "", "  ")

	//将对象转成字节码响应给客户端(http.ResponseWriter)
	//w.Write(b)

	//将对象转成json字符串
	t.Log(string(b))
}

//json反序列化：将json字符串或json字节码转化为对象
func TestUnmarshal(t *testing.T) {
	var s Serverslice
	str := `{"servers":[{"ID":127,"serverName":"Shanghai_VPN","serverIP":"127.0.0.1","code":"180"},{"serverName":"Beijing_VPN","serverIP":"","code":"300","timestamp":1542077370}]}`

	//json反序列化
	err := json.Unmarshal([]byte(str), &s)
	if err != nil {
		t.Error(err)
	}

	t.Log(s)
}

//自定义json反序列化：解析json.Number为map[string]string
func TestDecodeToMapString(t *testing.T) {
	str := `{"ID":"127","serverName":"Shanghai_VPN","serverIP":"127.0.0.1","code":180}`
	params, _ := DecodeToMapString([]byte(str))
	t.Log(params)
}

func TestEncode(t *testing.T) {
	fmt.Println(Encode(nil))
	fmt.Println(Encode(12))
	fmt.Println(Encode("12"))
	fmt.Println(Encode(true))
	fmt.Println(Encode(errors.New("12")))
}
