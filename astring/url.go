package astring

import "strings"

//拼接URL
func JoinURL(uris ...string) string {
    var url string
    var urisTemp []string
    for _, uri := range uris {
        uri = strings.Trim(uri, " ")
        uri = strings.ReplaceAll(uri, `\`, "/")
        uri = strings.Trim(uri, "/")
        if uri != "" {
            urisTemp = append(urisTemp, uri)
        }
    }
    url = strings.Join(urisTemp, "/")
    if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
        url = "/" + url
    }
    return url
}
