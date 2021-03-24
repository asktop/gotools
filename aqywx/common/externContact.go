package common

// ExternalContact 外部联系人
type ExternalContact struct {
	ExternalUserID  string          `json:"external_userid"`  // 外部联系人的userid
	Name            string          `json:"name"`             // 外部联系人的姓名或别名
	Position        string          `json:"position"`         // 外部联系人的职位，如果外部企业或用户选择隐藏职位，则不返回，仅当联系人类型是企业微信用户时有此字段
	Avatar          string          `json:"avatar"`           // 外部联系人头像，第三方不可获取
	CorpName        string          `json:"corp_name"`        // 外部联系人所在企业的简称，仅当联系人类型是企业微信用户时有此字段
	CorpFullName    string          `json:"corp_full_name"`   // 外部联系人所在企业的主体名称，仅当联系人类型是企业微信用户时有此字段
	Type            int             `json:"type"`             // 外部联系人的类型，1表示该外部联系人是微信用户，2表示该外部联系人是企业微信用户
	Gender          int             `json:"gender"`           // 外部联系人性别 0-未知 1-男性 2-女性
	UnionID         string          `json:"unionid"`          // 外部联系人在微信开放平台的唯一身份标识（微信unionid），通过此字段企业可将外部联系人与公众号/小程序用户关联起来。仅当联系人类型是微信用户，且企业绑定了微信开发者ID有此字段。
	ExternalProfile ExternalProfile `json:"external_profile"` // 外部联系人的自定义展示信息，可以有多个字段和多种类型，包括文本，网页和小程序，仅当联系人类型是企业微信用户时有此字段
}

// FollowUser 添加了外部成员的企业成员
type FollowUser struct {
	UserID        string          `json:"userid"`         // 添加了此外部联系人的企业成员userid
	Remark        string          `json:"remark"`         // 该成员对此外部联系人的备注
	Description   string          `json:"description"`    // 该成员对此外部联系人的描述
	CreateTime    int64           `json:"createtime"`     // 该成员添加此外部联系人的时间
	Tags          []FollowUserTag `json:"tags"`           // 该成员添加此外部联系人所打标签
	RemarkCompany string          `json:"remark_company"` // 该成员对此客户备注的企业名称
	RemarkMobiles []string        `json:"remark_mobiles"` // 该成员对此客户备注的手机号码
	State         string          `json:"state"`          // 该成员添加此客户的渠道
}

// FollowUserTag 添加了外部成员的企业成员所打的标签
type FollowUserTag struct {
	GroupName string `json:"group_name"` // 该成员添加此外部联系人所打标签的分组名称（标签功能需要企业微信升级到2.7.5及以上版本）
	TagName   string `json:"tag_name"`   // 该成员添加此外部联系人所打标签名称
	Type      int    `json:"type"`       // 该成员添加此外部联系人所打标签类型, 1-企业设置, 2-用户自定义
}
