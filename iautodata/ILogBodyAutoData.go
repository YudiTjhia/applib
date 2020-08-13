package iautodata

import (
	"applib/app"
	"applib/db"
	"applib/entity"
)

type ILogBodyAutoData interface {
	db.IData
	Insert(header app.RequestHeader, logBody entity.LogBody) (string, error)
	Delete(header app.RequestHeader, logBody entity.LogBody) error
	GetList(header app.RequestHeader, logBody entity.LogBody) ([]entity.LogBody, error)
	Count(header app.RequestHeader, logBody entity.LogBody) (int64, error)
	Search(header app.RequestHeader, logBody entity.LogBody, q string) ([]entity.LogBody, error)
	GetOne(header app.RequestHeader, logBody entity.LogBody) (entity.LogBody, error)
}
