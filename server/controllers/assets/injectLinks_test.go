package assets

import (
	"net/http"
	"net/url"
	"regexp"
	"testing"

	"github.com/simonchong/linny/server/paths"
)

var mockReq = &http.Request{URL: &url.URL{Host: "TestHost", Path: "TestPath/"}, Host: "TestHost"}

func TestInjectLinksEmpty(t *testing.T) {
	var res = InjectLinks("test", "test", mockReq)
	if res != "test" {
		t.Error("Expected 'test' got ", res)
	}
}

func TestInjectLinksILK(t *testing.T) {
	html := `test <script src="{{ilk 'someAsset.js'}}"></script> test`
	htmlRes := `test <script src="//TestHost/TestPath/someAsset.js"></script> test`
	var res = InjectLinks("test", html, mockReq)
	if res != htmlRes {
		t.Errorf("Expected '%s' got '%s'", htmlRes, res)
	}
}

func TestInjectLinksILKAbsolute(t *testing.T) {
	html := `test <script src="{{ilk '/someAsset.js'}}"></script> test`
	htmlRes := `test <script src="//TestHost/someAsset.js"></script> test`
	var res = InjectLinks("test", html, mockReq)
	if res != htmlRes {
		t.Errorf("Expected '%s' got '%s'", htmlRes, res)
	}
}

func TestInjectLinksMLK(t *testing.T) {
	html := `test <a href="{{mlk 'somePage.html'}}"></a> test`
	var res = InjectLinks("testAID", html, mockReq)
	reg := `^test <a href="//TestHost/` + paths.MeasureDir + `/k\?g=\d+&a=testAID&l=%2F%2FTestHost%2FTestPath%2FsomePage.html"></a> test$`
	matched, err := regexp.MatchString(reg, res)

	// log.Println(reg)
	if !matched || err != nil {
		t.Errorf("Not in the expected format: '%s'", res)
	}
}

func TestInjectLinksMLKTagged(t *testing.T) {
	html := `test <a href="{{mlk 'somePage.html' tag="testTAG"}}"></a> test`
	var res = InjectLinks("testAID", html, mockReq)
	reg := `^test <a href="//TestHost/` + paths.MeasureDir + `/k\?g=\d+&a=testAID&l=%2F%2FTestHost%2FTestPath%2FsomePage.html&t=testTAG"></a> test$`
	matched, err := regexp.MatchString(reg, res)

	// log.Println(reg)
	if !matched || err != nil {
		t.Errorf("Not in the expected format: '%s'", res)
	}
}

func TestInjectLinksMLKAbsolute(t *testing.T) {
	html := `test <a href="{{mlk '/somePage.html'}}"></a> test`
	var res = InjectLinks("testAID", html, mockReq)
	reg := `^test <a href="//TestHost/` + paths.MeasureDir + `/k\?g=\d+&a=testAID&l=%2F%2FTestHost%2FsomePage.html"></a> test$`
	matched, err := regexp.MatchString(reg, res)

	// log.Println(reg)
	if !matched || err != nil {
		t.Errorf("Not in the expected format: '%s'", res)
	}
}
