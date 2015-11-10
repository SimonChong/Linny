package controllers

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/simonchong/linny/insights"
	"github.com/simonchong/linny/insights/datatest"
	"github.com/simonchong/linny/server/wrappers"
	"github.com/zenazn/goji/web"
)

type mockACTable struct {
	data map[string]string
}

func (m *mockACTable) Init() {}
func (m *mockACTable) Insert(adID string, refererURL string, destinationURL string, originIP string, linkGeneratedOn time.Time, linkTag string, sessionID string) (sql.Result, error) {
	m.data["adID"] = adID
	m.data["refererURL"] = refererURL
	m.data["destinationURL"] = destinationURL
	m.data["originIP"] = originIP
	m.data["linkGeneratedOn"] = linkGeneratedOn.String()
	m.data["linkTag"] = linkTag
	m.data["sessionID"] = sessionID
	return &datatest.MockDBResult{1, 1}, nil
}

func TestClickTracking(t *testing.T) {

	var mockTable = &mockACTable{map[string]string{}}
	var mockData = &insights.Data{AdClickThroughs: mockTable}
	var mockAppContext = &wrappers.AppContext{Data: mockData}
	var mockReq, _ = http.NewRequest("GET", "/m/k", nil)
	mockReq.RemoteAddr = "127.0.0.1:80"
	mockReq.Form = map[string][]string{
		"g": []string{"1447100000"},
		"a": []string{"ADID123"},
		"l": []string{"http://www.test.test"},
		"t": []string{"TAG123"},
	}
	var w = httptest.NewRecorder()

	var code, err = ClickTracking(mockAppContext, "12345", web.C{}, w, mockReq)
	//TEST it inserts the click into the DB
	if mockTable.data["adID"] != "ADID123" {
		t.Errorf("adID: %s", mockTable.data["adID"])
	}
	if mockTable.data["refererURL"] != "" {
		t.Errorf("refererURL: %s", mockTable.data["refererURL"])
	}
	if mockTable.data["destinationURL"] != "http://www.test.test" {
		t.Errorf("destinationURL: %s", mockTable.data["destinationURL"])
	}
	if mockTable.data["originIP"] != "127.0.0.1" {
		t.Errorf("originIP: %s", mockTable.data["originIP"])
	}
	if mockTable.data["linkGeneratedOn"] != "2015-11-10 07:13:20 +1100 AEDT" {
		t.Errorf("linkGeneratedOn: %s", mockTable.data["linkGeneratedOn"])
	}
	if mockTable.data["linkTag"] != "TAG123" {
		t.Errorf("linkTag: %s", mockTable.data["linkTag"])
	}
	if mockTable.data["sessionID"] != "12345" {
		t.Errorf("sessionID: %s", mockTable.data["sessionID"])
	}
	//TEST it redirects 301
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if code != 301 {
		t.Errorf("Status: %d", code)
	}
}
