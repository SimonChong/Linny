package session

import (
	"net/http"
	"strings"
	"time"

	"github.com/satori/go.uuid"
)

const SessionCookieName = "lS"

func MakeSessionID() string {
	return strings.Replace(uuid.NewV1().String(), "-", "", -1)
}

func MakeSessionCookie(sessionID string) *http.Cookie {
	return &http.Cookie{
		Name:    SessionCookieName,
		Value:   sessionID,
		Path:    "/",
		Expires: time.Now().Add(time.Hour * 24 * 750),
	}
}

func SetSessionCookie(w http.ResponseWriter, r *http.Request) string {
	sessionID, err := GetSessionCookie(r)
	if err != nil {
		sessionID = MakeSessionID()
	}
	http.SetCookie(w, MakeSessionCookie(sessionID))
	return sessionID
}

func GetSessionCookie(r *http.Request) (string, error) {
	ckExisting, err := r.Cookie(SessionCookieName)
	if err == nil {
		return ckExisting.Value, nil
	}
	return "", err
}
