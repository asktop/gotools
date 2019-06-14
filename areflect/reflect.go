package areflect

import (
	"reflect"
	"runtime"
	"strings"
)

func GetFuncAllName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func GetFuncName(i interface{}) string {
	allName := GetFuncAllName(i)
	names := strings.Split(allName, ".")
	return names[len(names)-1]
}
