package cast

import (
	"time"
)

// ToBool casts an interface to a bool type.
func ToBool(i interface{}, defaultVal ...bool) bool {
	v, err := ToBoolE(i)
	if len(defaultVal) > 0 && (i == nil || err != nil) {
		return defaultVal[0]
	}
	return v
}

// ToUint casts an interface to a uint type.
func ToUint(i interface{}, defaultVal ...uint) uint {
	v, err := ToUintE(i)
	if len(defaultVal) > 0 && (i == nil || err != nil) {
		return defaultVal[0]
	}
	return v
}

// ToUint8 casts an interface to a uint8 type.
func ToUint8(i interface{}, defaultVal ...uint8) uint8 {
	v, err := ToUint8E(i)
	if len(defaultVal) > 0 && (i == nil || err != nil) {
		return defaultVal[0]
	}
	return v
}

// ToUint16 casts an interface to a uint16 type.
func ToUint16(i interface{}, defaultVal ...uint16) uint16 {
	v, err := ToUint16E(i)
	if len(defaultVal) > 0 && (i == nil || err != nil) {
		return defaultVal[0]
	}
	return v
}

// ToUint32 casts an interface to a uint32 type.
func ToUint32(i interface{}, defaultVal ...uint32) uint32 {
	v, err := ToUint32E(i)
	if len(defaultVal) > 0 && (i == nil || err != nil) {
		return defaultVal[0]
	}
	return v
}

// ToUint64 casts an interface to a uint64 type.
func ToUint64(i interface{}, defaultVal ...uint64) uint64 {
	v, err := ToUint64E(i)
	if len(defaultVal) > 0 && (i == nil || err != nil) {
		return defaultVal[0]
	}
	return v
}

// ToInt casts an interface to an int type.
func ToInt(i interface{}, defaultVal ...int) int {
	v, err := ToIntE(i)
	if len(defaultVal) > 0 && (i == nil || err != nil) {
		return defaultVal[0]
	}
	return v
}

// ToInt8 casts an interface to an int8 type.
func ToInt8(i interface{}, defaultVal ...int8) int8 {
	v, err := ToInt8E(i)
	if len(defaultVal) > 0 && (i == nil || err != nil) {
		return defaultVal[0]
	}
	return v
}

// ToInt16 casts an interface to an int16 type.
func ToInt16(i interface{}, defaultVal ...int16) int16 {
	v, err := ToInt16E(i)
	if len(defaultVal) > 0 && (i == nil || err != nil) {
		return defaultVal[0]
	}
	return v
}

// ToInt32 casts an interface to an int32 type.
func ToInt32(i interface{}, defaultVal ...int32) int32 {
	v, err := ToInt32E(i)
	if len(defaultVal) > 0 && (i == nil || err != nil) {
		return defaultVal[0]
	}
	return v
}

// ToInt64 casts an interface to an int64 type.
func ToInt64(i interface{}, defaultVal ...int64) int64 {
	v, err := ToInt64E(i)
	if len(defaultVal) > 0 && (i == nil || err != nil) {
		return defaultVal[0]
	}
	return v
}

// ToFloat32 casts an interface to a float32 type.
func ToFloat32(i interface{}, defaultVal ...float32) float32 {
	v, err := ToFloat32E(i)
	if len(defaultVal) > 0 && (i == nil || err != nil) {
		return defaultVal[0]
	}
	return v
}

// ToFloat64 casts an interface to a float64 type.
func ToFloat64(i interface{}, defaultVal ...float64) float64 {
	v, err := ToFloat64E(i)
	if len(defaultVal) > 0 && (i == nil || err != nil) {
		return defaultVal[0]
	}
	return v
}

// ToString casts an interface to a string type.
func ToString(i interface{}, defaultVal ...string) string {
	v, err := ToStringE(i)
	if len(defaultVal) > 0 && (i == nil || err != nil) {
		return defaultVal[0]
	}
	return v
}

// ToTime casts an interface to a time.Time type.
func ToTime(i interface{}) time.Time {
	v, _ := ToTimeE(i)
	return v
}

// ToDuration casts an interface to a time.Duration type.
func ToDuration(i interface{}) time.Duration {
	v, _ := ToDurationE(i)
	return v
}

// ToSlice casts an interface to a []interface{} type.
func ToSlice(i interface{}, defaultVal ...interface{}) []interface{} {
	v, err := ToSliceE(i)
	if len(defaultVal) > 0 && (i == nil || err != nil) {
		return defaultVal
	}
	return v
}

// ToBoolSlice casts an interface to a []bool type.
func ToBoolSlice(i interface{}, defaultVal ...bool) []bool {
	v, err := ToBoolSliceE(i)
	if len(defaultVal) > 0 && (i == nil || err != nil) {
		return defaultVal
	}
	return v
}

// ToIntSlice casts an interface to a []int type.
func ToIntSlice(i interface{}, defaultVal ...int) []int {
	v, err := ToIntSliceE(i)
	if len(defaultVal) > 0 && (i == nil || err != nil) {
		return defaultVal
	}
	return v
}

// ToStringSlice casts an interface to a []string type.
func ToStringSlice(i interface{}, defaultVal ...string) []string {
	v, err := ToStringSliceE(i)
	if len(defaultVal) > 0 && (i == nil || err != nil) {
		return defaultVal
	}
	return v
}

// ToDurationSlice casts an interface to a []time.Duration type.
func ToDurationSlice(i interface{}) []time.Duration {
	v, _ := ToDurationSliceE(i)
	return v
}

// ToStringMap casts an interface to a map[string]interface{} type.
func ToStringMap(i interface{}) map[string]interface{} {
	v, _ := ToStringMapE(i)
	return v
}

// ToStringMapBool casts an interface to a map[string]bool type.
func ToStringMapBool(i interface{}) map[string]bool {
	v, _ := ToStringMapBoolE(i)
	return v
}

// ToStringMapInt casts an interface to a map[string]int type.
func ToStringMapInt(i interface{}) map[string]int {
	v, _ := ToStringMapIntE(i)
	return v
}

// ToStringMapInt64 casts an interface to a map[string]int64 type.
func ToStringMapInt64(i interface{}) map[string]int64 {
	v, _ := ToStringMapInt64E(i)
	return v
}

// ToStringMapString casts an interface to a map[string]string type.
func ToStringMapString(i interface{}) map[string]string {
	v, _ := ToStringMapStringE(i)
	return v
}

// ToStringMapStringSlice casts an interface to a map[string][]string type.
func ToStringMapStringSlice(i interface{}) map[string][]string {
	v, _ := ToStringMapStringSliceE(i)
	return v
}
