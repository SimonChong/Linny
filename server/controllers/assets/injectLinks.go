package assets

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/simonchong/linny/server/paths"
)

var timeNow time.Time

func InjectLinks(adID string, content string, r *http.Request) string {

	// log.Println(r.Host)
	// log.Println(r.URL)
	timeNow = time.Now()

	host := r.Host

	//Get the last directory

	var curPath = ""
	prefixPth, err := regexp.MatchString("^[\\./]", r.URL.Path)
	if prefixPth && err == nil {
		curPath = path.Dir(r.URL.Path[1:])
	} else {
		curPath = path.Dir(r.URL.Path)
	}
	if curPath == "." || curPath == "/" {
		curPath = ""
	}

	// log.Println("HOST: " + host)
	// log.Println("URL: " + r.URL.String())
	// log.Println("PTH: " + curPath)

	content = replaceILK(content, host, curPath)
	content = replaceMLK(content, host, curPath, adID)

	return content
}

func resolveLink(host string, currentDir string, link string) string {
	currentDir = strings.TrimSuffix(currentDir, "/")
	currentDir = strings.TrimPrefix(currentDir, "/")
	if currentDir != "" {
		currentDir += "/"
	}

	if strings.HasPrefix(link, "./") {
		link = "//" + host + "/" + currentDir + strings.Replace(link, "./", "", 1)
	} else if strings.HasPrefix(link, "/") && !strings.HasPrefix(link, "//") {
		link = "//" + host + link
	} else if !strings.HasPrefix(link, "http://") && !strings.HasPrefix(link, "https://") {
		link = "//" + host + "/" + currentDir + link
	}
	// log.Println(host, currentDir, link)
	return link
}

var regILK = regexp.MustCompile("{{ilk\\s+[\"']?([^{}\"']+)[\"']?\\s*}}")

// Internal Links
func replaceILK(content string, host string, path string) string {
	return regILK.ReplaceAllStringFunc(content, func(src string) string {
		return resolveLink(host, path, regILK.FindStringSubmatch(src)[1])
	})
}

var regMLK = regexp.MustCompile("{{mlk\\s+[\"']?([^{}\"']+)[\"']?[^}]*}}")
var tagExtract = regexp.MustCompile("tag\\s*=\\s*[\"']([\\w ]+)[\"']")

func replaceMLK(content string, host string, path string, adID string) string {
	return regMLK.ReplaceAllStringFunc(content, func(src string) string {
		linkTo := resolveLink(host, path, regMLK.FindStringSubmatch(src)[1])
		tag := tagExtract.FindStringSubmatch(src)

		link := "//" + host + "/" + paths.MeasureDir + "/k?"
		link += "g=" + url.QueryEscape(fmt.Sprint(timeNow.Unix())) + "&"
		link += "a=" + url.QueryEscape(adID) + "&"
		link += "l=" + url.QueryEscape(linkTo)

		if len(tag) > 1 {
			link += "&t=" + url.QueryEscape(tag[1])
		}

		return link
	})
}
