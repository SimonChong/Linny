package conversions

import (
	"net/http"
	"time"
)

const conversionCookie = "lC"

func NewCookie(adID string) *http.Cookie {
	return &http.Cookie{
		Name:    conversionCookie,
		Value:   adID,
		Path:    "/",
		Expires: time.Now().Add(time.Hour * 24 * 7),
	}
}

func AddCookie(w http.ResponseWriter, adID string) {
	http.SetCookie(w, NewCookie(adID))
}

func GetCookie(r *http.Request) (string, error) {
	ckExisting, err := r.Cookie(conversionCookie)
	if err == nil {
		return ckExisting.Value, nil
	}
	return "", err
}
