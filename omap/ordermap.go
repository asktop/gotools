package omap

import (
	"encoding/json"
	"github.com/asktop/gotools/cast"
	"reflect"
	"sort"
	"strings"
)

//有序map
type OrderMap struct {
	keys   []string
	values map[string]interface{}
}

//创建有序map
func New() *OrderMap {
	o := OrderMap{}
	o.keys = []string{}
	o.values = map[string]interface{}{}
	return &o
}

//创建有序map
func SetMap(mapData interface{}) *OrderMap {
	o := OrderMap{}
	o.keys = []string{}
	o.values = map[string]interface{}{}

	mapVal := reflect.ValueOf(mapData)
	if mapVal.Kind() == reflect.Map {
		keyVals := mapVal.MapKeys()
		for _, keyVal := range keyVals {
			key := cast.ToString(keyVal.Interface())
			value := mapVal.MapIndex(keyVal).Interface()
			o.keys = append(o.keys, key)
			o.values[key] = value
		}
		return &o
	} else {
		return nil
	}
}

//获取所有的key
func (o *OrderMap) Keys() []string {
	return o.keys
}

//获取所有的值（map类型）
func (o *OrderMap) Values() map[string]interface{} {
	return o.values
}

//赋值
func (o *OrderMap) Set(key string, value interface{}) *OrderMap {
	_, ok := o.values[key]
	if !ok {
		o.keys = append(o.keys, key)
	}
	o.values[key] = value
	return o
}

//取值
func (o *OrderMap) Get(key string) interface{} {
	return o.values[key]
}

//获取 bool 类型
func (o *OrderMap) GetBool(key string) bool {
	return cast.ToBool(o.values[key])
}

//获取 int64 类型
func (o *OrderMap) GetInt64(key string) int64 {
	return cast.ToInt64(o.values[key])
}

//获取 string 类型
func (o *OrderMap) GetString(key string) string {
	return cast.ToString(o.values[key])
}

//取值并判断是否存在key
func (o *OrderMap) GetOk(key string) (interface{}, bool) {
	value, ok := o.values[key]
	return value, ok
}

//获取 bool 类型并判断是否存在key
func (o *OrderMap) GetBoolOk(key string) (bool, bool) {
	value, ok := o.values[key]
	if ok {
		return cast.ToBool(value), ok
	}
	return false, ok
}

//获取 int64 类型并判断是否存在key
func (o *OrderMap) GetInt64Ok(key string) (int64, bool) {
	value, ok := o.values[key]
	if ok {
		return cast.ToInt64(value), ok
	}
	return 0, ok
}

//获取 string 类型并判断是否存在key
func (o *OrderMap) GetStringOk(key string) (string, bool) {
	value, ok := o.values[key]
	if ok {
		return cast.ToString(value), ok
	}
	return "", ok
}

//删除
func (o *OrderMap) Delete(keys ...string) *OrderMap {
	for _, key := range keys {
		// check key is in use
		_, ok := o.values[key]
		if !ok {
			continue
		}
		// remove from keys
		for i, k := range o.keys {
			if k == key {
				o.keys = append(o.keys[:i], o.keys[i+1:]...)
				break
			}
		}
		// remove from values
		delete(o.values, key)
	}
	return o
}

//按key排序
func (o *OrderMap) SortKey() *OrderMap {
	sort.Strings(o.keys)
	return o
}

//按key倒序排序
func (o *OrderMap) RSortKey() *OrderMap {
	o.Sort(func(a *OrderMapPair, b *OrderMapPair) bool {
		return a.key > b.key
	})
	return o
}

//按value排序
func (o *OrderMap) SortValue() *OrderMap {
	o.Sort(func(a *OrderMapPair, b *OrderMapPair) bool {
		return cast.ToString(a.value) < cast.ToString(b.value)
	})
	return o
}

//按value倒序排序
func (o *OrderMap) RSortValue() *OrderMap {
	o.Sort(func(a *OrderMapPair, b *OrderMapPair) bool {
		return cast.ToString(a.value) > cast.ToString(b.value)
	})
	return o
}

