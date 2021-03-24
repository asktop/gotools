package base

import (
	"encoding/json"
	"fmt"

	"github.com/asktop/gotools/aqywx/util"
)

const (
	getUserURL              = "https://qyapi.weixin.qq.com/cgi-bin/user/get?access_token=%s&userid=%s"
	createUserURL           = "https://qyapi.weixin.qq.com/cgi-bin/user/create?access_token=%s"
	updateUserURL           = "https://qyapi.weixin.qq.com/cgi-bin/user/update?access_token=%s"
	deleteUserURL           = "https://qyapi.weixin.qq.com/cgi-bin/user/delete?access_token=%s&userid=%s"
	batchDeleteUserURL      = "https://qyapi.weixin.qq.com/cgi-bin/user/batchdelete?access_token=%s"
	getDepUserSimpleListURL = "https://qyapi.weixin.qq.com/cgi-bin/user/simplelist?access_token=%s&department_id=%v&fetch_child=%v"
	getDepUserListURL       = "https://qyapi.weixin.qq.com/cgi-bin/user/list?access_token=%s&department_id=%v&fetch_child=%v"
	userIDToOpenIDURL       = "https://qyapi.weixin.qq.com/cgi-bin/user/convert_to_openid?access_token=%s"
	openIDToUserIDURL       = "https://qyapi.weixin.qq.com/cgi-bin/user/convert_to_userid?access_token=%s"
	userAuthSuccessURL      = "https://qyapi.weixin.qq.com/cgi-bin/user/authsucc?access_token=%s&userid=%s"
	inviteUserURL           = "https://qyapi.weixin.qq.com/cgi-bin/batch/invite?access_token=%s"
)

// User 用户数据
type User struct {
	UserID           string          `json:"userid"`            // 成员UserID
	Name             string          `json:"name"`              // 成员名称
	Department       []int           `json:"department"`        // 所属部门id列表
	Order            []int           `json:"order"`             // 部门内排序值
	Position         string          `json:"position"`          // 职务信息
	Mobile           string          `json:"mobile"`            // 手机号码
	Gender           string          `json:"gender"`            // 性别，0未定义，1男性，2女性
	Email            string          `json:"email"`             // 邮箱
	IsLeaderInDept   []int           `json:"is_leader_in_dept"` // 表示所在部门内是否是上级
	Avatar           string          `json:"avatar"`            // 头像
	Telephone        string          `json:"telephone"`         // 座机
	Enable           int             `json:"enable"`            // 成员启用状态，1启用，0禁用
	Alias            string          `json:"alias"`             // 别名
	Address          string          `json:"address"`           // 地址
	ExtAttr          ExtAttr         `json:"extattr"`           // 扩展属性
	Status           int             `json:"status"`            // 激活状态，1已激活，2已禁用，4未激活
	QRCode           string          `json:"qr_code"`           // 员工个人二维码
	ExternalPosition string          `json:"external_position"` // 成员对外职务
	ExternalProfile  ExternalProfile `json:"external_profile"`  // 成员对外属性
}

// UserData 创建用户时传输的数据
type UserData struct {
	User
	AvatarMediaid string `json:"avatar_mediaid"`
	ToInvite      bool   `json:"to_invite"`
}

