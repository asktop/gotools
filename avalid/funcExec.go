package avalid

//执行自定义方法
type funcExec struct {
	f        func() (msg string, ok bool)
}

func (c *funcExec) Check() (msg string, ok bool) {
	return c.f()
}