// Sort Sort the map using your sort func
func (o *OrderMap) Sort(lessFunc func(a *OrderMapPair, b *OrderMapPair) bool) {
	pairs := make([]*OrderMapPair, len(o.keys))
	for i, key := range o.keys {
		pairs[i] = &OrderMapPair{key, o.values[key]}
	}

	sort.Sort(ByOrderMapPair{pairs, lessFunc})

	for i, pair := range pairs {
		o.keys[i] = pair.key
	}
}

//实现json序列化接口，保证对象序列化后和map序列化后一致
func (o OrderMap) MarshalJSON() ([]byte, error) {
	s := "{"
	for _, k := range o.keys {
		// add key
		kEscaped := strings.Replace(k, `"`, `\"`, -1)
		s = s + `"` + kEscaped + `":`
		// add value
		v := o.values[k]
		vBytes, err := json.Marshal(v)
		if err != nil {
			return []byte{}, err
		}
		s = s + string(vBytes) + ","
	}
	if len(o.keys) > 0 {
		s = s[0 : len(s)-1]
	}
	s = s + "}"
	return []byte(s), nil
}

//实现json反序列化接口，保证map的json可以序列化到ordermap中
func (o *OrderMap) UnmarshalJSON(b []byte) error {
	if o.values == nil {
		o.values = map[string]interface{}{}
	}
	var err error
	err = mapStringToOrderedMap(string(b), o)
	if err != nil {
		return err
	}
	return nil
}

func mapStringToOrderedMap(s string, o *OrderMap) error {
	// parse string into map
	m := map[string]interface{}{}
	err := json.Unmarshal([]byte(s), &m)
	if err != nil {
		return err
	}
	// Get the order of the keys
	orderedKeys := []OrderMapIndex{}
	for k, _ := range m {
		kEscaped := strings.Replace(k, `"`, `\"`, -1)
		kQuoted := `"` + kEscaped + `"`
		// Find how much content exists before this key.
		// If all content from this key and after is replaced with a close
		// brace, it should still form a valid json string.
		sTrimmed := s
		for len(sTrimmed) > 0 {
			lastIndex := strings.LastIndex(sTrimmed, kQuoted)
			if lastIndex == -1 {
				break
			}
			sTrimmed = sTrimmed[0:lastIndex]
			sTrimmed = strings.TrimSpace(sTrimmed)
			if len(sTrimmed) > 0 && sTrimmed[len(sTrimmed)-1] == ',' {
				sTrimmed = sTrimmed[0 : len(sTrimmed)-1]
			}
			maybeValidJson := sTrimmed + "}"
			testMap := map[string]interface{}{}
			err := json.Unmarshal([]byte(maybeValidJson), &testMap)
			if err == nil {
				// record the position of this key in s
				ki := OrderMapIndex{
					Key:   k,
					Index: len(sTrimmed),
				}
				orderedKeys = append(orderedKeys, ki)
				// shorten the string to get the next key
				startOfValueIndex := lastIndex + len(kQuoted)
				valueStr := s[startOfValueIndex : len(s)-1]
				valueStr = strings.TrimSpace(valueStr)
				if len(valueStr) > 0 && valueStr[0] == ':' {
					valueStr = valueStr[1:len(valueStr)]
				}
				valueStr = strings.TrimSpace(valueStr)
				if valueStr[0] == '{' {
					// if the value for this key is a map
					// find end of valueStr by removing everything after last }
					// until it forms valid json
					hasValidJson := false
					i := 1
					for i < len(valueStr) && !hasValidJson {
						if valueStr[i] != '}' {
							i = i + 1
							continue
						}
						subTestMap := map[string]interface{}{}
						testValue := valueStr[0 : i+1]
						err = json.Unmarshal([]byte(testValue), &subTestMap)
						if err == nil {
							hasValidJson = true
							valueStr = testValue
							break
						}
						i = i + 1
					}
					// convert to orderedmap
					// this may be recursive it values in the map are also maps
					if hasValidJson {
						newMap := New()
						err := mapStringToOrderedMap(valueStr, newMap)
						if err != nil {
							return err
						}
						m[k] = *newMap
					}
				} else if valueStr[0] == '[' {
					// if the value for this key is a slice
					// find end of valueStr by removing everything after last ]
					// until it forms valid json
					hasValidJson := false
					i := 1
					for i < len(valueStr) && !hasValidJson {
						if valueStr[i] != ']' {
							i = i + 1
							continue
						}
						subTestSlice := []interface{}{}
						testValue := valueStr[0 : i+1]
						err = json.Unmarshal([]byte(testValue), &subTestSlice)
						if err == nil {
							hasValidJson = true
							valueStr = testValue
							break
						}
						i = i + 1
					}
					// convert to slice with any map items converted to
					// orderedmaps
					// this may be recursive if values in the slice are slices
					if hasValidJson {
						var newSlice []interface{}
						err := sliceStringToSliceWithOrderedMaps(valueStr, &newSlice)
						if err != nil {
							return err
						}
						m[k] = newSlice
					}
				} else {
					o.Set(k, m[k])
				}
				break
			}
		}
	}
	// Sort the keys
	sort.Sort(ByOrderMapIndex(orderedKeys))
	// Convert sorted keys to string slice
	k := []string{}
	for _, ki := range orderedKeys {
		k = append(k, ki.Key)
	}
	// Set the OrderMap values
	o.values = m
	o.keys = k
	return nil
}

