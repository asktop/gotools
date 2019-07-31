package arand

import (
	"encoding/base32"
	"encoding/hex"
	"github.com/asktop/gotools/atime"
	"github.com/asktop/gotools/auuid"
	"math/rand"
	"strings"
)

func Rand() *rand.Rand {
	return rand.New(rand.NewSource(atime.NowNano() + rand.Int63n(1e10)))
}

func Intn(n int) int {
	random := Rand()
	return random.Intn(n)
}

func Int63n(n int64) int64 {
	random := Rand()
	return random.Int63n(n)
}

//生成在 min 和 max 之间的随机数（包含 min 和 max）
func RandInt(min int, max int) int {
	if min >= max {
		return min
	}
	return Intn(max-min+1) + min
}

//生成在 min 和 max 之间的随机数（包含 min 和 max）
func RandInt64(min int64, max int64) int64 {
	if min >= max {
		return min
	}
	return Int63n(max-min+1) + min
}

//随机字符串 指定长度
// @param sources 数据源 默认（0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz）
// 	空:数字+大小写字母
// 	0:只数字
//	a:只小写字母
//	A:只大写字母
//	Aa:大小写字母
//	_:数字+大小写字母+下划线
//	其他:自定义
func RandStr(length int, sources ...string) string {
	rs := make([]rune, length)
	var source []rune
	if len(sources) > 0 {
		typ := sources[0]
		if typ == "0" {
			source = []rune("0123456789")
		} else if typ == "A" {
			source = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
		} else if typ == "a" {
			source = []rune("abcdefghijklmnopqrstuvwxyz")
		} else if typ == "0A" || typ == "A0" {
			source = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
		} else if typ == "0a" || typ == "a0" {
			source = []rune("0123456789abcdefghijklmnopqrstuvwxyz")
		} else if typ == "Aa" || typ == "aA" {
			source = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
		} else if typ == "0Aa_" || typ == "_" {
			source = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_")
		} else {
			source = []rune(strings.Join(sources, ""))
		}
	} else {
		source = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	}
	sourceLen := len(source)
	if sourceLen == 0 {
		return ""
	}

	for i := range rs {
		rs[i] = source[Intn(sourceLen)]
	}

	return string(rs)
}

//随机UUID字符串，32位 dba4756aa25f49fc969f30b3ef234d68
func RandUUID() string {
	return strings.ReplaceAll(auuid.New().String(), "-", "")
}

//随机md5字符串，size默认16
// size = 16 返回32位 52fdfc072182654f163f5f0f9a621d72
// size = 10 返回20位 52fdfc072182654f163f
func RandMd5(size ...int) string {
	siz := 16
	if len(size) > 0 {
		siz = size[0]
	}
	data := make([]byte, siz)
	Rand().Read(data)
	return hex.EncodeToString(data)
}

//随机base32字符串，size默认16
// size = 16 返回32位 SVTMOTIQAN6E2653AQD5DYWGJE======
// size = 10 返回16位 KL67YBZBQJSU6FR7
func RandBase32(size ...int) string {
	siz := 16
	if len(size) > 0 {
		siz = size[0]
	}
	data := make([]byte, siz)
	Rand().Read(data)
	return base32.StdEncoding.EncodeToString(data)
}
