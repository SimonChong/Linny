package controllers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/simonchong/linny/common"
	"github.com/simonchong/linny/server/controllers/assets"
	"github.com/simonchong/linny/server/controllers/resources"
	"github.com/simonchong/linny/server/paths"
	"github.com/simonchong/linny/server/wrappers"

	"github.com/zenazn/goji/web"
)

//go:generate wgf -i=../../resources/tracking.js -o=./resources/trackingJS.go -p=resources -c=TrackingJS

func AssetHTML(ac *wrappers.AppContext, sID string, c web.C, w http.ResponseWriter, r *http.Request) (int, error) {

	// log.Println("Requested: ", r.URL.Path[1:])
	// log.Println(r.Proto)
	// log.Println(r.Host)
	// log.Println(r.URL)

	fileReq := c.URLParams["file"]

	fileAbs, err := common.ResolveSecure(ac.ConfLinny.ContentRoot+"/"+paths.AssetsDir, fileReq)
	if err != nil {
		log.Println("Secure Resolve Failed: ", err)
		return 404, errors.New("Secure Resolve Failed: " + err.Error())
	}
	exists, fileAbs := common.FileExistsHTML(fileAbs)
	if !exists {
		log.Println("File Does Not Exist: ", fileAbs)
		return 404, errors.New("File Does Not Exist " + fileAbs)
	}

	//Metrics collection
	originIP, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Println("IP error")
		originIP = ""
	}
	referer := r.Header.Get("referer")
	ac.Data.AdDownloads.Insert(ac.ConfAd.Id, r.URL.Path[1:], referer, originIP, sID)

	//Content Processing and rendering
	content, err := getWrappedContent(fileAbs, ac.ConfLinny.ContentRoot, ac.ConfAd.Id, r)
	if err != nil {
		log.Println("Content Error: ", err)
		http.NotFound(w, r)
		return 502, errors.New("Content Error: " + err.Error())
	}
	content = assets.InjectLinks(ac.ConfAd.Id, content, r)

	w.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprint(w, content)
	return 200, nil
}

func AssetFiles(ac *wrappers.AppContext, sID string, c web.C, w http.ResponseWriter, r *http.Request) (int, error) {

	absBaseDir, _ := filepath.Abs(ac.ConfLinny.ContentRoot)
	fileServeDir := absBaseDir + "/" + paths.AssetsDir
	log.Println(fileServeDir)
	handle := http.FileServer(http.Dir(fileServeDir))

	handle.ServeHTTP(w, r)
	return 200, nil
}

func getTrackingCode(adID string, host string) string {
	unix := strconv.FormatInt(time.Now().Unix(), 10)

	return "<script defer='defer'>(function(a, h, v, g) {" + resources.TrackingJS + "})('" + adID + "' , '" + host + "','" + paths.MeasureDir + "'," + unix + ");</script>"
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
