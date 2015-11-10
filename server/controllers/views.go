package controllers

import (
	"encoding/base64"
	"io"
	"log"
	"net"
	"net/http"

	"github.com/simonchong/linny/common"
	"github.com/simonchong/linny/server/wrappers"
	"github.com/zenazn/goji/web"
)

func ViewCounter(ac *wrappers.AppContext, sID string, c web.C, w http.ResponseWriter, r *http.Request) (int, error) {

	adID := r.FormValue("a")
	originIP, _, errIP := net.SplitHostPort(r.RemoteAddr)
	if errIP != nil {
		return http.StatusInternalServerError, errIP
	}
	timeGen, errT := common.FormTime("g", r)
	if errT != nil {
		return http.StatusInternalServerError, errT
	}
	referer := r.Header.Get("referer")

	log.Println("View ADID", adID)
	log.Println("View Origin IP", originIP)
	log.Println("View Gen Time", timeGen)
	log.Println("View Referer", referer)
	log.Println("View SessionID", sID)

	//Add to DB
	ac.Data.AdViews.Insert(adID, referer, originIP, timeGen, sID)

	//Send GIF response
	gif, _ := base64.StdEncoding.DecodeString("R0lGODlhAQABAIABAP///wAAACwAAAAAAQABAAACAkQBADs=")
	w.Header().Set("Content-Type", "image/gif")
	io.WriteString(w, string(gif))

	return 200, nil
}
