package db

import (
	"applib/app"
	"applib/conf"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"go.mongodb.org/mongo-driver/mongo"
)

type CApi struct {
	LogID        string
	ConnectionID string
	baseUrl      string
}

func (cApi *CApi) SetDataConnection(connection DataConnection) {
	cApi.baseUrl = connection.CApi.baseUrl
	cApi.ConnectionID = connection.CApi.ConnectionID
}

func (cApi *CApi) SetConnectionID(connectionID string) {
	cApi.ConnectionID = connectionID
}

func (cApi CApi) HeaderToMap(header app.RequestHeader) map[string]string {
	mapHead := map[string]string{}
	mapHead["requestAccount"] = header.RequestAccount
	mapHead["requestSystem"] = header.RequestSystem
	mapHead["requestApp"] = header.RequestApp
	mapHead["requestUser"] = header.RequestUser
	mapHead["requestSession"] = header.RequestSession
	return mapHead
}

func (cApi CApi) GetGormDB() *gorm.DB {
	return nil
}

func (cApi CApi) GetMongoDB() *mongo.Database {
	return nil
}

func (cApi *CApi) SetLogID(logID string) {
	cApi.LogID = logID
}

func (cApi CApi) send_(method string, url string, data interface{}) ([]byte, error) {

	var request *http.Request
	if method == "GET" {
		request, _ = http.NewRequest(method, cApi.baseUrl+url, nil)
	} else {
		body, err := json.Marshal(data)
		if err != nil {
			return []byte{}, err
		}
		request, _ = http.NewRequest(method, cApi.baseUrl+url, bytes.NewBuffer(body))
	}

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)
	strRespBody := string(respBody)

	if resp.StatusCode == http.StatusOK {
		return respBody, nil

	} else {
		return []byte{}, errors.New(strRespBody)
	}
}

func makeQueryUrl(data map[string]string) string {
	queryParams := ""
	for key, val := range data {
		queryParams += key + "=" + val + "&"
	}
	if len(queryParams) > 0 {
		queryParams = queryParams[0 : len(queryParams)-1]
	}
	return queryParams
}

func (cApi *CApi) SetBaseUrl(pBaseUrl string) {
	cApi.baseUrl = pBaseUrl
}

func (cApi CApi) Send(method string, url string, data interface{}, header map[string]string) ([]byte, error) {
	return cApi.send(method, url, data, header)
}
func (cApi CApi) send(method string, url string, data interface{}, header map[string]string) ([]byte, error) {

	var request *http.Request
	var err error

	if method == "GET" {
		//queryParams := makeQueryUrl(data.(map[string]string))
		queryUrl := cApi.baseUrl + url
		//if len(queryParams) > 0 {
		//	queryUrl += "?" + queryParams
		//}
		request, err = http.NewRequest(method, queryUrl, nil)
		if(err!=nil ){
			fmt.Println(err)
			return []byte{}, err
		}
		request.Header.Set("Content-Type", "application/json")
		pBody := data.(map[string]string)
		q := request.URL.Query()
		for key, val:=range pBody {
			q.Add(key, val)
		}
		request.URL.RawQuery = q.Encode()

	} else {
		pBody, err := json.Marshal(data)
		if err != nil {
			fmt.Println(err)
			return []byte{}, err
		}
		request, err = http.NewRequest(method, cApi.baseUrl+url, bytes.NewBuffer(pBody))
		if(err!=nil ){
			fmt.Println(err)
			return []byte{}, err
		}
		request.Header.Set("Content-Type", "application/json")
	}

	for key, head := range header {
		request.Header.Add(key, head)
	}

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)
	strRespBody := string(respBody)

	if resp.StatusCode == http.StatusOK {
		return respBody, nil

	} else {
		return []byte{}, errors.New(strRespBody)
	}
}

func (cApi CApi) GET(url string, data interface{}, header map[string]string) ([]byte, error) {
	return cApi.send("GET", url, data, header)
}

func (cApi CApi) POST(url string, data interface{}, header map[string]string) ([]byte, error) {
	return cApi.send("POST", url, data, header)
}

func (cApi CApi) PUT(url string, data interface{}, header map[string]string) ([]byte, error) {
	return cApi.send("PUT", url, data, header)
}

func (cApi CApi) DELETE(url string, data interface{}, header map[string]string) ([]byte, error) {
	return cApi.send("DELETE", url, data, header)
}

func createCApiConnString(cApiConf conf.DbConf) string {

	connString := ""
	if cApiConf.Protocol == "" {
		connString += "http://"
	} else {
		connString += cApiConf.Protocol + "://"
	}

	connString += cApiConf.Server
	if cApiConf.Port != 0 && cApiConf.Port != 80 {
		connString += ":" + strconv.Itoa(cApiConf.Port)
	}

	connString += cApiConf.BasePath
	return connString

}

func CreateCApiConnection(configFile string, connectionID string) CApi {

	conf := conf.GetDbConfig(configFile, connectionID)
	connString := createCApiConnString(conf)

	cApi := CApi{}
	cApi.ConnectionID = connectionID
	cApi.baseUrl = connString

	return cApi

}
