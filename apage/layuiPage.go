package apage

import (
	"github.com/asktop/dbr"
	"github.com/asktop/gotools/atime"
)

type LayuiPage struct {
	Page       uint64              `json:"page"`        //当前页码
	Limit      uint64              `json:"limit"`       //每页条数
	Count      int64               `json:"count"`       //总条数
	PageCount  int64               `json:"page_count"`  //总页数
	Data       []map[string]string `json:"data"`        //分页业务数据
	DataSource []map[string]string `json:"data_source"` //分页原始数据
	Code       int64               `json:"code"`
	Msg        string              `json:"msg"`
	pageData   *Page
	execFields []string               //自定义处理的所有字段
	execValues map[string]interface{} //自定义处理的所有字段的处理值
}

//layui
func NewLayuiPage(stmt *dbr.SelectStmt, page uint64, limit uint64, defLimit ...uint64) (p *LayuiPage, err error) {
	p = new(LayuiPage)
	p.pageData, err = NewPage(stmt, page, limit, defLimit...)
	p.Page = p.pageData.Page
	p.Limit = p.pageData.Limit
	p.Count = p.pageData.Total
	p.PageCount = p.pageData.TotalPage
	if err != nil {
		p.Code = 1
		p.Msg = err.Error()
		return p, err
	} else {
		p.Code = 0
		p.Data = p.pageData.Data
		return p, nil
	}
}

//默认值
func (p *LayuiPage) Default(name string, defaultVal string) *LayuiPage {
	p.pageData.Default(name, defaultVal)
	return p
}

//下拉选项处理
func (p *LayuiPage) Select(name string, options map[string]string) *LayuiPage {
	p.pageData.Select(name, options)
	return p
}

//时间戳格式化
func (p *LayuiPage) FormatTimestamp(name string, format string) *LayuiPage {
	p.pageData.FormatTimestamp(name, format)
	return p
}
func (p *LayuiPage) FormatDateTime(name string) *LayuiPage {
	return p.FormatTimestamp(name, atime.DATETIME)
}
func (p *LayuiPage) FormatDate(name string) *LayuiPage {
	return p.FormatTimestamp(name, atime.DATE)
}
func (p *LayuiPage) FormatTime(name string) *LayuiPage {
	return p.FormatTimestamp(name, atime.TIME)
}
func (p *LayuiPage) FormatMonth(name string) *LayuiPage {
	return p.FormatTimestamp(name, atime.MONTH)
}

//方法处理
func (p *LayuiPage) Func(name string, f func(value string, row map[string]string, rowSource map[string]string) string) *LayuiPage {
	p.pageData.Func(name, f)
	return p
}

//超链接
func (p *LayuiPage) URL(name string) *LayuiPage {
	p.pageData.URL(name)
	return p
}

//执行处理
func (p *LayuiPage) Exec() {
	p.pageData.Exec()
	p.Data = p.pageData.Data
	p.DataSource = p.pageData.DataSource
}