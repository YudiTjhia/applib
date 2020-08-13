package iautodata

import (
	"applib/app"
	"applib/db"
	"applib/entity"
)

type ILogResponseAutoData interface {
	db.IData
	Insert(header app.RequestHeader, logResponse entity.LogResponse) (string, error)
	Delete(header app.RequestHeader, logResponse entity.LogResponse) error
	GetList(header app.RequestHeader, logResponse entity.LogResponse) ([]entity.LogResponse, error)
	Count(header app.RequestHeader, logResponse entity.LogResponse) (int64, error)
	Search(header app.RequestHeader, logResponse entity.LogResponse, q string) ([]entity.LogResponse, error)
	GetOne(header app.RequestHeader, logResponse entity.LogResponse) (entity.LogResponse, error)
}
