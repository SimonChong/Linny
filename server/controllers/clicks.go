package controllers

import (
	"fmt"
	"net"
	"net/http"

	"github.com/simonchong/linny/common"
	"github.com/simonchong/linny/server/controllers/conversions"
	"github.com/simonchong/linny/server/wrappers"
	"github.com/zenazn/goji/web"
)

func ClickTracking(ac *wrappers.AppContext, sID string, c web.C, w http.ResponseWriter, r *http.Request) (int, error) {

	adID := r.FormValue("a")
	originIP, _, errIP := net.SplitHostPort(r.RemoteAddr)
	if errIP != nil {
		return http.StatusInternalServerError, errIP
	}

	timeGen, errT := common.FormTime("g", r)
	if errT != nil {
		return http.StatusInternalServerError, errT
	}

	destLink := r.FormValue("l")
	tag := r.FormValue("t")
	referer := r.Header.Get("referer")

	conversions.AddCookie(w, r, adID)

	fmt.Println("Link Path: ", r.URL.Path[1:])
	fmt.Println("Link Gen Time: ", timeGen)
	fmt.Println("Link ADID: ", adID)
	fmt.Println("Link Click Through: ", destLink)
	fmt.Println("Link Tag: ", tag)
	fmt.Println("Link Referer: ", referer)
	fmt.Println("Link SessionID", sID)

	ac.Data.AdClickThroughs.Insert(adID, referer, destLink, originIP,
		timeGen, tag, sID)

	//TODO redirect 301 to u
	// http.Redirect(w, r, destLink, http.StatusFound)

	return 301, nil
}
