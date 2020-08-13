package app

import (
	"applib/util"
	"net/http"
	"strconv"
	"time"
)

const (
	INVALID_REQUEST_ACCOUNT = "INVALID_REQUEST_ACCOUNT"
	INVALID_REQUEST_USER    = "INVALID_REQUEST_USER"
	INVALID_REQUEST_APP     = "INVALID_REQUEST_APP"
	INVALID_REQUEST_SYSTEM  = "INVALID_REQUEST_SYSTEM"
	INVALID_REQUEST_SERVICE = "INVALID_REQUEST_SERVICE"
	INVALID_REQUEST_SESSION = "INVALID_REQUEST_SESSION"
)

type RequestHeader struct {

	RequestAccount  string `json:"requestAccount"`
	RequestSystem   string `json:"requestSystem"`
	RequestUser     string `json:"requestUser"`
	RequestSession  string `json:"requestSession"`
	RequestApp      string `json:"requestApp"`
	RequestService  string `json:"requestService"`
	RequestEmployee string `json:"requestEmployee"`
	SessionStatus   bool   `json:"sessionStatus"`
	SessionDate     time.Time `json:"sessionDate"`
	ExpiredDate     time.Time `json:"expiredDate"`


	Page      int  `json:"page" gorm:"-"`
	PageSize  int  `json:"pageSize" gorm:"-"`
	UsePaging bool `json:"usePaging" gorm:"-"`

	AppConfigFile AppConfigFile `json:"-"`

}

func (header RequestHeader) GetOffset() int64 {
	return int64((header.Page - 1) * header.PageSize)
}

func(header RequestHeader) LogPaging(logger func(msg ...interface{})) {
	METHOD := "RequestHeader>LogPaging>"
	logger(header.RequestUser, METHOD, "Page>", header.Page)
	logger(header.RequestUser, METHOD, "PageSize>", header.PageSize)
	logger(header.RequestUser, METHOD, "offset>", header.GetOffset())
}

func (request RequestHeader) ToMap() map[string]string {
	mp := map[string]string{}
	mp["requestAccount"] = request.RequestAccount
	mp["requestSystem"] = request.RequestSystem
	mp["requestUser"] = request.RequestUser
	mp["requestApp"] = request.RequestApp
	mp["requestEmployee"] = request.RequestEmployee
	mp["requestApp"] = request.RequestApp
	mp["requestService"] = request.RequestService
	mp["requestSession"] = request.RequestSession
	mp["requestAccount"] = request.RequestAccount
	return mp
}

func (request *RequestHeader) FromHttpRequest(r *http.Request) {
	request.RequestAccount = r.Header.Get("requestAccount")
	request.RequestUser = r.Header.Get("requestUser")
	request.RequestSystem = r.Header.Get("requestSystem")
	request.RequestApp = r.Header.Get("requestApp")

	intPage, err := strconv.Atoi(r.Header.Get("page"))
	if err == nil {
		request.Page = intPage
	}

	intPageSize, err := strconv.Atoi(r.Header.Get("pageSize"))
	if err == nil {
		request.PageSize = intPageSize
	}

	bUsePaging, err := strconv.ParseBool(r.Header.Get("usePaging"))
	if err == nil {
		request.UsePaging = bUsePaging
	}

}

func (request *RequestHeader) SetValues(accountID string, requestUser string,
	requestApp string, requestSystem string, requestService string,
	requestSession string) {

	request.RequestAccount = accountID
	request.RequestUser = requestUser
	request.RequestApp = requestApp
	request.RequestSystem = requestSystem
	request.RequestService = requestService
	request.RequestSession = requestSession
}

func (request RequestHeader) validateAccountID() error {
	return util.Required(request.RequestAccount, INVALID_REQUEST_ACCOUNT)
}
func (request RequestHeader) validateRequestUser() error {
	return util.Required(request.RequestUser, INVALID_REQUEST_USER)
}
func (request RequestHeader) validateRequestApp() error {
	return util.Required(request.RequestApp, INVALID_REQUEST_APP)
}
func (request RequestHeader) validateRequestSystem() error {
	return util.Required(request.RequestSystem, INVALID_REQUEST_SYSTEM)
}

