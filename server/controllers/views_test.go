package controllers

import (
	"database/sql"
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/simonchong/linny/insights"
	"github.com/simonchong/linny/insights/datatest"
	"github.com/simonchong/linny/server/controllers/conversions"
	"github.com/simonchong/linny/server/session"
	"github.com/simonchong/linny/server/wrappers"
	"github.com/zenazn/goji/web"
)

type mockViewsTable struct {
	data map[string]string
}

func (m *mockViewsTable) Init() {}
func (m *mockViewsTable) Insert(adID string, refererURL string, originIP string, contentGeneratedOn time.Time, sessionID string) (sql.Result, error) {
	m.data["adID"] = adID
	m.data["refererURL"] = refererURL
	m.data["originIP"] = originIP
	m.data["contentGeneratedOn"] = contentGeneratedOn.String()
	m.data["sessionID"] = sessionID

	return &datatest.MockDBResult{1, 1}, nil
}

func TestViews(t *testing.T) {

	var mockTable = &mockViewsTable{map[string]string{}}
	var mockData = &insights.Data{AdViews: mockTable}
	var mockAppContext = &wrappers.AppContext{Data: mockData}
	var mockReq, _ = http.NewRequest("GET", "/m/k", nil)
	mockReq.RemoteAddr = "127.0.0.1:80"
	mockReq.Form = map[string][]string{
		"g": []string{"1447100000"},
		"a": []string{"ADID123"},
	}
	var sessionID = session.MakeIDSession()
	mockReq.AddCookie(conversions.NewCookie("ADID123"))
	mockReq.AddCookie(session.MakeSessionCookie(sessionID))
	mockReq.Header.Add("Referer", "http://test.test")

	var w = httptest.NewRecorder()

	var code, err = ViewCounter(mockAppContext, sessionID, web.C{}, w, mockReq)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if code != 200 {
		t.Errorf("Status: %d", code)
	}

	if mockTable.data["adID"] != "ADID123" {
		t.Errorf("Error - adID : %s", mockTable.data["adID"])
	}
	if mockTable.data["refererURL"] != "http://test.test" {
		t.Errorf("Error - refererURL : %s", mockTable.data["refererURL"])
	}
	if mockTable.data["originIP"] != "127.0.0.1" {
		t.Errorf("Error - originIP : %s", mockTable.data["originIP"])
	}
	if mockTable.data["contentGeneratedOn"] != "2015-11-10 07:13:20 +1100 AEDT" {
		t.Errorf("Error - contentGeneratedOn : %s", mockTable.data["contentGeneratedOn"])
	}
	if mockTable.data["sessionID"] != sessionID {
		t.Errorf("Error - sessionID : %s", mockTable.data["sessionID"])
	}

	gif, _ := base64.StdEncoding.DecodeString("R0lGODlhAQABAIABAP///wAAACwAAAAAAQABAAACAkQBADs=")
	if string(w.Body.Bytes()) != string(gif) {
		t.Errorf("Body does not contain GIF")
	}
}
