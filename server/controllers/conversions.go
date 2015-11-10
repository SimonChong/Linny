package controllers

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"

	"time"

	"github.com/simonchong/linny/common"
	"github.com/simonchong/linny/server/controllers/conversions"
	"github.com/simonchong/linny/server/controllers/resources"
	"github.com/simonchong/linny/server/paths"
	"github.com/simonchong/linny/server/session"
	"github.com/simonchong/linny/server/wrappers"
	"github.com/zenazn/goji/web"
)

//go:generate wgf -i=../../resources/conversion.js -o=./resources/conversionJS.go -p=resources -c=ConversionJS

func ConversionsJS(ac *wrappers.AppContext, c web.C, w http.ResponseWriter, r *http.Request) (int, error) {

	unix := strconv.FormatInt(time.Now().Unix(), 10)
	tag := r.FormValue("t")

	body := "(function(h, v, t, g) {" + string(resources.ConversionJS) + "})('" + r.Host + "','" + paths.MeasureDir + "','" + tag + "'," + unix + ");"

	w.Header().Set(
		"Content-Type",
		"application/javascript",
	)
	fmt.Fprint(w, body)

	return 200, nil
}

func Conversions(ac *wrappers.AppContext, c web.C, w http.ResponseWriter, r *http.Request) (int, error) {

	adID, errA := conversions.GetCookie(r)
	if errA != nil {
		return http.StatusOK, errA
	}

	originIP, _, errIP := net.SplitHostPort(r.RemoteAddr)
	if errIP != nil {
		return http.StatusOK, errIP
	}

	timeGen, errT := common.FormTime("g", r)
	if errT != nil {
		return http.StatusOK, errT
	}

	sessionID, errS := session.GetSessionCookie(r)
	if errS != nil {
		return http.StatusOK, errS
	}

	conversionTag := r.FormValue("t")
	//TODO limit to 255 characters

	referer := r.Header.Get("referer")

	log.Println("Conversion Origin IP", originIP)
	log.Println("Conversion Referer", referer)
	log.Println("Conversion Gen Time", timeGen)
	log.Println("Conversion adID", adID)
	log.Println("Conversion Conversion Tag", conversionTag)
	log.Println("Conversion Session", sessionID)

	//Add to DB
	ac.Data.AdConversions.Insert(adID, referer, originIP, conversionTag, sessionID)

	//Send GIF response
	gif, _ := base64.StdEncoding.DecodeString("R0lGODlhAQABAIABAP///wAAACwAAAAAAQABAAACAkQBADs=")
	w.Header().Set("Content-Type", "image/gif")
	io.WriteString(w, string(gif))

	return 200, nil
}
