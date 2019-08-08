package asort

import (
	"fmt"
	"github.com/asktop/gotools/acast"
	"sort"
	"strings"
)

func Ints(a []int) {
	sort.Sort(IntSlice(a))
}

func RInts(a []int) {
	sort.Sort(sort.Reverse(IntSlice(a)))
	// 等价于
	//sort.Slice(IntSlice(a), func(i, j int) bool {
	//	return a[i] > a[j]
	//})
}

func Int64s(a []int64) {
	sort.Sort(Int64Slice(a))
}

func RInt64s(a []int64) {
	sort.Sort(sort.Reverse(Int64Slice(a)))
	// 等价于
	//sort.Slice(Int64Slice(a), func(i, j int) bool {
	//	return a[i] > a[j]
	//})
}

func Float64s(a []float64) {
	sort.Sort(Float64Slice(a))
}

func RFloat64s(a []float64) {
	sort.Sort(sort.Reverse(Float64Slice(a)))
	// 等价于
	//sort.Slice(Float64Slice(a), func(i, j int) bool {
	//	return a[i] > a[j]
	//})
}

func Strings(a []string) {
	sort.Sort(StringSlice(a))
}

func RStrings(a []string) {
	sort.Sort(sort.Reverse(StringSlice(a)))
	// 等价于
	//sort.Slice(StringSlice(a), func(i, j int) bool {
	//	return a[i] > a[j]
	//})
}

//对 map[string]interface{} 排序
func StringMapInterfaces(data map[string]interface{}) StringMapInterfaceSlice {
	maps := StringMapInterfaceSlice{}
	for k, v := range data {
		maps = append(maps, StringMapInterface{Key: k, Value: v})
	}
	sort.Sort(maps)
	return maps
}

//对 map[string]interface{} 倒序排序
func RStringMapInterfaces(data map[string]interface{}) StringMapInterfaceSlice {
	s := StringMapInterfaceSlice{}
	for k, v := range data {
		s = append(s, StringMapInterface{Key: k, Value: v})
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i].Key > s[j].Key
	})
	return s
}

//对 map[string]string 排序
func StringMapStrings(data map[string]string) StringMapStringSlice {
	maps := StringMapStringSlice{}
	for k, v := range data {
		maps = append(maps, StringMapString{Key: k, Value: v})
	}
	sort.Sort(maps)
	return maps
}

//对 map[string]string 倒序排序
func RStringMapStrings(data map[string]string) StringMapStringSlice {
	s := StringMapStringSlice{}
	for k, v := range data {
		s = append(s, StringMapString{Key: k, Value: v})
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i].Key > s[j].Key
	})
	return s
}

//对 param map 参数排序，并按照 {key}={value} 格式拼接，最后通过 {sep} 拼接
func SortParamString(params map[string]string, sep string) string {
	if len(params) == 0 {
		return ""
	}
	newParams := StringMapStrings(params)
	var paramStrs = make([]string, 0)
	for _, param := range newParams {
		paramStrs = append(paramStrs, fmt.Sprintf("%s=%s", param.Key, param.Value))
	}
	return strings.Join(paramStrs, sep)
}

//对 param map 参数排序，并按照 {key}={value} 格式拼接，最后通过 {sep} 拼接
func SortParamInterface(params map[string]interface{}, sep string) string {
	if len(params) == 0 {
		return ""
	}
	newParams := StringMapInterfaces(params)
	var paramStrs = make([]string, 0)
	for _, param := range newParams {
		paramStrs = append(paramStrs, fmt.Sprintf("%s=%s", param.Key, acast.ToString(param.Value)))
	}
	return strings.Join(paramStrs, sep)
}
