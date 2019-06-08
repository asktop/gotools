package unique

import (
	"encoding/hex"
	"math/rand"
)

//随机md5字符串 32位
func RandMd5() string {
	data := make([]byte, 16)
	rand.Read(data)
	return hex.EncodeToString(data)
}
