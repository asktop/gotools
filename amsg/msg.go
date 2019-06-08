package amsg

import "fmt"

//自定义错误信息接口
type Msg interface {
	error
	Format() (format string, args []interface{})
	SetSuccess() Msg
	Status() bool
	SetData(data interface{}) Msg
	Data() (data interface{})
}

//创建带format格式化的error信息
func New(format string, args ...interface{}) Msg {
	e := new(msgString)
	e.format = format
	e.args = args
	return e
}

//快速创建error信息
func NewErr(err error) Msg {
	e := new(msgString)
	e.format = err.Error()
	return e
}

type msgString struct {
	format string
	args   []interface{}
	status bool
	data   interface{}
}

//格式化返回错误信息内容
func (e *msgString) Error() string {
	if e.args == nil {
		return e.format
	}
	return fmt.Sprintf(e.format, e.args...)
}

//全部返回错误信息内容
func (e *msgString) Format() (format string, args []interface{}) {
	if e.args == nil {
		return e.format, nil
	}
	return e.format, e.args
}

//设置信息状态为成功（默认失败）
func (e *msgString) SetSuccess() Msg {
	e.status = true
	return e
}

//获取信息状态
func (e *msgString) Status() bool {
	return e.status
}

//设置信息数据
func (e *msgString) SetData(data interface{}) Msg {
	e.data = data
	return e
}

//获取信息数据
func (e *msgString) Data() (data interface{}) {
	return e.data
}
