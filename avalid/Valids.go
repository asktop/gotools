package avalid

import (
	"github.com/asktop/gotools/acast"
)

type Valids struct {
	valids []*Valid
}

func News() *Valids {
	return &Valids{}
}

func (vs *Valids) Valid(name string, value interface{}, title ...string) *Valids {
	valid := &Valid{}
	valid.Name = name
	if len(title) > 0 {
		valid.title = title[0]
	} else {
		valid.title = name
	}
	valid.value = value
	valid.valueStr = acast.ToString(value)
	valid.isCheck = true
	vs.valids = append(vs.valids, valid)
	return vs
}

//执行是否进行校验方法
func (vs *Valids) IsCheck(f func() bool) *Valids {
	l := len(vs.valids)
	if l != 0 {
		v := vs.valids[l-1]
		v.isCheck = f()
	}
	return vs
}

//执行自定义方法
func (vs *Valids) Func(f func() (msg string, ok bool)) *Valids {
	l := len(vs.valids)
	if l != 0 {
		v := vs.valids[l-1]
		v.checks = append(v.checks, &funcExec{
			f: f,
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

//不能有空格
func (vs *Valids) HasNoBlank(msg ...string) *Valids {
	l := len(vs.valids)
	if l != 0 {
		v := vs.valids[l-1]
		v.checks = append(v.checks, &hasNoBlank{
			title:    v.title,
			value:    v.value,
			valueStr: v.valueStr,
			msgs:     msg,
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

//在切片中
func (vs *Valids) InSlice(slice []string, msg ...string) *Valids {
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

//数值比较
// rs：比较状态 0：等于；1：大于；-1：小于；10：大于等于；-10：小于等于
func (vs *Valids) Cmp(number interface{}, rs int, msg ...string) *Valids {
	l := len(vs.valids)
	if l != 0 {
		v := vs.valids[l-1]
		v.checks = append(v.checks, &cmp{
			title:    v.title,
			value:    v.value,
			valueStr: v.valueStr,
			msgs:     msg,
			number:   number,
			rs:       rs,
		})
	}
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

//必须为整数
func (vs *Valids) IsInt(msg ...string) *Valids {
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

//必须为数值
func (vs *Valids) IsDecimal(length []uint, msg ...string) *Valids {
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

//必须为手机号
func (vs *Valids) IsMobile(msg ...string) *Valids {
	l := len(vs.valids)
	if l != 0 {
		v := vs.valids[l-1]
		v.checks = append(v.checks, &isMobile{
			title:    v.title,
			value:    v.value,
			valueStr: v.valueStr,
			msgs:     msg,
		})
	}
	return vs
}

//必须为电话号码
func (vs *Valids) IsTel(msg ...string) *Valids {
	l := len(vs.valids)
	if l != 0 {
		v := vs.valids[l-1]
		v.checks = append(v.checks, &isTel{
			title:    v.title,
			value:    v.value,
			valueStr: v.valueStr,
			msgs:     msg,
		})
	}
	return vs
}

//必须为Email
func (vs *Valids) IsEmail(msg ...string) *Valids {
	l := len(vs.valids)
	if l != 0 {
		v := vs.valids[l-1]
		v.checks = append(v.checks, &isEmail{
			title:    v.title,
			value:    v.value,
			valueStr: v.valueStr,
			msgs:     msg,
		})
	}
	return vs
}

//必须为身份证号码
func (vs *Valids) IsIDCard(msg ...string) *Valids {
	l := len(vs.valids)
	if l != 0 {
		v := vs.valids[l-1]
		v.checks = append(v.checks, &isIDCard{
			title:    v.title,
			value:    v.value,
			valueStr: v.valueStr,
			msgs:     msg,
		})
	}
	return vs
}

//必须为银行卡号
func (vs *Valids) IsBankCard(msg ...string) *Valids {
	l := len(vs.valids)
	if l != 0 {
		v := vs.valids[l-1]
		v.checks = append(v.checks, &isBankCard{
			title:    v.title,
			value:    v.value,
			valueStr: v.valueStr,
			msgs:     msg,
		})
	}
	return vs
}

//检查账号（字母开头，数字字母下划线）
func (vs *Valids) IsAccount(length []uint, msg ...string) *Valids {
	l := len(vs.valids)
	if l != 0 {
		v := vs.valids[l-1]
		v.checks = append(v.checks, &isAccount{
			title:    v.title,
			value:    v.value,
			valueStr: v.valueStr,
			msgs:     msg,
			length:   length,
		})
	}
	return vs
}

//检查密码
// level: 密码强度级别
// 	1：包含数字、字母
// 	2：包含数字、字母、下划线
// 	3：包含数字、字母、特殊字符
// 	4：包含数字、大小写字母
// 	5：包含数字、大小写字母、下划线
// 	6：包含数字、大小写字母、特殊字符
func (vs *Valids) IsPwd(level uint, length []uint, msg ...string) *Valids {
	l := len(vs.valids)
	if l != 0 {
		v := vs.valids[l-1]
		v.checks = append(v.checks, &isPwd{
			title:    v.title,
			value:    v.value,
			valueStr: v.valueStr,
			msgs:     msg,
			level:    level,
			length:   length,
		})
	}
	return vs
}

//执行验证
func (vs *Valids) Check() (msg string, ok bool) {
	for _, v := range vs.valids {
		if v.isCheck {
			for _, vc := range v.checks {
				msg, ok = vc.Check()
				if !ok {
					return
				}
			}
		}
	}
	return "", true
}

//执行验证
func (vs *Valids) Checks() (msgs map[string]string, ok bool) {
	msgs = map[string]string{}
	for _, v := range vs.valids {
		if v.isCheck {
			for _, vc := range v.checks {
				msg, ok := vc.Check()
				if !ok {
					msgs[v.Name] = msg
					break
				}
			}
		}
	}
	if len(msgs) > 0 {
		return msgs, false
	} else {
		return msgs, true
	}
}
