package base

import (
	"encoding/json"
	"fmt"

	"github.com/asktop/gotools/aqywx/util"
)

const (
	batchSyncUserURL    = "https://qyapi.weixin.qq.com/cgi-bin/batch/syncuser?access_token=%s"
	batchReplaceUserURL = "https://qyapi.weixin.qq.com/cgi-bin/batch/replaceuser?access_token=%s"
	batchReplaceDeptURL = "https://qyapi.weixin.qq.com/cgi-bin/batch/replaceparty?access_token=%s"
	getSycResultURL     = "https://qyapi.weixin.qq.com/cgi-bin/batch/getresult?access_token=%s&jobid=%s"
)

// Callback 回调信息
type Callback struct {
	URL            string `json:"url"`
	Token          string `json:"token"`
	EncodingAesKey string `json:"encodingaeskey"`
}

// SyncUserParams 增量更新和全量覆盖成员参数
type SyncUserParams struct {
	MediaID  string   `json:"media_id"`
	ToInvite bool     `json:"to_invite"`
	Callback Callback `json:"callback"`
}

// SyncResp 异步返回
type SyncResp struct {
	util.CommonError
	JobID string `json:"jobid"`
}

// BatchSyncUser 增量更新成员
func (base *Base) BatchSyncUser(params SyncUserParams) (jobID string, err error) {
	var accessToken string
	accessToken, err = base.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(batchSyncUserURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, params)
	if err != nil {
		return
	}
	respData := new(SyncResp)
	err = json.Unmarshal(response, respData)
	if err != nil {
		return
	}
	if respData.ErrCode != 0 {
		err = fmt.Errorf("BatchSyncUser Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}

	jobID = respData.JobID
	return
}

// BatchReplaceUser 全量覆盖成员
func (base *Base) BatchReplaceUser(params SyncUserParams) (jobID string, err error) {
	var accessToken string
	accessToken, err = base.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(batchReplaceUserURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, params)
	if err != nil {
		return
	}
	respData := new(SyncResp)
	err = json.Unmarshal(response, respData)
	if err != nil {
		return
	}
	if respData.ErrCode != 0 {
		err = fmt.Errorf("BatchReplaceUser Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}

	jobID = respData.JobID
	return
}

// SyncDeptParams 全量覆盖部门参数
type SyncDeptParams struct {
	MediaID  string   `json:"media_id"`
	Callback Callback `json:"callback"`
}

// BatchReplaceParty 全量覆盖部门
func (base *Base) BatchReplaceParty(params SyncDeptParams) (jobID string, err error) {
	var accessToken string
	accessToken, err = base.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(batchReplaceDeptURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, params)
	if err != nil {
		return
	}
	respData := new(SyncResp)
	err = json.Unmarshal(response, respData)
	if err != nil {
		return
	}
	if respData.ErrCode != 0 {
		err = fmt.Errorf("BatchReplaceParty Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}

	jobID = respData.JobID
	return
}

// SyncReultItem 异步结果返回接口
type SyncReultItem interface {
}

// SyncUserResultItem 用户数据同步结果
type SyncUserResultItem struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	UserID  string `json:"userid"`
}

// SyncPartyResultItem 部门数据同步结果
type SyncPartyResultItem struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	Action  int    `json:"action"`
	PartyID int    `json:"partyid"`
}

// GetSyncResultResp 获取异步任务结果返回
type GetSyncResultResp struct {
	util.CommonError
	Status     int             `json:"status"`
	Type       string          `json:"type"`
	Total      int             `json:"total"`
	PercentAge int             `json:"percentage"`
	Result     []SyncReultItem `json:"result"`
}

// GetSyncResult 获取异步任务结果
func (base *Base) GetSyncResult(jobID string) (respData *GetSyncResultResp, err error) {
	var accessToken string
	accessToken, err = base.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getSycResultURL, accessToken, jobID)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	respData = new(GetSyncResultResp)
	err = json.Unmarshal(response, respData)
	if err != nil {
		return
	}
	if respData.ErrCode != 0 {
		err = fmt.Errorf("GetSyncResult Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}
	return
}
