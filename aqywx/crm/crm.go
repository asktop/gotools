package crm

import "github.com/asktop/gotools/aqywx/context"

// Crm 外部联系人管理
type Crm struct {
	*context.Context
}

// NewCrm 实例化
func NewCrm(context *context.Context) *Crm {
	crm := new(Crm)
	crm.Context = context
	return crm
}
