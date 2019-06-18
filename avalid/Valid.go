package avalid

import "github.com/asktop/gotools/acast"

type Valid struct {
	Name     string
	title    string
	value    interface{}
	valueStr string
	isCheck  bool
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
	valid.isCheck = true
	return valid
}

//执行是否进行校验方法
func (v *Valid) IsCheck(f func() bool) *Valid {
	v.isCheck = f()
	return v
}

//执行自定义方法
func (v *Valid) Func(f func() (msg string, ok bool)) *Valid {
	v.checks = append(v.checks, &funcExec{
		f: f,
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

//在切片中
func (v *Valid) InSlice(slice []string, msg ...string) *Valid {
	v.checks = append(v.checks, &inSlice{
		title:    v.title,
		value:    v.value,
		valueStr: v.valueStr,
		msgs:     msg,
		slice:    slice,
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

//数值比较
// rs：比较状态 0：等于；1：大于；-1：小于；10：大于等于；-10：小于等于
func (v *Valid) Cmp(number interface{}, rs int, msg ...string) *Valid {
	v.checks = append(v.checks, &cmp{
		title:    v.title,
		value:    v.value,
		valueStr: v.valueStr,
		msgs:     msg,
		number:   number,
		rs:       rs,
	})
	return v
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

//必须为数值
func (v *Valid) IsDecimal(length []int, msg ...string) *Valid {
	v.checks = append(v.checks, &isDecimal{
		title:    v.title,
		value:    v.value,
		valueStr: v.valueStr,
		msgs:     msg,
		length:   length,
	})
	return v
}

//必须为手机号
func (v *Valid) IsPhone(msg ...string) *Valid {
	v.checks = append(v.checks, &isPhone{
		title:    v.title,
		value:    v.value,
		valueStr: v.valueStr,
		msgs:     msg,
	})
	return v
}

//必须为电话号码
func (v *Valid) IsTel(msg ...string) *Valid {
	v.checks = append(v.checks, &isTel{
		title:    v.title,
		value:    v.value,
		valueStr: v.valueStr,
		msgs:     msg,
	})
	return v
}

//必须为Email
func (v *Valid) IsEmail(msg ...string) *Valid {
	v.checks = append(v.checks, &isEmail{
		title:    v.title,
		value:    v.value,
		valueStr: v.valueStr,
		msgs:     msg,
	})
	return v
}

//必须为身份证号码
func (v *Valid) IsIDCard(msg ...string) *Valid {
	v.checks = append(v.checks, &isIDCard{
		title:    v.title,
		value:    v.value,
		valueStr: v.valueStr,
		msgs:     msg,
	})
	return v
}

//检查账号（字母开头，数字字母下划线）
func (v *Valid) IsAccount(length []int, msg ...string) *Valid {
	v.checks = append(v.checks, &isAccount{
		title:    v.title,
		value:    v.value,
		valueStr: v.valueStr,
		msgs:     msg,
		length:   length,
	})
	return v
}

//检查密码
// level: 密码强度级别
// 	1：包含数字、字母
// 	2：包含数字、字母、下划线
// 	3：包含数字、字母、特殊字符
// 	4：包含数字、大小写字母
// 	5：包含数字、大小写字母、下划线
// 	6：包含数字、大小写字母、特殊字符
func (v *Valid) IsPwd(level int, length []int, msg ...string) *Valid {
	v.checks = append(v.checks, &isPwd{
		title:    v.title,
		value:    v.value,
		valueStr: v.valueStr,
		msgs:     msg,
		level:    level,
		length:   length,
	})
	return v
}

//执行验证
func (v *Valid) Check() (msg string, ok bool) {
	if v.isCheck {
		for _, vc := range v.checks {
			msg, ok = vc.Check()
			if !ok {
				return
			}
		}
	}
	return "", true
}

//执行验证
func (v *Valid) Checks() (msgs map[string]string, ok bool) {
	msgs = map[string]string{}
	if v.isCheck {
		for _, vc := range v.checks {
			msg, ok := vc.Check()
			if !ok {
				msgs[v.Name] = msg
				return msgs, false
			}
		}
	}
	return msgs, true
}
