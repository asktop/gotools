package ajson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

/*
1.1、string 转换成 []byte
	b := []byte(str)
1.2、[]byte 转换成 string （以下两种方式转换结构并不相同）
	str := string(b)
	str := fmt.Sprintf("%x", b)

2.1、io.Reader 转换成 []byte
	reply, err := ioutil.ReadAll(resp.Body)
2.2、[]byte 转换成 io.Reader
	bytes.NewReader(data)
*/

//将 map或结构 转换成 json类型string
func Encode(v interface{}) string {
	switch s := v.(type) {
	case nil:
		return ""
	case string:
		return s
	case error:
		return s.Error()
	}
	objVal := reflect.ValueOf(v)
	if objVal.Kind() == reflect.Ptr {
		if objVal.IsNil() {
			return ""
		}
	}
	data, _ := json.Marshal(v)
	return string(data)
}

//将 json类型string 转换成 map
func Decode(v string) (map[string]interface{}, error) {
	dataMap := map[string]interface{}{}
	err := json.Unmarshal([]byte(v), &dataMap)
	return dataMap, err
}

//json反序列化：[]byte -> map[string]string (解析数值为string)
func DecodeToMapString(data []byte) (dataMap map[string]string, err error) {
	dataMap = map[string]string{}
	tempMap := map[string]interface{}{}
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	err = decoder.Decode(&tempMap)
	if err != nil {
		return
	}
	for k, v := range tempMap {
		switch t := v.(type) {
		case string:
			dataMap[k] = v.(string)
		case json.Number:
			dataMap[k] = v.(json.Number).String()
		default:
			_ = t
			return nil, fmt.Errorf("unknown type, key:%v, type:%v", k, reflect.TypeOf(t).String())
		}
	}
	return
}

//将 结构 转换成 map
func StructToMap(v interface{}) (map[string]interface{}, error) {
	dataMap := map[string]interface{}{}
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		return nil, err
	}
	return dataMap, nil
}

//将 map 转换成 结构
func MapToStruct(source map[string]interface{}, result interface{}) error {
	data, err := json.Marshal(source)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &result)
}
