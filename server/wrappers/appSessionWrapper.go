package wrappers

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/satori/go.uuid"
	"github.com/zenazn/goji/web"
)

const SessionCookieName = "lS"

type AppSessionHandler struct {
	AppContext *AppContext
	Handler    func(*AppContext, string, web.C, http.ResponseWriter, *http.Request) (int, error)
}

func (handle AppSessionHandler) ServeHTTPC(c web.C, w http.ResponseWriter, r *http.Request) {

	//Generate sessionID
	sessionID := ""
	ckExisting, err := r.Cookie(SessionCookieName)
	if err != nil {
		sessionID = ckExisting.Value
	} else {
		sessionID = strings.Replace(uuid.NewV1().String(), "-", "", -1)
	}
	ck := &http.Cookie{
		Name:    SessionCookieName,
		Value:   sessionID,
		Path:    "/",
		Expires: time.Now().Add(time.Hour * 24 * 750),
	}
	http.SetCookie(w, ck)

	//Handle request
	status, err := handle.Handler(handle.AppContext, sessionID, c, w, r)
	if err != nil {
		log.Printf("HTTP %d: %q", status, err)
		switch status {
		case http.StatusNotFound:
			http.NotFound(w, r)
		case http.StatusInternalServerError:
			http.Error(w, http.StatusText(status), status)
		default:
			http.Error(w, http.StatusText(status), status)
		}
	}
}
