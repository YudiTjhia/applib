package api

import (
	"applib/app"
	"applib/encryption"
	"cbadmin/middleware"
	"cbadmin/svcent"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tidwall/pretty"
	"log"
	"net/http"
	"strings"
)

const (
	METHOD_POST   = "post"
	METHOD_PUT    = "put"
	METHOD_DELETE = "delete"
	METHOD_GET    = "get"
	ACTION_DETAILS = "details"
	ACTION_SEARCH  = "search"
	
)

type IBaseApi interface {
	Insert(w http.ResponseWriter, r *http.Request)
	DecodeBody(r *http.Request)
	ValidateHeader(r *http.Request) app.ErrorCollection
	OK(w http.ResponseWriter, data interface{})
	Fail(w http.ResponseWriter, err error)
	Fails(w http.ResponseWriter, errs app.ErrorCollection)
}

type BaseApi struct {
	app.App
	ServiceName string 
}

func (baseApi BaseApi) DecryptQuery(r *http.Request) (string, error) {

	val := r.URL.Query()["message"]
	if val != nil {
		decrypted, err := encryption.AES256_Decrypt(val[0])
		if err != nil {
			return "", err
		}
		return decrypted, nil
	}

	errMsg := "invalid_message"
	log.Println(errMsg)
	return "", errors.New(errMsg)
}

func (baseApi BaseApi) DecryptBody(r *http.Request) (string, error) {
	encryptedEnt := svcent.EncryptedSvcEnt{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&encryptedEnt)
	if err!=nil {
		return "", err
	}
	decrypted, err := encryption.AES256_Decrypt(encryptedEnt.Message)
	if err!=nil {
		return "", err
	}
	return decrypted, nil
}

func isRequestEmpty(r *http.Request, key string) error {
	reqAcc := r.Header.Get(key)
	if reqAcc == "" {
		return errors.New("invalid header[" + key + "]")
	}
	return nil
}

func (baseApi BaseApi) ValidateHeader(r *http.Request) app.ErrorCollection {
	appErr := app.ErrorCollection{}
	appErr.Add(isRequestEmpty(r, "requestAccount"))
	if appErr.HasErrors() {
		return appErr
	}

	appErr.Add(isRequestEmpty(r, "requestUser"))
	if appErr.HasErrors() {
		return appErr
	}

	appErr.Add(isRequestEmpty(r, "requestApp"))
	if appErr.HasErrors() {
		return appErr
	}

	appErr.Add(isRequestEmpty(r, "requestSystem"))
	if appErr.HasErrors() {
		return appErr
	}

	appErr.Add(isRequestEmpty(r, "requestSession"))
	if appErr.HasErrors() {
		return appErr
	}

	return appErr
}

func (baseApi BaseApi) GetQuery(r *http.Request, key string) string {
	val := r.URL.Query()[key]
	if val != nil {
		return val[0]
	}
	return ""
}

func (baseApi BaseApi) GetMethod(r *http.Request) string {
	return strings.ToLower(r.Method)

}

func (baseApi BaseApi) OKString(w http.ResponseWriter, data string) {

	w.WriteHeader(http.StatusOK)
	middleware.Log_("OKString>", data)
	_, err := w.Write([]byte(data))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(pretty.Pretty([]byte(data))))
}

func (baseApi BaseApi) OKPretty(w http.ResponseWriter, data interface{}) {

	if data != nil {
		w.WriteHeader(http.StatusOK)
		bytesEnt, err := json.Marshal(data)
		strPretty := string(pretty.Pretty(bytesEnt))
		_, err = w.Write([]byte(strPretty))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(strPretty)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (baseApi BaseApi) OK(w http.ResponseWriter, data interface{}) {

	if data != nil {
		w.WriteHeader(http.StatusOK)
		bytesEnt, err := json.Marshal(data)
		//middleware.Log_("OK>", string(bytesEnt))
		//fmt.Println(string(pretty.Pretty(bytesEnt)))
		_, err = w.Write(bytesEnt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}


func (baseApi BaseApi) OKEncrypted(w http.ResponseWriter, data interface{}) {

	if data != nil {
		bytesEnt, err := json.Marshal(data)
		middleware.Log_("OK>", string(bytesEnt))

		EncryptedEnt:= svcent.EncryptedSvcEnt{}
		message, err := encryption.AES256_Encrypt(string(bytesEnt))
		if err!=nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		EncryptedEnt.Message = message
		bytes1, err := json.Marshal(EncryptedEnt)
		if err!=nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Println(string(pretty.Pretty([]byte(bytes1))))
		_, err = w.Write([]byte(bytes1))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	w.WriteHeader(http.StatusOK)
}


func (baseApi BaseApi) Fail(w http.ResponseWriter, err error) {
	middleware.Log_("Fail>", err.Error())
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
}

func (baseApi BaseApi) Fails(w http.ResponseWriter, errs app.ErrorCollection) {
	b, _ := json.Marshal(errs)
	middleware.Log_("Fail>", string(b))
	http.Error(w, string(b), http.StatusInternalServerError)
	return
}

func (baseApi BaseApi) IsGet(r *http.Request) bool {
	if strings.ToLower(r.Method) == "get" {
		return true
	}
	return false
}

func (baseApi BaseApi) IsCUD(r *http.Request) bool {
	m := strings.ToLower(r.Method)
	if m == "post" || m == "put" || m == "delete" {
		return true
	}
	return false
}
