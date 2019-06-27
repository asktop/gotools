package amsg

import (
	"encoding/json"
	"errors"
	"fmt"
)

//自定义信息接口，可携带状态和数据
type Msg interface {
	Msg(msgf string, msgargs ...interface{}) Msg
	Status() (out bool)
	Error() error
	String() string
	Stringf() (msgf string, msgargs []interface{})
	Data(in ...interface{}) (out interface{})
}

//创建成功信息
func New(msgf string, msgargs ...interface{}) Msg {
	e := new(errMsg)
	e.msgf = msgf
	e.msgargs = msgargs
	e.status = true
	return e
}

//创建错误信息
func NewErr(err error) Msg {
	e := new(errMsg)
	if err != nil {
		e.err = err.Error()
		e.msgf = err.Error()
	}
	return e
}

type errMsg struct {
	status  bool
	err     string
	msgf    string
	msgargs []interface{}
	data    interface{}
}

//设置信息内容
func (e *errMsg) Msg(msgf string, msgargs ...interface{}) Msg {
	e.msgf = msgf
	e.msgargs = msgargs
	if !e.status && e.err == "" {
		//信息状态为错误，且error信息为空时，将error信息设置为自定义信息
		e.err = e.String()
	}
	return e
}

//获取信息状态
func (e *errMsg) Status() (out bool) {
	return e.status
}

//获取错误
func (e *errMsg) Error() error {
	if e.Status() {
		return nil
	}
	return errors.New(e.err)
}

//获取格式化后的信息内容
func (e *errMsg) String() string {
	if len(e.msgargs) == 0 {
		return e.msgf
	}
	return fmt.Sprintf(e.msgf, e.msgargs...)
}

//获取信息内容的格式化信息
func (e *errMsg) Stringf() (msgf string, msgargs []interface{}) {
	if len(e.msgargs) == 0 {
		return e.msgf, nil
	}
	return e.msgf, e.msgargs
}

//设置|获取信息数据
func (e *errMsg) Data(in ...interface{}) (out interface{}) {
	if len(in) > 0 {
		e.data = in[0]
	}
	return e.data
}

func (e *errMsg) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}
