package controllers

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/zenazn/goji/web"
)

func (f *Factory) MeasureClick() func(web.C, http.ResponseWriter, *http.Request) {
	return func(c web.C, w http.ResponseWriter, r *http.Request) {

		timeGen := r.FormValue("g")
		adID := r.FormValue("a")
		destLink := r.FormValue("l")
		tag := r.FormValue("t")
		originIP, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			fmt.Println("IP error")
			originIP = ""
		}
		referer := r.Header.Get("referer")

		timeGenUnix, err := strconv.ParseInt(timeGen, 10, 64)
		if err != nil {
			fmt.Println("Parse ERROR: ", err)
		} else {
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

			f.Data.ClickThroughs.Insert("uid", referer, destLink, originIP, timeGenTime, tag)

			//TODO redirect 301 to u
			//TODO add conversion cookie
			// http.Redirect(w, r, destLink, http.StatusFound)

		}
	}
}
