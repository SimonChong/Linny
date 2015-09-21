package controllers

import (
	"encoding/base64"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/zenazn/goji/web"
)

func (f *Factory) ViewCounter() func(web.C, http.ResponseWriter, *http.Request) {
	return func(c web.C, w http.ResponseWriter, r *http.Request) {

		adID := r.FormValue("a")
		originIP, _, err := net.SplitHostPort(r.RemoteAddr)
		timeGen := r.FormValue("g")
		referer := r.Header.Get("referer")

		fmt.Println("View ADID", adID)
		fmt.Println("View Origin IP", originIP)
		fmt.Println("View Gen Time", timeGen)
		fmt.Println("View Referer", referer)

		//Add to DB
		timeGenUnix, err := strconv.ParseInt(timeGen, 10, 64)
		if err != nil {
			timeGenTime := time.Unix(timeGenUnix, 0)

			f.Data.AdViews.Insert(adID, referer, originIP, timeGenTime, "TODO SESSION")
		}

		//Send GIF response
		gif, _ := base64.StdEncoding.DecodeString("R0lGODlhAQABAIABAP///wAAACwAAAAAAQABAAACAkQBADs=")
		w.Header().Set("Content-Type", "image/gif")
		io.WriteString(w, string(gif))
	}
}
