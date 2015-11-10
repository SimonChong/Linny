package wrappers

import (
	"log"
	"net/http"

	"github.com/simonchong/linny/server/session"
	"github.com/zenazn/goji/web"
)

type AppSessionHandler struct {
	AppContext *AppContext
	Handler    func(*AppContext, string, web.C, http.ResponseWriter, *http.Request) (int, error)
}

func (handle AppSessionHandler) ServeHTTPC(c web.C, w http.ResponseWriter, r *http.Request) {

	//Generate sessionID
	sessionID := session.SetSessionCookie(w, r)

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