func sliceStringToSliceWithOrderedMaps(valueStr string, newSlice *[]interface{}) error {
	// if the value for this key is a []interface, convert any map items to an orderedmap.
	// find end of valueStr by removing everything after last ]
	// until it forms valid json
	itemsStr := strings.TrimSpace(valueStr)
	itemsStr = itemsStr[1 : len(itemsStr)-1]
	// get next item in the slice
	itemIndex := 0
	startItem := 0
	endItem := 0
	for endItem <= len(itemsStr) {
		couldBeItemEnd := false
		couldBeItemEnd = couldBeItemEnd || endItem == len(itemsStr)
		couldBeItemEnd = couldBeItemEnd || (endItem < len(itemsStr) && itemsStr[endItem] == ',')
		if !couldBeItemEnd {
			endItem = endItem + 1
			continue
		}
		// if this substring compiles to json, it's the next item
		possibleItemStr := strings.TrimSpace(itemsStr[startItem:endItem])
		var possibleItem interface{}
		err := json.Unmarshal([]byte(possibleItemStr), &possibleItem)
		if err != nil {
			endItem = endItem + 1
			continue
		}
		if possibleItemStr[0] == '{' {
			// if item is map, convert to orderedmap
			oo := New()
			err := mapStringToOrderedMap(possibleItemStr, oo)
			if err != nil {
				return err
			}
			// add new orderedmap item to new slice
			slice := *newSlice
			slice = append(slice, *oo)
			*newSlice = slice
		} else if possibleItemStr[0] == '[' {
			// if item is slice, convert to slice with orderedmaps
			var newItem []interface{}
			err := sliceStringToSliceWithOrderedMaps(possibleItemStr, &newItem)
			if err != nil {
				return err
			}
			// replace original slice item with new slice
			slice := *newSlice
			slice = append(slice, newItem)
			*newSlice = slice
		} else {
			// any non-slice and non-map item, just add json parsed item
			slice := *newSlice
			slice = append(slice, possibleItem)
			*newSlice = slice
		}
		// remove this item from itemsStr
		startItem = endItem + 1
		endItem = endItem + 1
		itemIndex = itemIndex + 1
	}
	return nil
}

type OrderMapPair struct {
	key   string
	value interface{}
}

func (kv *OrderMapPair) Key() string {
	return kv.key
}

func (kv *OrderMapPair) Value() interface{} {
	return kv.value
}

type ByOrderMapPair struct {
	Pairs    []*OrderMapPair
	LessFunc func(a *OrderMapPair, j *OrderMapPair) bool
}

func (a ByOrderMapPair) Len() int           { return len(a.Pairs) }
func (a ByOrderMapPair) Swap(i, j int)      { a.Pairs[i], a.Pairs[j] = a.Pairs[j], a.Pairs[i] }
func (a ByOrderMapPair) Less(i, j int) bool { return a.LessFunc(a.Pairs[i], a.Pairs[j]) }

type OrderMapIndex struct {
	Key   string
	Index int
}

type ByOrderMapIndex []OrderMapIndex

func (a ByOrderMapIndex) Len() int           { return len(a) }
func (a ByOrderMapIndex) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByOrderMapIndex) Less(i, j int) bool { return a[i].Index < a[j].Index }
