package app

import (
	"applib/util"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"fmt"
)

type App struct {
	//RequestAccount string
	//UserID string
	DB *gorm.DB
}

func (app *App) SetApp(pApp App) {
	//app.RequestAccount = pApp.RequestAccount
	app.DB = pApp.DB
}

func (app App) GetApp() App {
	app1 := App{
		//RequestAccount: app.RequestAccount,
		DB: app.DB,
	}
	return app1
}

func (app App) GetDb() *gorm.DB {
	return app.DB
}

func (app *App) SetDb(db *gorm.DB) {
	app.DB = db
}

func DecodeRequestHeader2(r *http.Request) RequestHeader {
	header := RequestHeader{}
	header.RequestAccount = r.Header.Get("requestAccount")
	header.RequestUser = r.Header.Get("requestUser")
	header.RequestApp = r.Header.Get("requestApp")
	header.RequestSystem = r.Header.Get("requestSystem")
	header.RequestSession = r.Header.Get("requestSession")
	intPage, err := strconv.Atoi(r.Header.Get("page"))
	if err == nil {
		header.Page = intPage
	}
	intPageSize, err := strconv.Atoi(r.Header.Get("pageSize"))
	if err == nil {
		header.PageSize = intPageSize
	}
	bUsePaging, err := strconv.ParseBool(r.Header.Get("usePaging"))
	if err == nil {
		header.UsePaging = bUsePaging
	}

	return header
}



func DecodeRequestHeader(r *http.Request) RequestHeader {
	header := RequestHeader{}
	header.RequestAccount = getRequestQuery(r, "requestAccount")
	header.RequestUser = getRequestQuery(r, "requestUser")
	header.RequestApp = getRequestQuery(r, "requestApp")
	header.RequestSystem = getRequestQuery(r, "requestSystem")
	//header.RequestService = getRequestQuery(r,"requestService")
	header.RequestSession = getRequestQuery(r, "requestSession")

	intPage, err := strconv.Atoi(getRequestQuery(r, "page"))
	if err == nil {
		header.Page = intPage
	}
	intPageSize, err := strconv.Atoi(getRequestQuery(r, "pageSize"))
	if err == nil {
		header.PageSize = intPageSize
	}
	bUsePaging, err := strconv.ParseBool(getRequestQuery(r, "usePaging"))
	if err == nil {
		header.UsePaging = bUsePaging
	}

	fmt.Println("header", util.ToJson(header))

	return header
}

func ParseHeader(r *http.Request) RequestHeader {
	header := RequestHeader{}
	header.RequestAccount = r.Header.Get("Request-Account")
	header.RequestUser = r.Header.Get("Request-User")
	header.RequestApp = r.Header.Get("Request-App")
	header.RequestSystem = r.Header.Get("Request-System")
	header.RequestSession = r.Header.Get("Request-Session")
	intPage, err := strconv.Atoi(r.Header.Get("page"))
	if err == nil {
		header.Page = intPage
	}
	intPageSize, err := strconv.Atoi(r.Header.Get("Page-Size"))
	if err == nil {
		header.PageSize = intPageSize
	}
	bUsePaging, err := strconv.ParseBool(r.Header.Get("Use-Paging"))
	if err == nil {
		header.UsePaging = bUsePaging
	}
	fmt.Println("header", util.ToJson(header))
	return header
}

func DecodeQ(r *http.Request) string {
	return getRequestQuery(r, "q")
}

func getRequestQuery(r *http.Request, key string) string {
	val := r.URL.Query()[key]
	fmt.Println("getRequestQuery[" + key + "]=", val)
	if val != nil {
		return val[0]
	}
	return ""
}


func ParseHeaderKey(r *http.Request, key string) string {
	val := r.Header.Get(key)
	fmt.Println("ParseQueryKey[" + key + "]=", val)
	return val
}
