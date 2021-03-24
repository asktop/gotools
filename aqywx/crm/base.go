package crm

import (
	"encoding/json"
	"fmt"

	"github.com/asktop/gotools/aqywx/common"
	"github.com/asktop/gotools/aqywx/util"
)

const (
	getCustomerContactsURL    = "https://qyapi.weixin.qq.com/cgi-bin/crm/get_customer_contacts?access_token=%s"
	getExternalContactListURL = "https://qyapi.weixin.qq.com/cgi-bin/crm/get_external_contact_list?access_token=%s&userid=%s"
	getExternalContactURL     = "https://qyapi.weixin.qq.com/cgi-bin/crm/get_external_contact?access_token=%s&external_userid=%s"
)

// GetCustomerContactsResp 获取配置了客户联系功能的成员列表返回
type GetCustomerContactsResp struct {
	util.CommonError
	CustomerContacts []string `json:"customer_contacts"`
}

// GetCustomerContacts 获取配置了客户联系功能的成员列表
func (crm *Crm) GetCustomerContacts() (ret []string, err error) {
	var accessToken string
	accessToken, err = crm.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getCustomerContactsURL, accessToken)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	respData := new(GetCustomerContactsResp)
	err = json.Unmarshal(response, respData)
	if err != nil {
		return
	}
	if respData.ErrCode != 0 {
		err = fmt.Errorf("GetCustomerContacts Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}

	ret = respData.CustomerContacts
	return
}

// GetExternalContactListResp 获取外部联系人列表返回
type GetExternalContactListResp struct {
	util.CommonError
	ExternalUserIDs []string `json:"external_userid"`
}

// GetExternalContactList 获取外部联系人列表
func (crm *Crm) GetExternalContactList(userID string) (ret []string, err error) {
	var accessToken string
	accessToken, err = crm.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getExternalContactListURL, accessToken, userID)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	respData := new(GetExternalContactListResp)
	err = json.Unmarshal(response, respData)
	if err != nil {
		return
	}
	if respData.ErrCode != 0 {
		err = fmt.Errorf("GetCustomerContacts Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}

	ret = respData.ExternalUserIDs
	return
}

// GetExternalContactResp 获取外部联系人详情返回
type GetExternalContactResp struct {
	util.CommonError
	ExternalContact common.ExternalContact `json:"external_contact"`
	FollowUsers     []common.FollowUser    `json:"follow_user"`
}

// GetExternalContact 获取外部联系人详情
func (crm *Crm) GetExternalContact(externalUserID string) (respData *GetExternalContactResp, err error) {
	var accessToken string
	accessToken, err = crm.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getExternalContactURL, accessToken, externalUserID)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	respData = new(GetExternalContactResp)
	err = json.Unmarshal(response, respData)
	if err != nil {
		return
	}
	if respData.ErrCode != 0 {
		err = fmt.Errorf("GetExternalContact Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}
	return
}
