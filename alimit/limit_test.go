package alimit

import "testing"

func TestApiRateLimit(t *testing.T) {
	apiLimit := NewAPIRateLimit(300)
	apiUniqueKey := "Request.Method" + "Request.URL" + "IP or UserId"
	if !apiLimit.Check(apiUniqueKey) {
		//频率受限，返回
		return
	}
	//进行其他业务操作
}
