package controllers

import (
	"database/sql"
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

type mockConTable struct {
	data map[string]string
}

func (m *mockConTable) Init() {}
func (m *mockConTable) Insert(adID string, refererURL string, originIP string, jsGeneratedOn time.Time, conversionTag string, sessionID string) (sql.Result, error) {
	m.data["adID"] = adID
	m.data["refererURL"] = refererURL
	m.data["originIP"] = originIP
	m.data["jsGeneratedOn"] = jsGeneratedOn.String()
	m.data["conversionTag"] = conversionTag
	m.data["sessionID"] = sessionID

	return &datatest.MockDBResult{1, 1}, nil
}

func TestConversions(t *testing.T) {

	var mockTable = &mockConTable{map[string]string{}}
	var mockData = &insights.Data{AdConversions: mockTable}
	var mockAppContext = &wrappers.AppContext{Data: mockData}
	var mockReq, _ = http.NewRequest("GET", "/m/k", nil)
	mockReq.RemoteAddr = "127.0.0.1:80"
	mockReq.Form = map[string][]string{
		"g": []string{"1447100000"},
		"a": []string{"ADID123"},
		"l": []string{"http://www.test.test"},
		"t": []string{"TAG123"},
	}
	var sessionID = session.MakeSessionID()
	mockReq.AddCookie(conversions.NewCookie("ADID123"))
	mockReq.AddCookie(session.MakeSessionCookie(sessionID))
	mockReq.Header.Add("Referer", "http://test.test")

	var w = httptest.NewRecorder()

	var code, err = Conversions(mockAppContext, web.C{}, w, mockReq)
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
	if mockTable.data["jsGeneratedOn"] != "2015-11-10 07:13:20 +1100 AEDT" {
		t.Errorf("Error - jsGeneratedOn : %s", mockTable.data["jsGeneratedOn"])
	}
	if mockTable.data["conversionTag"] != "TAG123" {
		t.Errorf("Error - conversionTag : %s", mockTable.data["conversionTag"])
	}
	if mockTable.data["sessionID"] != sessionID {
		t.Errorf("Error - sessionID : %s", mockTable.data["sessionID"])
	}
}

func TestConversionsTagLimit(t *testing.T) {

	var mockTable = &mockConTable{map[string]string{}}
	var mockData = &insights.Data{AdConversions: mockTable}
	var mockAppContext = &wrappers.AppContext{Data: mockData}
	var mockReq, _ = http.NewRequest("GET", "/m/k", nil)
	mockReq.RemoteAddr = "127.0.0.1:80"
	mockReq.Form = map[string][]string{
		"g": []string{"1447100000"},
		"a": []string{"ADID123"},
		"l": []string{"http://www.test.test"},
		"t": []string{"123456678123456678123456678123456678123456678123456678123456678123456678123456678123456678123456678123456678123456678123456678123456678"},
	}
	var sessionID = session.MakeSessionID()
	mockReq.AddCookie(conversions.NewCookie("ADID123"))
	mockReq.AddCookie(session.MakeSessionCookie(sessionID))
	mockReq.Header.Add("Referer", "http://test.test")

	var w = httptest.NewRecorder()

	var code, err = Conversions(mockAppContext, web.C{}, w, mockReq)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if code != 200 {
		t.Errorf("Status: %d", code)
	}
	if len(mockTable.data["conversionTag"]) > 64 {
		t.Errorf("Error - conversionTag : %s", mockTable.data["conversionTag"])
	}

}
