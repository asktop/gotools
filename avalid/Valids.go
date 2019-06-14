package avalid

import "github.com/asktop/gotools/acast"

type Valids struct {
	valids []*Valid
}

func News() *Valids {
	return &Valids{}
}

func (vs *Valids) Valid(title string, value interface{}) *Valids {
	valid := &Valid{}
	valid.title = title
	valid.value = value
	valid.valueStr = acast.ToString(value)
	vs.valids = append(vs.valids, valid)
	return vs
}

//数值的范围
func (vs *Valids) Between(min interface{}, max interface{}, msg ...string) *Valids {
	l := len(vs.valids)
	if l != 0 {
		v := vs.valids[l-1]
		v.checks = append(v.checks, &between{
			title:    v.title,
			value:    v.value,
			valueStr: v.valueStr,
			msgs:     msg,
			min:      min,
			max:      max,
		})
	}
	return vs
}

//数值相等
func (vs *Valids) Equal(equalVal interface{}, msg ...string) *Valids {
	l := len(vs.valids)
	if l != 0 {
		v := vs.valids[l-1]
		v.checks = append(v.checks, &equal{
			title:    v.title,
			value:    v.value,
			valueStr: v.valueStr,
			msgs:     msg,
			equalVal: equalVal,
		})
	}
	return vs
}

//在切片中
func (vs *Valids) InSlice(slice interface{}, msg ...string) *Valids {
	l := len(vs.valids)
	if l != 0 {
		v := vs.valids[l-1]
		v.checks = append(v.checks, &inSlice{
			title:    v.title,
			value:    v.value,
			valueStr: v.valueStr,
			msgs:     msg,
			slice:    slice,
		})
	}
	return vs
}

//必须为数值
func (vs *Valids) IsDecimal(length interface{}, msg ...string) *Valids {
	l := len(vs.valids)
	if l != 0 {
		v := vs.valids[l-1]
		v.checks = append(v.checks, &isDecimal{
			title:    v.title,
			value:    v.value,
			valueStr: v.valueStr,
			msgs:     msg,
			length:   length,
		})
	}
	return vs
}

//必须为整数
func (vs *Valids) IsInt(length interface{}, msg ...string) *Valids {
	l := len(vs.valids)
	if l != 0 {
		v := vs.valids[l-1]
		v.checks = append(v.checks, &isInt{
			title:    v.title,
			value:    v.value,
			valueStr: v.valueStr,
			msgs:     msg,
		})
	}
	return vs
}

//必须为数字字符串
func (vs *Valids) IsNumber(length interface{}, msg ...string) *Valids {
	l := len(vs.valids)
	if l != 0 {
		v := vs.valids[l-1]
		v.checks = append(v.checks, &isNumber{
			title:    v.title,
			value:    v.value,
			valueStr: v.valueStr,
			msgs:     msg,
			length:   length,
		})
	}
	return vs
}

//字符串长度范围
func (vs *Valids) Length(min interface{}, max interface{}, msg ...string) *Valids {
	l := len(vs.valids)
	if l != 0 {
		v := vs.valids[l-1]
		v.checks = append(v.checks, &length{
			title:    v.title,
			value:    v.value,
			valueStr: v.valueStr,
			msgs:     msg,
			min:      min,
			max:      max,
		})
	}
	return vs
}

//正则表达式验证
func (vs *Valids) Regex(exp string, msg ...string) *Valids {
	l := len(vs.valids)
	if l != 0 {
		v := vs.valids[l-1]
		v.checks = append(v.checks, &regex{
			title:    v.title,
			value:    v.value,
			valueStr: v.valueStr,
			msgs:     msg,
			exp:      exp,
		})
	}
	return vs
}

//必需
func (vs *Valids) Required(msg ...string) *Valids {
	l := len(vs.valids)
	if l != 0 {
		v := vs.valids[l-1]
		v.checks = append(v.checks, &required{
			title:    v.title,
			value:    v.value,
			valueStr: v.valueStr,
			msgs:     msg,
		})
	}
	return vs
}

//相同
func (vs *Valids) Same(sameVal interface{}, msg ...string) *Valids {
	l := len(vs.valids)
	if l != 0 {
		v := vs.valids[l-1]
		v.checks = append(v.checks, &same{
			title:    v.title,
			value:    v.value,
			valueStr: v.valueStr,
			msgs:     msg,
			sameVal:  sameVal,
		})
	}
	return vs
}

//执行验证
func (vs *Valids) Check() (msg string, ok bool) {
	for _, v := range vs.valids {
		for _, vc := range v.checks {
			msg, ok = vc.Check()
			if !ok {
				return
			}
		}
	}
	return "", true
}
