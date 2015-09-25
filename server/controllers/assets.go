package controllers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/simonchong/linny/common"
	"github.com/simonchong/linny/constants"
	"github.com/simonchong/linny/server/wrappers"

	"github.com/zenazn/goji/web"
)

func AssetHTML(ac *wrappers.AppContext, sID string, c web.C, w http.ResponseWriter, r *http.Request) (int, error) {

	fmt.Println("Requested: ", r.URL.Path[1:])
	// fmt.Println(r.Proto)
	// fmt.Println(r.Host)
	// fmt.Println(r.URL)

	fileReq := c.URLParams["file"]

	fileAbs, err := common.ResolveSecure(ac.ConfLinny.ContentRoot+"/"+constants.AssetsDir, fileReq)
	if err != nil {
		fmt.Println("Secure Resolve Failed: ", err)
		return 404, errors.New("Secure Resolve Failed: " + err.Error())
	}
	exists, fileAbs := common.FileExistsHTML(fileAbs)
	if !exists {
		fmt.Println("File Does Not Exist: ", fileAbs)
		return 404, errors.New("File Does Not Exist " + fileAbs)
	}

	//Metrics collection
	originIP, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Println("IP error")
		originIP = ""
	}
	referer := r.Header.Get("referer")
	ac.Data.AdDownloads.Insert(ac.ConfAd.Id, r.URL.Path[1:], referer, originIP, "SESSIONID TODO")

	//Content Processing and rendering
	content, err := getWrappedContent(fileAbs, ac.ConfLinny.ContentRoot, ac.ConfAd.Id, r)
	if err != nil {
		fmt.Println("Content Error: ", err)
		http.NotFound(w, r)
		return 502, errors.New("Content Error: " + err.Error())
	}
	content = common.InjectLinks(ac.ConfAd.Id, content, r)

	w.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprint(w, content)
	return 200, nil
}

func AssetFiles(ac *wrappers.AppContext, sID string, c web.C, w http.ResponseWriter, r *http.Request) (int, error) {

	absBaseDir, _ := filepath.Abs(ac.ConfLinny.ContentRoot)
	fileServeDir := absBaseDir + "/" + constants.AssetsDir
	fmt.Println(fileServeDir)
	handle := http.StripPrefix("/"+constants.AssetsDir+"/", http.FileServer(http.Dir(fileServeDir)))

	handle.ServeHTTP(w, r)
	return 200, nil
}

func getTrackingCode(adID string, host string) string {
	code, err := ioutil.ReadFile("./resources/tracking.js")
	if err != nil {
		panic(err)
	}
	unix := strconv.FormatInt(time.Now().Unix(), 10)

	return "<script defer='defer'>(function(a, h, v, g) {" + string(code) + "})('" + adID + "' , '" + host + "','" + constants.ViewsDir + "'," + unix + ");</script>"
}

func getWrappedContent(path string, root string, adID string, r *http.Request) (string, error) {

	content, err0 := ioutil.ReadFile(path)
	if err0 != nil {
		return "", err0
	}
	header, err1 := ioutil.ReadFile(root + "/header.frag")
	if err1 != nil {
		return "", err1
	}
	footer, err2 := ioutil.ReadFile(root + "/footer.frag")
	if err2 != nil {
		return "", err2
	}

	rtn := string(header) + string(content) + getTrackingCode(adID, r.Host) + string(footer)

	return rtn, nil
}
