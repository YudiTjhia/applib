package isvc

import (
	"applib/app"
	"applib/svcent"
)

type ILogBodySvc interface {
	Insert(logBody svcent.LogBodySvcEnt) (svcent.LogBodySvcEnt, app.ErrorCollection)
	Delete(logBody svcent.LogBodySvcEnt) app.ErrorCollection
	GetList(logBody svcent.LogBodySvcEnt) ([]svcent.LogBodySvcEnt, app.ErrorCollection)
	Search(logBody svcent.LogBodySvcEnt, q string) ([]svcent.LogBodySvcEnt, app.ErrorCollection)
	GetOne(logBody svcent.LogBodySvcEnt) (svcent.LogBodySvcEnt, app.ErrorCollection)
}
