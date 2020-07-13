package apage

import (
	"errors"
	"github.com/asktop/dbr"
	"github.com/asktop/gotools/acast"
	"github.com/asktop/gotools/atime"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

type BsPage struct {
	Page       uint64              `json:"page"`        //当前页码
	Limit      uint64              `json:"limit"`       //每页条数
	Total      int64               `json:"total"`       //总条数
	TotalPage  int64               `json:"total_page"`  //总页数
	Data       []map[string]string `json:"data"`        //分页展示数据
	DataSource []map[string]string `json:"data_source"` //分页原始数据
	request    *http.Request                            //请求request
	pageData   *Page
	pages      []uint64 //分页标签 当前显示页码枚举
}

//bootstrap
func NewBsPage(req *http.Request, stmt *dbr.SelectStmt, defLimit ...uint64) (p *BsPage, err error) {
	p = new(BsPage)
	if req == nil {
		return nil, errors.New("Request can not be nil")
	}
	p.request = req
	page := acast.ToUint64(p.request.Form.Get("page"))
	limit := acast.ToUint64(p.request.Form.Get("limit"))
	p.pageData, err = NewPage(stmt, page, limit, defLimit...)
	p.Page = p.pageData.Page
	p.Limit = p.pageData.Limit
	p.Total = p.pageData.Total
	p.TotalPage = p.pageData.TotalPage
	if err != nil {
		return p, err
	} else {
		return p, nil
	}
}

//默认值
func (p *BsPage) Default(name string, defaultVal string) *BsPage {
	p.pageData.Default(name, defaultVal)
	return p
}

//下拉选项处理
func (p *BsPage) Select(name string, options map[string]string) *BsPage {
	p.pageData.Select(name, options)
	return p
}

//时间戳格式化
func (p *BsPage) FormatTimestamp(name string, format string) *BsPage {
	p.pageData.FormatTimestamp(name, format)
	return p
}
func (p *BsPage) FormatDateTime(name string) *BsPage {
	return p.FormatTimestamp(name, atime.DATETIME)
}
func (p *BsPage) FormatDate(name string) *BsPage {
	return p.FormatTimestamp(name, atime.DATE)
}
func (p *BsPage) FormatTime(name string) *BsPage {
	return p.FormatTimestamp(name, atime.TIME)
}
func (p *BsPage) FormatMonth(name string) *BsPage {
	return p.FormatTimestamp(name, atime.MONTH)
}

//方法处理
func (p *BsPage) Func(name string, f func(value string, row map[string]string, rowSource map[string]string) string) *BsPage {
	p.pageData.Func(name, f)
	return p
}

//超链接
func (p *BsPage) URL(name string) *BsPage {
	p.pageData.URL(name)
	return p
}

//执行处理
func (p *BsPage) Exec() {
	p.pageData.Exec()
	p.Data = p.pageData.Data
	p.DataSource = p.pageData.DataSource
}

//去指定页码
func (p *BsPage) GoPage(page uint64) (link string) {
	requestUrl, _ := url.ParseRequestURI(p.request.RequestURI)
	values := requestUrl.Query()
	if page == 1 {
		values.Del("page")
	} else {
		values.Set("page", strconv.Itoa(int(page)))
	}
	requestUrl.RawQuery = values.Encode()
	link = requestUrl.String()
	return
}

//是否有分页（页码总数大于1）
func (p *BsPage) HasPage() bool {
	return p.TotalPage > 1
}

//去首页
func (p *BsPage) GoFirst() (link string) {
	return p.GoPage(1)
}

//是否有上一页
func (p *BsPage) HasPrev() bool {
	return p.Page > 1
}

//去上一页
func (p *BsPage) GoPrev() (link string) {
	if p.HasPrev() {
		link = p.GoPage(p.Page - 1)
	}
	return
}

//当前显示的页码集合
func (p *BsPage) Pages() []uint64 {
	if p.pages == nil && p.Total > 0 {
		var pages []uint64
		var totalPage uint64 = uint64(p.TotalPage)
		page := p.Page
		switch {
		case page >= totalPage-4 && totalPage > 9:
			start := totalPage - 9 + 1
			pages = make([]uint64, 9)
			for i, _ := range pages {
				pages[i] = start + uint64(i)
			}
		case page >= 5 && totalPage > 9:
			start := page - 5 + 1
			pages = make([]uint64, uint64(math.Min(9, float64(page+4+1))))
			for i, _ := range pages {
				pages[i] = start + uint64(i)
			}
		default:
			pages = make([]uint64, uint64(math.Min(9, float64(totalPage))))
			for i, _ := range pages {
				pages[i] = uint64(i) + 1
			}
		}
		p.pages = pages
	}
	return p.pages
}

//当前页是否是指定页码
func (p *BsPage) IsPage(page uint64) bool {
	return p.Page == page
}

//是否有下一页
func (p *BsPage) HasNext() bool {
	return p.Page < uint64(p.TotalPage)
}

//去下一页
func (p *BsPage) GoNext() (link string) {
	if p.HasNext() {
		link = p.GoPage(p.Page + 1)
	}
	return
}

//去尾页
func (p *BsPage) GoLast() (link string) {
	return p.GoPage(uint64(p.TotalPage))
}
