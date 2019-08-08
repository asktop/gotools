package asort

// int 切片
type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// int64 切片
type Int64Slice []int64

func (p Int64Slice) Len() int           { return len(p) }
func (p Int64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// float64 切片
type Float64Slice []float64

func (p Float64Slice) Len() int           { return len(p) }
func (p Float64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Float64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// string 切片
type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// map[string]interface 切片
type StringMapInterface struct {
	Key   string
	Value interface{}
}
type StringMapInterfaceSlice []StringMapInterface

func (s StringMapInterfaceSlice) Len() int           { return len(s) }
func (s StringMapInterfaceSlice) Less(i, j int) bool { return s[i].Key < s[j].Key }
func (s StringMapInterfaceSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// map[string]string 切片
type StringMapString struct {
	Key   string
	Value string
}
type StringMapStringSlice []StringMapString

func (s StringMapStringSlice) Len() int           { return len(s) }
func (s StringMapStringSlice) Less(i, j int) bool { return s[i].Key < s[j].Key }
func (s StringMapStringSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
