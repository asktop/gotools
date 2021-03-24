package common

// ContactWay 联系我的方式
type ContactWay struct {
	ConfigID   string   `json:"config_id"`
	Type       int      `json:"type"`
	Scene      int      `json:"scene"`
	Style      int      `json:"style"`
	Remark     string   `json:"remark"`
	SkipVerify bool     `json:"skip_verify"`
	State      string   `json:"state"`
	User       []string `json:"user"`
	Party      []int    `json:"party"`
}

// ContactWayItem 联系我的方式的返回结果
type ContactWayItem struct {
	ConfigID   string `json:"config_id"`
	Type       int    `json:"type"`
	Scene      int    `json:"scene"`
	Style      int    `json:"style"`
	Remark     string `json:"remark"`
	SkipVerify bool   `json:"skip_verify"`
	State      string `json:"state"`
	QRCode     string `json:"qr_code"`
}
