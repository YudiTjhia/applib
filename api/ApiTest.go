package api

import (
	"applib/db"
	"bytes"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ApiTest struct {
	T         *testing.T
	MuxRouter *mux.Router
	db        *gorm.DB
}

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("An error occurred. %v", err)
	}
}

func (at *ApiTest) Init(t *testing.T) {
	at.T = t
	at.MuxRouter = mux.NewRouter()
	at.db = db.CreateGormConnection("", "")
}

func (at *ApiTest) Close() {
	at.db.Close()
}

func (at *ApiTest) AddRouter(path string, f func(w http.ResponseWriter, r *http.Request)) {
	at.MuxRouter.HandleFunc(path, f)
}

func (at *ApiTest) GET(url string, success func(body []byte), fail func(body []byte)) {
	request, err := http.NewRequest("GET", url, nil)
	checkError(err, at.T)

	recorder := httptest.NewRecorder()
	at.MuxRouter.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		fail(recorder.Body.Bytes())
	}
	success(recorder.Body.Bytes())
}

func (at *ApiTest) POST(url string, body []byte, success func(body []byte), fail func(body []byte)) {

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	checkError(err, at.T)

	recorder := httptest.NewRecorder()
	at.MuxRouter.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		fail(recorder.Body.Bytes())
	}

	success(recorder.Body.Bytes())

}

func (at *ApiTest) PUT(url string, body []byte, success func(body []byte), fail func(body []byte)) {

	request, err := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	checkError(err, at.T)

	recorder := httptest.NewRecorder()
	at.MuxRouter.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		fail(recorder.Body.Bytes())
	}

	success(recorder.Body.Bytes())

}

func (at *ApiTest) DELETE(url string, body []byte, success func(body []byte), fail func(body []byte)) {

	request, err := http.NewRequest("DELETE", url, bytes.NewBuffer(body))
	checkError(err, at.T)

	recorder := httptest.NewRecorder()
	at.MuxRouter.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		fail(recorder.Body.Bytes())
	}

	success(recorder.Body.Bytes())

}
