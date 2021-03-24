package base

// ExtAttr 扩展属性
type ExtAttr struct {
	Attrs []Attr `json:"attrs"`
}

// Attr 属性
type Attr struct {
	Type        int             `json:"type"`
	Name        string          `json:"name"`
	Text        TextAttr        `json:"text"`
	Web         WebAttr         `json:"web"`
	MiniProgram MiniProgramAttr `json:"miniprogram"`
}

// TextAttr 文本属性
type TextAttr struct {
	Value string `json:"value"`
}

// WebAttr 网页属性
type WebAttr struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}

// MiniProgramAttr 小程序属性
type MiniProgramAttr struct {
	AppID    string `json:"appid"`
	Title    string `json:"title"`
	PagePath string `json:"pagepath"`
}

// ExternalProfile 扩展资料
type ExternalProfile struct {
	ExternalCorpName string  `json:"external_corp_name"`
	ExternalAttr     ExtAttr `json:"external_attr"`
}
