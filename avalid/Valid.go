package avalid

import "github.com/asktop/gotools/acast"

type Valid struct {
	Name     string
	title    string
	value    interface{}
	valueStr string
	checks   []checkIface
}

//创建验证
func New(name string, value interface{}, title ...string) *Valid {
	valid := &Valid{}
	valid.Name = name
	if len(title) > 0 {
		valid.title = title[0]
	} else {
		valid.title = name
	}
	valid.value = value
	valid.valueStr = acast.ToString(value)
	return valid
}

//数值的范围
func (v *Valid) Between(min interface{}, max interface{}, msg ...string) *Valid {
	v.checks = append(v.checks, &between{
		title:    v.title,
		value:    v.value,
		valueStr: v.valueStr,
		msgs:     msg,
		min:      min,
		max:      max,
	})
	return v
}

//数值相等
func (v *Valid) Equal(equalVal interface{}, msg ...string) *Valid {
	v.checks = append(v.checks, &equal{
		title:    v.title,
		value:    v.value,
		valueStr: v.valueStr,
		msgs:     msg,
		equalVal: equalVal,
	})
	return v
}

//在切片中
func (v *Valid) InSlice(slice interface{}, msg ...string) *Valid {
	v.checks = append(v.checks, &inSlice{
		title:    v.title,
		value:    v.value,
		valueStr: v.valueStr,
		msgs:     msg,
		slice:    slice,
	})
	return v
}

//必须为数值
func (v *Valid) IsDecimal(length interface{}, msg ...string) *Valid {
	v.checks = append(v.checks, &isDecimal{
		title:    v.title,
		value:    v.value,
		valueStr: v.valueStr,
		msgs:     msg,
		length:   length,
	})
	return v
}

//必须为整数
func (v *Valid) IsInt(msg ...string) *Valid {
	v.checks = append(v.checks, &isInt{
		title:    v.title,
		value:    v.value,
		valueStr: v.valueStr,
		msgs:     msg,
	})
	return v
}

//必须为数字字符串
func (v *Valid) IsNumber(length interface{}, msg ...string) *Valid {
	v.checks = append(v.checks, &isNumber{
		title:    v.title,
		value:    v.value,
		valueStr: v.valueStr,
		msgs:     msg,
		length:   length,
	})
	return v
}

//字符串长度范围
func (v *Valid) Length(min interface{}, max interface{}, msg ...string) *Valid {
	v.checks = append(v.checks, &length{
		title:    v.title,
		value:    v.value,
		valueStr: v.valueStr,
		msgs:     msg,
		min:      min,
		max:      max,
	})
	return v
}

//正则表达式验证
func (v *Valid) Regex(exp string, msg ...string) *Valid {
	v.checks = append(v.checks, &regex{
		title:    v.title,
		value:    v.value,
		valueStr: v.valueStr,
		msgs:     msg,
		exp:      exp,
	})
	return v
}

//必需
func (v *Valid) Required(msg ...string) *Valid {
	v.checks = append(v.checks, &required{
		title:    v.title,
		value:    v.value,
		valueStr: v.valueStr,
		msgs:     msg,
	})
	return v
}

//相同
func (v *Valid) Same(sameVal interface{}, msg ...string) *Valid {
	v.checks = append(v.checks, &same{
		title:    v.title,
		value:    v.value,
		valueStr: v.valueStr,
		msgs:     msg,
		sameVal:  sameVal,
	})
	return v
}

//执行验证
func (v *Valid) Check() (msg string, ok bool) {
	for _, vc := range v.checks {
		msg, ok = vc.Check()
		if !ok {
			return
		}
	}
	return "", true
}

//执行验证
func (v *Valid) Checks() (msgs map[string]string, ok bool) {
	msgs = map[string]string{}
	for _, vc := range v.checks {
		msg, ok := vc.Check()
		if !ok {
			msgs[v.Name] = msg
			return msgs, false
		}
	}
	return msgs, true
}