func (request RequestHeader) validateRequestService() error {
	return util.Required(request.RequestService, INVALID_REQUEST_SERVICE)
}

func (request RequestHeader) validateRequestSession() error {
	return util.Required(request.RequestSession, INVALID_REQUEST_SESSION)
}

func (header *RequestHeader) DecodeRequestHeader(r *http.Request) ErrorCollection {
	header.RequestAccount = getRequestQuery(r, "requestAccount")
	header.RequestUser = getRequestQuery(r, "requestUser")
	header.RequestApp = getRequestQuery(r, "requestApp")
	header.RequestSystem = getRequestQuery(r, "requestSystem")
	header.RequestSession = getRequestQuery(r, "requestSession")

	//header.Page = getRequestQuery(r, "page")
	//header.PageSize = getRequestQuery(r, "pageSize")
	//header.UsePaging = getRequestQuery(r, "usePaging")

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

	return header.ValidateHeader()
}

func (header *RequestHeader) DecodeOnlyRequestHeader(r *http.Request) {
	header.RequestAccount = getRequestQuery(r, "requestAccount")
	header.RequestUser = getRequestQuery(r, "requestUser")
	header.RequestApp = getRequestQuery(r, "requestApp")
	header.RequestSystem = getRequestQuery(r, "requestSystem")
	header.RequestSession = getRequestQuery(r, "requestSession")

}

func (request RequestHeader) ValidateHeader() ErrorCollection {

	appErr := ErrorCollection{}
	err := request.validateAccountID()
	appErr.Add(err)
	if appErr.HasErrors() {
		return appErr
	}

	err = request.validateRequestUser()
	appErr.Add(err)
	if appErr.HasErrors() {
		return appErr
	}

	err = request.validateRequestApp()
	appErr.Add(err)
	if appErr.HasErrors() {
		return appErr
	}

	err = request.validateRequestSystem()
	appErr.Add(err)
	if appErr.HasErrors() {
		return appErr
	}

	err = request.validateRequestSession()
	appErr.Add(err)
	if appErr.HasErrors() {
		return appErr
	}

	return appErr
}


func (request RequestHeader) ValidateUserAndSession() error {

	err := util.Required(request.RequestUser, INVALID_REQUEST_ACCOUNT)
	if err != nil { return err }

	err = util.Required(request.RequestSession, INVALID_REQUEST_SESSION)
	if err != nil { return err }

	return nil
}

func (request *RequestHeader) SetAppConfigFile(appConfigFile AppConfigFile) {
	request.AppConfigFile = appConfigFile
}

func (request RequestHeader) ToHeader() RequestHeader {
	n := RequestHeader{}
	n.RequestSystem = request.RequestSystem
	n.RequestAccount = request.RequestAccount
	n.RequestService = request.RequestService
	n.RequestApp = request.RequestApp
	n.RequestUser = request.RequestUser

	n.RequestSession = request.RequestSession
	n.AppConfigFile = request.AppConfigFile

	n.PageSize = request.PageSize
	n.Page = request.Page
	n.UsePaging = request.UsePaging

	return n
}

func (request *RequestHeader) FromHeader(header RequestHeader) {
	request.RequestSystem = header.RequestSystem
	request.RequestAccount = header.RequestAccount
	request.RequestService = header.RequestService
	request.RequestApp = header.RequestApp
	request.RequestUser = header.RequestUser
	request.RequestSession = header.RequestSession
	request.AppConfigFile = header.AppConfigFile

	request.Page = header.Page
	request.PageSize = header.PageSize
	request.UsePaging = header.UsePaging

}

func (request RequestHeader) ToQueryParams() string {

	queryParams := "requestSystem=" + request.RequestSystem
	queryParams += "&requestAccount=" + request.RequestAccount
	queryParams += "&requestService=" + request.RequestService
	queryParams += "&requestApp=" + request.RequestApp
	queryParams += "&requestUser=" + request.RequestUser
	queryParams += "&requestSession=" + request.RequestSession
	queryParams += "&page=" + strconv.Itoa(request.Page)
	queryParams += "&pageSize=" + strconv.Itoa(request.PageSize)
	queryParams += "&usePaging=" + strconv.FormatBool(request.UsePaging)

	return queryParams

}
