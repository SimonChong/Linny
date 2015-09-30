package wrappers

import (
	"log"
	"net/http"

	"github.com/zenazn/goji/web"
)

type AppContextHandler struct {
	AppContext *AppContext
	Handler    func(*AppContext, web.C, http.ResponseWriter, *http.Request) (int, error)
}

func (handle AppContextHandler) ServeHTTPC(c web.C, w http.ResponseWriter, r *http.Request) {

	//Handle request
	status, err := handle.Handler(handle.AppContext, c, w, r)
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
