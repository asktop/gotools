package extra

import (
    "github.com/asktop/gotools/acast"
    "strings"
    "sync"
)

//常量盒子
type Box struct {
    mu   sync.RWMutex
    boxs []BoxOne
}

type BoxOne struct {
    Value   string `json:"value"`   //值
    Title   string `json:"title"`   //展示名
    Checked bool   `json:"checked"` //是否选中
}

func NewBox(boxs []BoxOne) *Box {
    b := new(Box)
    b.boxs = boxs
    return b
}

//在前面追加元素
func (b *Box) PushFront(boxone BoxOne) *Box {
    b.mu.Lock()
    newboxs := []BoxOne{}
    newboxs = append(newboxs, boxone)
    for _, bo := range b.boxs {
        newboxs = append(newboxs, bo)
    }
    b.boxs = newboxs
    b.mu.Unlock()
    return b
}

//在后面追加元素
func (b *Box) PushBack(boxone BoxOne) *Box {
    b.mu.Lock()
    b.boxs = append(b.boxs, boxone)
    b.mu.Unlock()
    return b
}

//临时在前面追加元素
func (b *Box) AddFront(boxone BoxOne) *Box {
    nb := new(Box)
    nb.boxs = append(nb.boxs, boxone)
    for _, bo := range b.boxs {
        nb.boxs = append(nb.boxs, bo)
    }
    return nb
}

//临时在后面追加元素
func (b *Box) AddBack(boxone BoxOne) *Box {
    nb := new(Box)
    nb.boxs = b.boxs
    nb.boxs = append(nb.boxs, boxone)
    return nb
}

//获取选中常量列表（有序）
func (b *Box) GetCheckedBox(checkedVs ...string) []BoxOne {
    if len(checkedVs) > 0 {
        var allCheckedVs []string
        for _, checkId := range checkedVs {
            ids := strings.Split(checkId, ",")
            allCheckedVs = append(allCheckedVs, ids...)
        }
        for i, box := range b.boxs {
            var checked bool
            for _, v := range allCheckedVs {
                if v == box.Value {
                    checked = true
                    break
                }
            }
            box.Checked = checked
            b.boxs[i] = box
        }
    }
    return b.boxs
}

//获取常量列表（无序）
func (b *Box) GetMap() map[string]string {
    rs := map[string]string{}
    for _, box := range b.boxs {
        rs[box.Value] = box.Title
    }
    return rs
}

//获取常量列表（无序）
func (b *Box) GetMapInt64() map[int64]string {
    rs := map[int64]string{}
    for _, box := range b.boxs {
        rs[acast.ToInt64(box.Value)] = box.Title
    }
    return rs
}

//获取单个常量
func (b *Box) GetExist(value string) *BoxOne {
    for _, box := range b.boxs {
        b := box
        if b.Value == value {
            return &b
        }
    }
    return nil
}

//获取Switch开关HTML
func (b *Box) GetSwitchHtml(dataId string, name string, value string, checkedValue string) string {
    var html, text, checked string
    for _, bo := range b.boxs {
        text += bo.Title + "|"
    }
    text = strings.TrimSuffix(text, "|")
    if value == checkedValue {
        checked = "checked"
    }
    html = `<input type="checkbox" id="` + dataId + `" name="` + name + `" value="` + checkedValue + `" lay-filter="switch" lay-skin="switch" lay-text="` + text + `" ` + checked + `>`
    return html
}
