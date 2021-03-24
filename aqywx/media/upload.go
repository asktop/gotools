package media

import (
	"encoding/json"
	"fmt"

	"github.com/asktop/gotools/aqywx/util"
)

const (
	uploadTmpURL   = "https://qyapi.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s"
	uploadImageURL = "https://qyapi.weixin.qq.com/cgi-bin/media/uploadimg?access_token=%s"
)

// Type 媒体文件类型
type Type string

const (
	// MediaTypeImage 媒体文件:图片
	MediaTypeImage Type = "image"
	// MediaTypeVoice 媒体文件:声音
	MediaTypeVoice = "voice"
	// MediaTypeVideo 媒体文件:视频
	MediaTypeVideo = "video"
	// MediaTypeFile 媒体文件:普通文件
	MediaTypeFile = "file"
)

// UploadTmpResp 临时素材上传返回信息
type UploadTmpResp struct {
	util.CommonError
	Type      Type   `json:"type"`
	MediaID   string `json:"media_id"`
	CreatedAt int64  `json:"created_at"`
}

// UploadTmp 上传临时素材
func (media *Media) UploadTmp(fileName string, mediaType Type) (respData *UploadTmpResp, err error) {
	var accessToken string
	accessToken, err = media.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(uploadTmpURL, accessToken, mediaType)
	var response []byte
	response, err = util.PostFile("media", fileName, uri)
	if err != nil {
		return
	}
	respData = new(UploadTmpResp)
	err = json.Unmarshal(response, respData)
	if err != nil {
		return
	}
	if respData.ErrCode != 0 {
		err = fmt.Errorf("UploadTmp Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}
	return
}

// UploadImageResp 上传图片返回
type UploadImageResp struct {
	util.CommonError
	URL string `json:"url"`
}

// UploadImage 上传永久图片
func (media *Media) UploadImage(fileName string) (urlStr string, err error) {
	var accessToken string
	accessToken, err = media.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(uploadImageURL, accessToken)
	var response []byte
	response, err = util.PostFile("media", fileName, uri)
	if err != nil {
		return
	}
	respData := new(UploadImageResp)
	err = json.Unmarshal(response, respData)
	if err != nil {
		return
	}
	if respData.ErrCode != 0 {
		err = fmt.Errorf("UploadTmp Error , errcode=%d , errmsg=%s", respData.ErrCode, respData.ErrMsg)
		return
	}

	urlStr = respData.URL
	return
}
