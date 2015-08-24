package controllers

import (
	"fmt"
	"net/http"

	"github.com/simonchong/linny/common"
)

func (f *Factory) AdHtml() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Requested: ", r.URL.Path[1:])
		// fmt.Println(r.Proto)
		// fmt.Println(r.Host)
		// fmt.Println(r.URL)

		contentRequested := r.URL.Path[1:]

		content, err := common.GetWrappedContent(f.Conf.ContentRoot, contentRequested)
		content = common.InjectLinks(content, r)

		if err != nil {
			fmt.Println("Content Controller Error: ", err)
			http.NotFound(w, r)
			return
		}

		w.Header().Set(
			"Content-Type",
			"text/html",
		)
		fmt.Fprint(w, content)
	}
}
