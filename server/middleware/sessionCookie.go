package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/satori/go.uuid"
)

const SessionCookieName = "lS"

// Session Cookie Generator
func SessionCookieGen(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		lSessionID := ""
		ckExisting, err := r.Cookie(SessionCookieName)
		if err != nil {
			lSessionID = ckExisting.Value
		} else {
			lSessionID = strings.Replace(uuid.NewV1().String(), "-", "", -1)
		}
		ck := &http.Cookie{
			Name:    SessionCookieName,
			Value:   lSessionID,
			Path:    "/",
			Expires: time.Now().Add(time.Hour * 24 * 750),
		}
		http.SetCookie(w, ck)

		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
