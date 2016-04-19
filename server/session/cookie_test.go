package session

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetCookie(t *testing.T) {

	r, _ := http.NewRequest("GET", "/somePATH", nil)
	w := httptest.NewRecorder()

	SetSessionCookie(w, r)

	ck := w.HeaderMap["Set-Cookie"]

	if len(ck) < 1 {
		t.Error("Cookie does not exist")
	}
}

func TestGetCookie(t *testing.T) {

	r, _ := http.NewRequest("GET", "/somePATH", nil)
	sID := MakeIDSession()
	cookie := MakeSessionCookie(sID)
	r.AddCookie(cookie)

	ck, err := GetSessionCookie(r)

	if len(ck) < 1 {
		t.Error("Cookie does not exist")
	}
	if ck != sID {
		t.Errorf("Cookie value expected %s , got %s", sID, ck)
	}
	if err != nil {
		t.Errorf("Error setting cookie: %s", err)
	}
}
