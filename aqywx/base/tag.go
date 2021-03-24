package base

import (
	"encoding/json"
	"fmt"

	"github.com/asktop/gotools/aqywx/util"
)

const (
	createTagURL   = "https://qyapi.weixin.qq.com/cgi-bin/tag/create?access_token=%s"
	updateTagURL   = "https://qyapi.weixin.qq.com/cgi-bin/tag/update?access_token=%s"
	deleteTagURL   = "https://qyapi.weixin.qq.com/cgi-bin/tag/delete?access_token=%s&tagid=%d"
	getTagUsersURL = "https://qyapi.weixin.qq.com/cgi-bin/tag/get?access_token=%s&tagid=%d"
	addTagUsersURL = "https://qyapi.weixin.qq.com/cgi-bin/tag/addtagusers?access_token=%s"
	delTagUsersURL = "https://qyapi.weixin.qq.com/cgi-bin/tag/deltagusers?access_token=%s"
	getTagListURL  = "https://qyapi.weixin.qq.com/cgi-bin/tag/list?access_token=%s"
)

// Tag 标签
type Tag struct {
	TagName string `json:"tagname"`
	TagID   int    `json:"tagid"`
}

// CreateTagResp 创建标签返回
type CreateTagResp struct {
	util.CommonError
	TagID int `json:"tagid"`
}

// CreateTag 创建标签
func (base *Base) CreateTag(tagData Tag) (ret int, err error) {
	var accessToken string
	accessToken, err = base.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(createTagURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, tagData)
	if err != nil {
		return
	}
	respData := new(CreateTagResp)
	err = json.Unmarshal(response, respData)
	if err != nil {
		return
	}
	if respData.ErrCode != 0 {
		err = fmt.Errorf("CreateTag Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}

	ret = respData.TagID
	return
}

// UpdateTag 更新标签名字
func (base *Base) UpdateTag(tagData Tag) (err error) {
	var accessToken string
	accessToken, err = base.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(updateTagURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, tagData)
	if err != nil {
		return
	}
	respData := new(util.CommonError)
	err = json.Unmarshal(response, respData)
	if err != nil {
		return
	}
	if respData.ErrCode != 0 {
		err = fmt.Errorf("UpdateTag Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}

	return
}

// DeleteTag 删除标签
func (base *Base) DeleteTag(tagID int) (err error) {
	var accessToken string
	accessToken, err = base.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(deleteTagURL, accessToken, tagID)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	respData := new(util.CommonError)
	err = json.Unmarshal(response, respData)
	if err != nil {
		return
	}
	if respData.ErrCode != 0 {
		err = fmt.Errorf("DeleteTag Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}

	return
}

// GetTagUsersResp 获取标签成员返回
type GetTagUsersResp struct {
	util.CommonError
	TagName   string       `json:"tagname"`
	UserList  []UserSample `json:"userlist"`
	PartyList []int        `json:"partylist"`
}

// GetTagUsers 获取标签成员
func (base *Base) GetTagUsers(tagID int) (ret []UserSample, err error) {
	var accessToken string
	accessToken, err = base.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getTagUsersURL, accessToken, tagID)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	respData := new(GetTagUsersResp)
	err = json.Unmarshal(response, respData)
	if err != nil {
		return
	}
	if respData.ErrCode != 0 {
		err = fmt.Errorf("GetTagUsers Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}

	ret = respData.UserList
	return
}

// CommonTagUsersParams 增加和删除标签用户参数
type CommonTagUsersParams struct {
	TagID     int      `json:"tagid"`
	UserList  []string `json:"userlist"`
	PartyList []int    `json:"partylist"`
}

// CommonTagUsersResp 添加和删除标签用户返回
type CommonTagUsersResp struct {
	util.CommonError
	InvalidList  string `json:"invalidlist"`
	InvalidParty []int  `json:"invalidparty"`
}

// AddTagUsers 增加标签成员
func (base *Base) AddTagUsers(params CommonTagUsersParams) (respData *CommonTagUsersResp, err error) {
	var accessToken string
	accessToken, err = base.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(addTagUsersURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, params)
	if err != nil {
		return
	}
	respData = new(CommonTagUsersResp)
	err = json.Unmarshal(response, respData)
	if err != nil {
		return
	}
	if respData.ErrCode != 0 {
		err = fmt.Errorf("AddTagUsers Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}

	return
}

// DelTagUsers 删除标签成员
func (base *Base) DelTagUsers(params CommonTagUsersParams) (respData *CommonTagUsersResp, err error) {
	var accessToken string
	accessToken, err = base.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(delTagUsersURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, params)
	if err != nil {
		return
	}
	respData = new(CommonTagUsersResp)
	err = json.Unmarshal(response, respData)
	if err != nil {
		return
	}
	if respData.ErrCode != 0 {
		err = fmt.Errorf("DelTagUsers Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}

	return
}

// GetTagListResp 获取标签列表返回
type GetTagListResp struct {
	util.CommonError
	TagList []Tag `json:"taglist"`
}

// GetTagList 获取标签列表
func (base *Base) GetTagList() (ret []Tag, err error) {
	var accessToken string
	accessToken, err = base.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getTagListURL, accessToken)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	respData := new(GetTagListResp)
	err = json.Unmarshal(response, respData)
	if err != nil {
		return
	}
	if respData.ErrCode != 0 {
		err = fmt.Errorf("GetTagList Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}

	ret = respData.TagList
	return
}
