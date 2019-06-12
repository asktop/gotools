package acast

import "time"

// ToIfaceSlice casts an interface to a []interface{} type.
func ToIfaceSlice(i interface{}, defaultVal ...interface{}) []interface{} {
	v, err := ToIfaceSliceE(i)
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
