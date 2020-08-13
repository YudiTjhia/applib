package testhelper

import (
	"applib/app"
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Expected(label string, expect interface{}, actual interface{}, t *testing.T) {
	if expect!= actual {
		t.Fatal(label + ",expected=", expect, ",actual=", actual)
	}
}

func Contains(label string, expect string, actual string, t *testing.T) {
	if !strings.Contains(strings.ToLower(actual), strings.ToLower(expect)) {
		t.Fatal(label + ",expected=", expect, ",actual=", actual)
	}
}

func Send(t *testing.T, method string, url string, pHeader app.RequestHeader,  body interface{},
	handlerFunc func(w http.ResponseWriter, r *http.Request)) (*httptest.ResponseRecorder, bool) {
	var req *http.Request
	var err error
	if method == "GET" || method == "get" {
		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			t.Fatal(err)
		}

		pBody := body.(map[string]string)
		q := req.URL.Query()
		for key, val:=range pBody {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	} else {
		pBody := body.([]byte)
		req, err = http.NewRequest(method, url, bytes.NewBuffer(pBody))
		if err != nil {
			t.Fatal(err)
		}

	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("requestUser", pHeader.RequestUser)
	req.Header.Set("requestSession", pHeader.RequestSession)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlerFunc)
	handler.ServeHTTP(rr, req)

	b:= true
	if status := rr.Code; status != http.StatusOK {
		b = false
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		t.Fatal(rr.Body.String())
	}
	return rr, b
}
