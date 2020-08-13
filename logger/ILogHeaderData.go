package logger

import (
	"applib/app"
	"applib/entity"
	"time"
)

type ILogHeaderData interface {
	Insert(accountID string, userID string, systemID string, appID string, dbInfo string) (string, error)

	GetList(accountID string, systemID string, fromDate time.Time, toDate time.Time) ([]entity.LogHeader, error)

	Search(accountID string, q string) ([]entity.LogHeader, error)
}

type ILogHeaderData2 interface {
	Insert(header app.RequestHeader, dbInfo string) (string, error)

	GetList(header app.RequestHeader, fromDate time.Time, toDate time.Time) ([]entity.LogHeader, error)

	Search(header app.RequestHeader, q string) ([]entity.LogHeader, error)
}
