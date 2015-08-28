package common

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strings"
	"time"
)

var regLK = regexp.MustCompile("{{lk\\s+[\"']?([^{}\"']+)[\"']?\\s*}}")
var regTLK = regexp.MustCompile("{{tlk\\s+[\"']?([^{}\"']+)[\"']?\\s*}}")
var timeNow time.Time

func InjectLinks(content string, r *http.Request) string {

	// fmt.Println(r.Host)
	// fmt.Println(r.URL)
	timeNow = time.Now()

	host := r.Host
	curPath := path.Dir(r.URL.Path[1:])

	content = replaceLK(content, host, curPath)
	content = replaceLKT(content, host, curPath)

	return content
}

func resolveLink(host string, currentDir string, link string) string {
	currentDir = strings.TrimSuffix(currentDir, "/")
	currentDir = strings.TrimPrefix(currentDir, "/")
	if strings.HasPrefix(link, "./") {
		link = "//" + host + "/" + currentDir + "/" + strings.Replace(link, "./", "", 1)
	} else if strings.HasPrefix(link, "/") && !strings.HasPrefix(link, "//") {
		link = "//" + host + link
	} else if !strings.HasPrefix(link, "http://") && !strings.HasPrefix(link, "https://") {
		link = "//" + host + "/" + currentDir + "/" + link
	}
	return link
}

func replaceLK(content string, host string, path string) string {
	return regLK.ReplaceAllStringFunc(content, func(src string) string {
		return resolveLink(host, path, regLK.FindStringSubmatch(src)[1])
	})
}

func replaceLKT(content string, host string, path string) string {
	return regTLK.ReplaceAllStringFunc(content, func(src string) string {
		linkTo := resolveLink(host, path, regTLK.FindStringSubmatch(src)[1])
		link := "//" + host + "/" + "track/click?"
		link += "u=" + url.QueryEscape(linkTo)
		link += "&t=" + fmt.Sprint(timeNow.Unix())
		return link
	})
}
