package astring

import "strings"

//拼接URL
func JoinURL(url ...string) string {
	var rs []string
	for _, u := range url {
		u = strings.TrimSpace(u)
		u = strings.TrimPrefix(u, "/")
		u = strings.TrimPrefix(u, `\`)
		u = strings.TrimSuffix(u, "/")
		u = strings.TrimSuffix(u, `\`)
		rs = append(rs, u)
	}
	return strings.Join(rs, `/`)
}
