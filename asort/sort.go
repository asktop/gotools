package asort

import "sort"

func Ints(a []int) {
	sort.Sort(IntSlice(a))
}

func RInts(a []int) {
	sort.Slice(IntSlice(a), func(i, j int) bool {
		return a[i] > a[j]
	})
}

func Int64s(a []int) {
	sort.Sort(Int64Slice(a))
}

func RInt64s(a []int) {
	sort.Slice(Int64Slice(a), func(i, j int) bool {
		return a[i] > a[j]
	})
}

func Float64s(a []float64) {
	sort.Sort(Float64Slice(a))
}

func RFloat64s(a []float64) {
	sort.Slice(Float64Slice(a), func(i, j int) bool {
		return a[i] > a[j]
	})
}

func Strings(a []string) {
	sort.Sort(StringSlice(a))
}

func RStrings(a []string) {
	sort.Slice(StringSlice(a), func(i, j int) bool {
		return a[i] > a[j]
	})
}

func StringMapInterfaces(data map[string]interface{}) StringMapInterfaceSlice {
	maps := StringMapInterfaceSlice{}
	for k, v := range data {
		maps = append(maps, StringMapInterface{Key: k, Value: v})
	}
	sort.Sort(maps)
	return maps
}

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

func StringMapStrings(data map[string]string) StringMapStringSlice {
	maps := StringMapStringSlice{}
	for k, v := range data {
		maps = append(maps, StringMapString{Key: k, Value: v})
	}
	sort.Sort(maps)
	return maps
}

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
