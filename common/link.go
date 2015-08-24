package common

import (
	// "fmt"
	"net/http"
	"net/url"
	"regexp"
)

var regLK = regexp.MustCompile("{{lk\\s+[\"']?([^{}\"']+)[\"']?\\s*}}")
var regTLK = regexp.MustCompile("{{tlk\\s+[\"']?([^{}\"']+)[\"']?\\s*}}")

func InjectLinks(content string, r *http.Request) string {

	// fmt.Println(r.Host)
	// fmt.Println(r.URL)
	baseURL := "//" + r.Host + "/"
	content = replaceLK(content, baseURL)
	content = replaceTLK(content, baseURL)

	return content
}

func replaceLK(content string, baseURL string) string {
	return regLK.ReplaceAllString(content, baseURL+"${1}")
}

func replaceTLK(content string, baseURL string) string {
	return regTLK.ReplaceAllStringFunc(content, func(src string) string {
		// fmt.Println(regTLK.FindStringSubmatch(src)[1])
		return baseURL + "track/u=" + url.QueryEscape(regTLK.FindStringSubmatch(src)[1])
	})
}
