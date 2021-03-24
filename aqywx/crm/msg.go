package crm

import (
	"encoding/json"
	"fmt"

	"github.com/asktop/gotools/aqywx/util"
)

const (
	addMsgTplURL         = "https://qyapi.weixin.qq.com/cgi-bin/crm/add_msg_template?access_token=%s"
	getGroupMsgResultURL = "https://qyapi.weixin.qq.com/cgi-bin/crm/get_group_msg_result?access_token=%s"
)

// MsgTemplate 消息模板
type MsgTemplate struct {
	ExternalUserID []string     `json:"external_userid"`
	Sender         string       `json:"sender"`
	Text           *Text        `json:"text"`
	Image          *Image       `json:"image"`
	Link           *Link        `json:"link"`
	Miniprogram    *Miniprogram `json:"miniprogram"`
}

// Text 消息文本内容
type Text struct {
	Content string `json:"content"`
}

// Image 图片数据
type Image struct {
	MediaID string `json:"media_id"`
}

// Link 链接数据
type Link struct {
	Title  string `json:"title"`
	PicURL string `json:"picurl"`
	Desc   string `json:"desc"`
	URL    string `json:"url"`
}

// Miniprogram 小程序数据
type Miniprogram struct {
	Title      string `json:"title"`
	PicMediaID string `json:"pic_media_id"`
	AppID      string `json:"appid"`
	Page       string `json:"page"`
}

// AddMsgTemplateResp 添加企业群发消息模板返回
type AddMsgTemplateResp struct {
	util.CommonError
	FailList []string `json:"fail_list"`
	MsgID    string   `json:"msgid"`
}

// AddMsgTemplate 添加企业群发消息模板
func (crm *Crm) AddMsgTemplate(params MsgTemplate) (respData *AddMsgTemplateResp, err error) {
	var accessToken string
	accessToken, err = crm.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(addMsgTplURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, params)
	if err != nil {
		return
	}
	respData = new(AddMsgTemplateResp)
	err = json.Unmarshal(response, respData)
	if err != nil {
		return
	}
	if respData.ErrCode != 0 {
		err = fmt.Errorf("AddMsgTemplate Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}

	return
}

// GetGroupMsgResultResp 获取企业群发消息发送结果返回
type GetGroupMsgResultResp struct {
	util.CommonError
	CheckStatus int                `json:"check_status"`
	DetailList  []ResultDetailItem `json:"detail_list"`
}

// ResultDetailItem 结果详情条目
type ResultDetailItem struct {
	ExternalUserID string `json:"external_userid"`
	UserID         string `json:"userid"`
	Status         int    `json:"status"`
	SendTime       int64  `json:"send_time"`
}

// GetGroupMsgResult 获取企业群发消息发送结果
func (crm *Crm) GetGroupMsgResult(msgID string) (respData *GetGroupMsgResultResp, err error) {
	var accessToken string
	accessToken, err = crm.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getGroupMsgResultURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, map[string]string{
		"msgid": msgID,
	})
	if err != nil {
		return
	}
	respData = new(GetGroupMsgResultResp)
	err = json.Unmarshal(response, respData)
	if err != nil {
		return
	}
	if respData.ErrCode != 0 {
		err = fmt.Errorf("GetGroupMsgResult Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}

	return
}