// CreateUser 创建用户
func (base *Base) CreateUser(userData UserData) (err error) {
	var accessToken string
	accessToken, err = base.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(createUserURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, userData)
	if err != nil {
		return
	}
	respData := new(util.CommonError)
	err = json.Unmarshal(response, respData)
	if err != nil {
		return
	}
	if respData.ErrCode != 0 {
		err = fmt.Errorf("CreateUser Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}
	return
}

// UpdateUser 更新用户
func (base *Base) UpdateUser(userData UserData) (err error) {
	var accessToken string
	accessToken, err = base.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(updateUserURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, userData)
	if err != nil {
		return
	}
	respData := new(util.CommonError)
	err = json.Unmarshal(response, respData)
	if err != nil {
		return
	}
	if respData.ErrCode != 0 {
		err = fmt.Errorf("UpdateUser Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}
	return
}

// GetUserResp 获取用户信息返回
type GetUserResp struct {
	util.CommonError
	User
}

// GetUser 获取用户信息
func (base *Base) GetUser(userID string) (ret User, err error) {
	var accessToken string
	accessToken, err = base.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getUserURL, accessToken, userID)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	respData := new(GetUserResp)
	err = json.Unmarshal(response, respData)
	if err != nil {
		return
	}
	if respData.ErrCode != 0 {
		err = fmt.Errorf("GetUser Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}
	ret = respData.User
	return
}

// DeleteUser 删除用户
func (base *Base) DeleteUser(userID string) (err error) {
	var accessToken string
	accessToken, err = base.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(deleteUserURL, accessToken, userID)
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
		err = fmt.Errorf("DeleteUser Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}
	return
}

// BatchDeleteUserParams 批量删除用户参数
type BatchDeleteUserParams struct {
	UserIDList []string `json:"useridlist"`
}

// BatchDeleteUser 批量删除用户
func (base *Base) BatchDeleteUser(userIDs []string) (err error) {
	var accessToken string
	accessToken, err = base.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(batchDeleteUserURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, BatchDeleteUserParams{
		UserIDList: userIDs,
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
		err = fmt.Errorf("BatchDeleteUser Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}
	return
}

// UserSample 用户基础信息
type UserSample struct {
	UserID     string `json:"userid"`
	Name       string `json:"name"`
	Department []int  `json:"department"`
}

// GetDepUserSimpleListResp 获取部门成员基础信息返回
type GetDepUserSimpleListResp struct {
	util.CommonError
	UserList []UserSample `json:"userlist"`
}

// GetDepUserSimpleList 获取部门成员基础信息
func (base *Base) GetDepUserSimpleList(depID int, fetchChild bool) (ret []UserSample, err error) {
	var accessToken string
	accessToken, err = base.GetAccessToken()
	if err != nil {
		return
	}

	fetchChildFlag := 0
	if fetchChild {
		fetchChildFlag = 1
	}

	uri := fmt.Sprintf(getDepUserSimpleListURL, accessToken, depID, fetchChildFlag)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	respData := new(GetDepUserSimpleListResp)
	err = json.Unmarshal(response, respData)
	if err != nil {
		return
	}
	if respData.ErrCode != 0 {
		err = fmt.Errorf("GetDepUserSimpleList Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}

	ret = respData.UserList
	return
}

// GetDepUserListResp 获取部门成员详细信息返回
type GetDepUserListResp struct {
	util.CommonError
	UserList []User `json:"userlist"`
}

// GetDepUserList 获取部门成员详细信息
func (base *Base) GetDepUserList(depID int, fetchChild bool) (ret []User, err error) {
	var accessToken string
	accessToken, err = base.GetAccessToken()
	if err != nil {
		return
	}

	fetchChildFlag := 0
	if fetchChild {
		fetchChildFlag = 1
	}

	uri := fmt.Sprintf(getDepUserListURL, accessToken, depID, fetchChildFlag)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	respData := new(GetDepUserListResp)
	err = json.Unmarshal(response, respData)
	if err != nil {
		return
	}
	if respData.ErrCode != 0 {
		err = fmt.Errorf("GetDepUserList Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}

	ret = respData.UserList
	return
}

// UserIDToOpenIDResp userid转openid返回
type UserIDToOpenIDResp struct {
	util.CommonError
	OpenID string `json:"openid"`
}

// UserIDToOpenID userid转openid
func (base *Base) UserIDToOpenID(userID string) (ret string, err error) {
	var accessToken string
	accessToken, err = base.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(userIDToOpenIDURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, map[string]string{
		"userid": userID,
	})
	if err != nil {
		return
	}
	respData := new(UserIDToOpenIDResp)
	err = json.Unmarshal(response, respData)
	if err != nil {
		return
	}
	if respData.ErrCode != 0 {
		err = fmt.Errorf("UserIDToOpenID Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}

	ret = respData.OpenID
	return
}

// OpenIDToUserIDResp openid转userid返回
type OpenIDToUserIDResp struct {
	util.CommonError
	UserID string `json:"userid"`
}

// OpenIDToUserID openid转userid
func (base *Base) OpenIDToUserID(openID string) (ret string, err error) {
	var accessToken string
	accessToken, err = base.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(openIDToUserIDURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, map[string]string{
		"openid": openID,
	})
	if err != nil {
		return
	}
	respData := new(OpenIDToUserIDResp)
	err = json.Unmarshal(response, respData)
	if err != nil {
		return
	}
	if respData.ErrCode != 0 {
		err = fmt.Errorf("OpenIDToUserID Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}

	ret = respData.UserID
	return
}

// AuthUserSuccess 用户二次验证
func (base *Base) AuthUserSuccess(userID string) (err error) {
	var accessToken string
	accessToken, err = base.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(userAuthSuccessURL, accessToken, userID)
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
		err = fmt.Errorf("GetDepUserList Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}

	return
}

// InviteUserParams 邀请用户参数
type InviteUserParams struct {
	User  []string `json:"user"`
	Party []int    `json:"party"`
	Tag   []int    `json:"tag"`
}

// InviteUserResp 邀请用户返回
type InviteUserResp struct {
	util.CommonError
	InvalidUser  []string `json:"invaliduser"`
	InvalidParty []string `json:"invalidparty"`
	InvalidTag   []string `json:"invalidtag"`
}

// InviteUser 邀请用户
func (base *Base) InviteUser(params InviteUserParams) (respData *InviteUserResp, err error) {
	var accessToken string
	accessToken, err = base.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(inviteUserURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, params)
	if err != nil {
		return
	}
	respData = new(InviteUserResp)
	err = json.Unmarshal(response, respData)
	if err != nil {
		return
	}
	if respData.ErrCode != 0 {
		err = fmt.Errorf("InviteUser Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}

	return
}
