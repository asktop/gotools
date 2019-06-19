package acast

import "github.com/asktop/gotools/amapstruct"

func MapToStruct(input interface{}, output interface{}) error {
	return amapstruct.Decode(input, output)
}

func StructToMap(input interface{}, output interface{}) error {
	return amapstruct.Decode(input, output)
}
