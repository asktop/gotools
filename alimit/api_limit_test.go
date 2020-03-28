package alimit

import "testing"

func TestNewAPILimit(t *testing.T) {
	apiLimit := NewAPILimit(300)
	apiUniqueKey := "Request.Method" + "Request.URL.Path" + "Input.IP or UserId or Token"
	if !apiLimit.Check(apiUniqueKey) {
		//频率受限，返回
		return
	}
	//进行其他业务操作
}
