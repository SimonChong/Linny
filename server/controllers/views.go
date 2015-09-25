package controllers

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/simonchong/linny/server/wrappers"
	"github.com/zenazn/goji/web"
)

func ViewCounter(ac *wrappers.AppContext, sID string, c web.C, w http.ResponseWriter, r *http.Request) (int, error) {

	adID := r.FormValue("a")
	originIP, _, errIP := net.SplitHostPort(r.RemoteAddr)
	if errIP != nil {
		return http.StatusInternalServerError, errIP
	}
	timeGen := r.FormValue("g")
	fmt.Println(timeGen)
	if !isNum.MatchString(timeGen) {
		return http.StatusServiceUnavailable, errors.New("ViewCounter: timeGen is not a number")
	}

	timeGenUnix, errT := strconv.ParseInt(timeGen, 10, 64)
	if errT != nil {
		return http.StatusInternalServerError, errT
	}
	referer := r.Header.Get("referer")

	//Add to DB
	timeGenTime := time.Unix(timeGenUnix, 0)
	fmt.Println("View ADID", adID)
	fmt.Println("View Origin IP", originIP)
	fmt.Println("View Gen Time", timeGen)
	fmt.Println("View Referer", referer)
	fmt.Println("View SessionID", sID)

	ac.Data.AdViews.Insert(adID, referer, originIP, timeGenTime, sID)

	//Send GIF response
	gif, _ := base64.StdEncoding.DecodeString("R0lGODlhAQABAIABAP///wAAACwAAAAAAQABAAACAkQBADs=")
	w.Header().Set("Content-Type", "image/gif")
	io.WriteString(w, string(gif))

	return 200, nil
}
