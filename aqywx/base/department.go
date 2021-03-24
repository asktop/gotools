package base

import (
	"encoding/json"
	"fmt"

	"github.com/asktop/gotools/aqywx/util"
)

const (
	createDepURL  = "https://qyapi.weixin.qq.com/cgi-bin/department/create?access_token=%s"
	updateDepURL  = "https://qyapi.weixin.qq.com/cgi-bin/department/update?access_token=%s"
	deleteDepURL  = "https://qyapi.weixin.qq.com/cgi-bin/department/delete?access_token=%s&id=%d"
	getDepListURL = "https://qyapi.weixin.qq.com/cgi-bin/department/list?access_token=%s&id=%v"
)

// Department 部门
type Department struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ParentID int    `json:"parentid"`
	Order    int    `json:"order"`
}

// CreateDepartmentResp 创建部门返回
type CreateDepartmentResp struct {
	util.CommonError
	ID int `json:"id"`
}

// CreateDepartment 创建部门
func (base *Base) CreateDepartment(depData Department) (ret int, err error) {
	var accessToken string
	accessToken, err = base.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(createDepURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, depData)
	if err != nil {
		return
	}
	respData := new(CreateDepartmentResp)
	err = json.Unmarshal(response, respData)
	if err != nil {
		return
	}
	if respData.ErrCode != 0 {
		err = fmt.Errorf("CreateDepartment Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}

	ret = respData.ID
	return
}

// UpdateDepartment 更新部门
func (base *Base) UpdateDepartment(depData Department) (err error) {
	var accessToken string
	accessToken, err = base.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(updateDepURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, depData)
	if err != nil {
		return
	}
	respData := new(util.CommonError)
	err = json.Unmarshal(response, respData)
	if err != nil {
		return
	}
	if respData.ErrCode != 0 {
		err = fmt.Errorf("UpdateDepartment Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}

	return
}

// DeleteDepartment 删除部门
func (base *Base) DeleteDepartment(depID int) (err error) {
	var accessToken string
	accessToken, err = base.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(deleteDepURL, accessToken, depID)
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
		err = fmt.Errorf("DeleteDepartment Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}
	return
}

// GetDepartmentListResp 获取部门列表返回
type GetDepartmentListResp struct {
	util.CommonError
	Department []Department `json:"department"`
}

// GetDepartmentList 获取部门列表
func (base *Base) GetDepartmentList(depIDs ...int) (ret []Department, err error) {
	var accessToken string
	accessToken, err = base.GetAccessToken()
	if err != nil {
		return
	}

	uri := ""
	if len(depIDs) > 0 {
		uri = fmt.Sprintf(getDepListURL, accessToken, depIDs[0])
	} else {
		uri = fmt.Sprintf(getDepListURL, accessToken, "")
	}
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	respData := new(GetDepartmentListResp)
	err = json.Unmarshal(response, respData)
	if err != nil {
		return
	}
	if respData.ErrCode != 0 {
		err = fmt.Errorf("GetDepartmentList Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}

	ret = respData.Department
	return
}
