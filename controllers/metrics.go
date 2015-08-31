package controllers

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji/web"
)

func (f *Factory) MetricsClick() func(web.C, http.ResponseWriter, *http.Request) {
	return func(c web.C, w http.ResponseWriter, r *http.Request) {

		fmt.Println("Metrics Click: ", r.URL.Path[1:])
		fmt.Println("Click Through: ", r.FormValue("u"))

		//TODO redirect 301 to u
		// redirectURL := r.FormValue("u")
		// http.Redirect(w, r, redirectURL, http.StatusFound)
	}
}
