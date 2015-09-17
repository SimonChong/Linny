package controllers

import (
	"fmt"
	"net"
	"net/http"
	"path/filepath"

	"github.com/simonchong/linny/common"
	"github.com/simonchong/linny/constants"

	"github.com/zenazn/goji/web"
)

func (f *Factory) AssetHTML() func(web.C, http.ResponseWriter, *http.Request) {
	return func(c web.C, w http.ResponseWriter, r *http.Request) {

		fmt.Println("Requested: ", r.URL.Path[1:])
		// fmt.Println(r.Proto)
		// fmt.Println(r.Host)
		// fmt.Println(r.URL)

		fileReq := c.URLParams["file"]

		fileAbs, err := common.ResolveSecure(f.ConfLinny.ContentRoot+"/"+constants.AssetsDir, fileReq)
		if err != nil {
			fmt.Println("Secure Resolve Failed: ", err)
			http.NotFound(w, r)
			return
		}
		exists, fileAbs := common.FileExistsHTML(fileAbs)
		if !exists {
			fmt.Println("File Does Not Exist: ", fileAbs)
			http.NotFound(w, r)
			return
		}

		//Metrics collection
		originIP, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			fmt.Println("IP error")
			originIP = ""
		}
		referer := r.Header.Get("referer")
		f.Data.AdDownloads.Insert(f.ConfAd.Id, r.URL.Path[1:], referer, originIP, "SESSIONID TODO")

		//Content Processing and rendering
		content, err := common.GetWrappedContent(fileAbs, f.ConfLinny.ContentRoot)
		if err != nil {
			fmt.Println("Content Error: ", err)
			http.NotFound(w, r)
			return
		}
		content = common.InjectLinks(f.ConfAd.Id, content, r)

		w.Header().Set(
			"Content-Type",
			"text/html",
		)
		fmt.Fprint(w, content)
	}
}

func (f *Factory) AssetFiles() http.Handler {

	absBaseDir, _ := filepath.Abs(f.ConfLinny.ContentRoot)
	fileServeDir := absBaseDir + "/" + constants.AssetsDir
	fmt.Println(fileServeDir)
	return http.StripPrefix("/"+constants.AssetsDir+"/", http.FileServer(http.Dir(fileServeDir)))
}
