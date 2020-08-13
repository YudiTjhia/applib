package isvc

import (
	"applib/app"
	"applib/svcent"
)

type ILogResponseSvc interface {
	Insert(logResponse svcent.LogResponseSvcEnt) (svcent.LogResponseSvcEnt, app.ErrorCollection)
	Delete(logResponse svcent.LogResponseSvcEnt) app.ErrorCollection
	GetList(logResponse svcent.LogResponseSvcEnt) ([]svcent.LogResponseSvcEnt, app.ErrorCollection)
	Search(logResponse svcent.LogResponseSvcEnt, q string) ([]svcent.LogResponseSvcEnt, app.ErrorCollection)
	GetOne(logResponse svcent.LogResponseSvcEnt) (svcent.LogResponseSvcEnt, app.ErrorCollection)
}
