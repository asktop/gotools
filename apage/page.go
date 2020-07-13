package apage

import (
	"errors"
	"github.com/asktop/dbr"
	"github.com/asktop/gotools/acast"
	"github.com/asktop/gotools/atime"
	"math"
	"reflect"
	"strings"
)

const (
	DEFAULT         = "DEFAULT"         //默认值
	SELECT          = "SELECT"          //下拉选项处理
	FORMATTIMESTAMP = "FORMATTIMESTAMP" //时间戳格式化
	FUNC            = "FUNC"            //方法处理
	URL             = "URL"             //超链接
)

type Page struct {
	Page       uint64 `json:"page"`              //当前页码
	Limit      uint64 `json:"limit"`             //每页条数
	Total      int64  `json:"total"`             //总条数
	TotalPage  int64                             //总页数
	Data       []map[string]string `json:"data"` //分页展示数据
	DataSource []map[string]string               //分页原始数据
	execFields []string                          //自定义处理的所有字段
	execValues map[string]interface{}            //自定义处理的所有字段的处理值
}

func NewPage(stmt *dbr.SelectStmt, page uint64, limit uint64, defLimit ...uint64) (p *Page, err error) {
	p = new(Page)
	if page <= 0 {
		page = 1
	}
	p.Page = page
	if limit <= 0 {
		if len(defLimit) > 0 {
			limit = defLimit[0]
		} else {
			limit = 10
		}
	}
	p.Limit = limit
	p.execValues = map[string]interface{}{}
	//查询总条数
	count, err := stmt.Count()
	if err != nil {
		return p, err
	}
	p.Total = int64(count)
	p.TotalPage = acast.ToInt64(math.Ceil(float64(p.Total) / float64(p.Limit)))
	//查询分页数据
	data := []map[string]string{}
	_, err = stmt.Paginate(page, limit).Load(&data)
	if err != nil {
		return p, err
	}
	p.Data = data
	return p, nil
}

//默认值
func (p *Page) Default(name string, defaultVal string) *Page {
	key := name + " " + DEFAULT
	if _, ok := p.execValues[key]; !ok {
		p.execFields = append(p.execFields, key)
	}
	p.execValues[key] = defaultVal
	return p
}

//下拉选项处理
func (p *Page) Select(name string, options map[string]string) *Page {
	key := name + " " + SELECT
	if _, ok := p.execValues[key]; !ok {
		p.execFields = append(p.execFields, key)
	}
	p.execValues[key] = options
	return p
}

//时间戳格式化
func (p *Page) FormatTimestamp(name string, format string) *Page {
	key := name + " " + FORMATTIMESTAMP
	if _, ok := p.execValues[key]; !ok {
		p.execFields = append(p.execFields, key)
	}
	p.execValues[key] = format
	return p
}
func (p *Page) FormatDateTime(name string) *Page {
	return p.FormatTimestamp(name, atime.DATETIME)
}
func (p *Page) FormatDate(name string) *Page {
	return p.FormatTimestamp(name, atime.DATE)
}
func (p *Page) FormatTime(name string) *Page {
	return p.FormatTimestamp(name, atime.TIME)
}
func (p *Page) FormatMonth(name string) *Page {
	return p.FormatTimestamp(name, atime.MONTH)
}

//方法处理
func (p *Page) Func(name string, f func(value string, row map[string]string, rowSource map[string]string) string) *Page {
	key := name + " " + FUNC
	if _, ok := p.execValues[key]; !ok {
		p.execFields = append(p.execFields, key)
	}
	p.execValues[key] = f
	return p
}

//超链接
func (p *Page) URL(name string) *Page {
	key := name + " " + URL
	if _, ok := p.execValues[key]; !ok {
		p.execFields = append(p.execFields, key)
	}
	p.execValues[key] = ""
	return p
}

//执行处理
func (p *Page) Exec() *Page {
	for i, data := range p.Data {
		//原始数据备份
		dataSource := map[string]string{}
		for k, v := range data {
			dataSource[k] = v
		}
		p.DataSource = append(p.DataSource, dataSource)
		//处理业务数据
		for _, key := range p.execFields {
			keys := strings.Split(key, " ")
			name := keys[0]     //字段名
			value := data[name] //字段值
			execType := keys[1]
			exec := p.execValues[key]
			switch execType {
			case DEFAULT:
				execVal := exec.(string)
				if value == "" {
					value = execVal
				}
			case SELECT:
				if execVal, ok := exec.(map[string]string); ok {
					value = execVal[value]
				}
			case FORMATTIMESTAMP:
				execVal := exec.(string)
				if value != "" && value != "0" {
					value = atime.FormatTimestamp(execVal, value)
				} else if value == "0" {
					value = ""
				}
			case FUNC:
				execVal := exec.(func(value string, row map[string]string, rowSource map[string]string) string)
				value = execVal(value, data, dataSource)
			case URL:
				value = `<a href="` + value + `" target="_blank">` + value + `</a>`
			}
			data[name] = value
		}
		p.Data[i] = data
	}
	return p
}

//将Paginator数据封装到结构中
func (p *Page) LoadData(value interface{}) error {
	//封装数据
	if len(p.Data) == 0 {
		return nil
	}

	data := reflect.ValueOf(value)
	if data.Kind() == reflect.Ptr {
		data = data.Elem()
	}
	if data.IsValid() {
		if data.Kind() == reflect.Slice {
			dataPtr := reflect.New(data.Type())
			err := acast.MapsStrToStructs(p.Data, dataPtr.Interface())
			if err != nil {
				return err
			}
			data.Set(dataPtr.Elem())
			return nil
		}
	}
	return errors.New("Page.LoadData : no field Data or Data isn't slice")
}

func (p *Page) LoadDataSource(value interface{}) error {
	//封装数据
	if len(p.DataSource) == 0 {
		return nil
	}

	data := reflect.ValueOf(value)
	if data.Kind() == reflect.Ptr {
		data = data.Elem()
	}
	if data.IsValid() {
		if data.Kind() == reflect.Slice {
			dataPtr := reflect.New(data.Type())
			err := acast.MapsStrToStructs(p.DataSource, dataPtr.Interface())
			if err != nil {
				return err
			}
			data.Set(dataPtr.Elem())
			return nil
		}
	}
	return errors.New("Page.LoadData : no field DataSource or DataSource isn't slice")
}

func (p *Page) GetPageLimit() (page, limit uint64, totol int64) {
	return p.Page, p.Limit, p.Total
}

//分页请求体
type PageReq struct {
	//当前页码
	Page uint64 `json:"page"`
	//每页条数
	Limit uint64 `json:"limit"`
}

//分页响应体
type PageRes struct {
	//当前页码
	Page uint64 `json:"page"`
	//每页条数
	Limit uint64 `json:"limit"`
	//总条数
	Total int64 `json:"total"`
}

func (p *PageRes) SetPageLimit(page, limit uint64, total int64) {
	p.Page = page
	p.Limit = limit
	p.Total = total
}
