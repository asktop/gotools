package media

import (
	"fmt"
)

const (
	getTmpMediaURL = "https://qyapi.weixin.qq.com/cgi-bin/media/get?access_token=%s&media_id=%s"
	getVoiceURL    = "https://qyapi.weixin.qq.com/cgi-bin/media/get/jssdk?access_token=%s&media_id=%s"
)

// GetTmpMediaURL 获取临时素材下载地址
func (media *Media) GetTmpMediaURL(mediaID string) (urlStr string, err error) {
	var accessToken string
	accessToken, err = media.GetAccessToken()
	if err != nil {
		return
	}

	urlStr = fmt.Sprintf(getTmpMediaURL, accessToken, mediaID)
	return
}

// GetVoiceURL 获取高清语音下载地址
func (media *Media) GetVoiceURL(mediaID string) (urlStr string, err error) {
	var accessToken string
	accessToken, err = media.GetAccessToken()
	if err != nil {
		return
	}

	urlStr = fmt.Sprintf(getVoiceURL, accessToken, mediaID)
	return
}
