package logger

import (
	"applib/app"
	"applib/entity"
	"github.com/jinzhu/gorm"
	"time"
)

type ILogBodyData interface {
	Insert(accountID string, userID string, systemID string, appID string,
		logID string, logType string, logTag string, message string)

	InsertCount(accountID string, userID string, systemID string, appID string,
		logID string, count int)

	InsertInterface(accountID string, userID string, systemID string, appID string,
		logID string, logType string, logTag string, message interface{})

	InsertData(accountID string, userID string, systemID string, appID string,
		logID string, logTag string, message interface{})

	InsertParams(accountID string, userID string, systemID string, appID string,
		logID string, params interface{})

	InsertQuery(accountID string, userID string, systemID string, appID string,
		logID string, query *gorm.DB)

	GetList(accountID string, fromDate time.Time, toDate time.Time,
		logType string) ([]entity.LogBody, error)
}

type ILogBodyData2 interface {
	Insert(header app.RequestHeader,
		logID string, logType string, logTag string, message string)

	InsertCount(header app.RequestHeader,
		logID string, count int)

	InsertInterface(header app.RequestHeader,
		logID string, logType string, logTag string, message interface{})

	InsertData(header app.RequestHeader,
		logID string, logTag string, message interface{})

	InsertParams(header app.RequestHeader,
		logID string, params interface{})

	InsertQuery(header app.RequestHeader,
		logID string, query *gorm.DB)

	GetList(header app.RequestHeader,
		fromDate time.Time, toDate time.Time,
		logType string) ([]entity.LogBody, error)
}
