package conversions

import (
	"net/http"
	"time"
)

const conversionCookie = "lC"

func AddCookie(w http.ResponseWriter, r *http.Request, adID string) {
	ck := &http.Cookie{
		Name:    conversionCookie,
		Value:   adID,
		Path:    "/",
		Expires: time.Now().Add(time.Hour * 24 * 7),
	}
	http.SetCookie(w, ck)
}

func GetCookie(r *http.Request) (string, error) {
	ckExisting, err := r.Cookie(conversionCookie)
	if err == nil {
		return ckExisting.Value, nil
	}
	return "", err
}
