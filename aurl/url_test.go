package aurl

import (
	"net/url"
	"testing"
)

var urlStr string = `https://golang.org/x/crypto?go-get=1 +`
var urlEncode string = `https%3A%2F%2Fgolang.org%2Fx%2Fcrypto%3Fgo-get%3D1+%2B`
var rawUrlEncode string = `https%3A%2F%2Fgolang.org%2Fx%2Fcrypto%3Fgo-get%3D1%20%2B`

func TestEncodeAndDecode(t *testing.T) {
	t.Log(Encode(urlStr) == urlEncode)
	t.Log(Encode(urlStr))


	res, err := Decode(urlEncode)
	if err != nil {
		t.Errorf("decode failed. %v", err)
		return
	}
	t.Log(res == urlStr)
	t.Log(res)
}

func TestRowEncodeAndDecode(t *testing.T) {
	t.Log(RawEncode(urlStr) == rawUrlEncode)
	t.Log(RawEncode(urlStr))

	res, err := RawDecode(rawUrlEncode)
	if err != nil {
		t.Errorf("decode failed. %v", err)
		return
	}
	t.Log(res == urlStr)
	t.Log(res)
}

func TestBuildQuery(t *testing.T) {
	src := url.Values{
		"a": {"a2", "a1"},
		"b": {"b2", "b1"},
		"c": {"c1", "c2"},
	}
	expect := "a=a2&a=a1&b=b2&b=b1&c=c1&c=c2"

	t.Log(BuildQuery(src) == expect)
	t.Log(BuildQuery(src))
}

func TestParseURL(t *testing.T) {
	src := `http://username:password@hostname:9090/path?arg=value#anchor`
	expect := map[string]string{
		"scheme":   "http",
		"host":     "hostname",
		"port":     "9090",
		"user":     "username",
		"pass":     "password",
		"path":     "/path",
		"query":    "arg=value",
		"fragment": "anchor",
	}

	component := 0
	for k, v := range []string{"all", "scheme", "host", "port", "user", "pass", "path", "query", "fragment"} {
		if v == "all" {
			component = -1
		} else {
			component = 1 << (uint(k - 1))
		}

		res, err := ParseURL(src, component)
		if err != nil {
			t.Errorf("ParseURL failed. component:%v, err:%v", component, err)
			return
		}

		if v == "all" {
			t.Log(res)
		} else {
			t.Log(res[v] == expect[v])
			t.Log(res[v])
		}
	}
}
