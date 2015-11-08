package assets

import (
	"net/http"
	"net/url"
	"testing"
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
