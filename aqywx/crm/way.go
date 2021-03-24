package crm

import (
	"encoding/json"
	"fmt"

	"github.com/asktop/gotools/aqywx/common"
	"github.com/asktop/gotools/aqywx/util"
)

const (
	addContactWayURL    = "https://qyapi.weixin.qq.com/cgi-bin/crm/add_contact_way?access_token=%s"
	getContactWayURL    = "https://qyapi.weixin.qq.com/cgi-bin/crm/get_contact_way?access_token=%s"
	updateContactWayURL = "https://qyapi.weixin.qq.com/cgi-bin/crm/update_contact_way?access_token=%s"
	delContactWayURL    = "https://qyapi.weixin.qq.com/cgi-bin/crm/del_contact_way?access_token=%s"
)

// AddContactWayResp 添加联系我方式的返回
type AddContactWayResp struct {
	util.CommonError
	ConfigID string `json:"config_id"`
}

// AddContactWay 添加联系我的方式
func (crm *Crm) AddContactWay(params common.ContactWay) (configID string, err error) {
	var accessToken string
	accessToken, err = crm.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(addContactWayURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, params)
	if err != nil {
		return
	}
	respData := new(AddContactWayResp)
	err = json.Unmarshal(response, respData)
	if err != nil {
		return
	}
	if respData.ErrCode != 0 {
		err = fmt.Errorf("AddContactWay Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}

	configID = respData.ConfigID
	return
}

// GetContactWayResp 获取已经配置的联系我的方式的返回
type GetContactWayResp struct {
	util.CommonError
	ContactWay []common.ContactWayItem `json:"contact_way"`
}

// GetContactWay 获取已经配置的联系我的方式
func (crm *Crm) GetContactWay(configID string) (contactWay []common.ContactWayItem, err error) {
	var accessToken string
	accessToken, err = crm.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getContactWayURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, map[string]string{
		"config_id": configID,
	})
	if err != nil {
		return
	}
	respData := new(GetContactWayResp)
	err = json.Unmarshal(response, respData)
	if err != nil {
		return
	}
	if respData.ErrCode != 0 {
		err = fmt.Errorf("AddContactWay Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}

	contactWay = respData.ContactWay
	return
}

// UpdateContactWay 更新联系我的方式
func (crm *Crm) UpdateContactWay(params common.ContactWay) (err error) {
	var accessToken string
	accessToken, err = crm.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(updateContactWayURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, params)
	if err != nil {
		return
	}
	respData := new(util.CommonError)
	err = json.Unmarshal(response, respData)
	if err != nil {
		return
	}
	if respData.ErrCode != 0 {
		err = fmt.Errorf("UpdateContactWay Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}

	return
}

// DelContactWay 删除联系我的方式
func (crm *Crm) DelContactWay(configID string) (err error) {
	var accessToken string
	accessToken, err = crm.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(delContactWayURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, map[string]string{
		"config_id": configID,
	})
	if err != nil {
		return
	}
	respData := new(util.CommonError)
	err = json.Unmarshal(response, respData)
	if err != nil {
		return
	}
	if respData.ErrCode != 0 {
		err = fmt.Errorf("DelContactWay Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}

	return
}
