package base

import "github.com/asktop/gotools/aqywx/context"

// Base 基础信息管理
type Base struct {
	*context.Context
}

// NewBase 实例化
func NewBase(context *context.Context) *Base {
	base := new(Base)
	base.Context = context
	return base
}
