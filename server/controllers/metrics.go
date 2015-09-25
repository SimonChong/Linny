package controllers

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/simonchong/linny/server/wrappers"
	"github.com/zenazn/goji/web"
)

var isNum = regexp.MustCompile(`/\d/`)

func MeasureClick(ac *wrappers.AppContext, sID string, c web.C, w http.ResponseWriter, r *http.Request) (int, error) {

	adID := r.FormValue("a")
	originIP, _, errIP := net.SplitHostPort(r.RemoteAddr)
	if errIP != nil {
		return http.StatusInternalServerError, errIP
	}
	timeGen := r.FormValue("g")
	if !isNum.MatchString(timeGen) {
		return 404, errors.New("MeasureClick: timeGen is not a number")
	}
	timeGenUnix, errT := strconv.ParseInt(timeGen, 10, 64)
	if errT != nil {
		return http.StatusInternalServerError, errT
	}
	destLink := r.FormValue("l")
	tag := r.FormValue("t")
	referer := r.Header.Get("referer")

	now := time.Now().Unix()
	if timeGenUnix > now {
		timeGenUnix = now
	}
	timeGenTime := time.Unix(timeGenUnix, 0)

	fmt.Println("Link Path: ", r.URL.Path[1:])
	fmt.Println("Link Gen Time: ", timeGenTime)
	fmt.Println("Link ADID: ", adID)
	fmt.Println("Link Click Through: ", destLink)
	fmt.Println("Link Tag: ", tag)
	fmt.Println("Link Referer: ", referer)
	fmt.Println("Link SessionID", sID)

	ac.Data.AdClickThroughs.Insert(adID, referer, destLink, originIP,
		timeGenTime, tag, sID)

	//TODO redirect 301 to u
	//TODO add conversion cookie
	// http.Redirect(w, r, destLink, http.StatusFound)

	return 301, nil
}
