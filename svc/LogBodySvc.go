package svc

import (
	"applib/app"
	"applib/funcsauto"
	"applib/logger"
	"applib/svcent"
)

type LogBodySvc struct {
	app.BaseSvc
}

func (service LogBodySvc) Insert(logBody svcent.LogBodySvcEnt) (svcent.LogBodySvcEnt, app.ErrorCollection) {

	appErr := app.ErrorCollection{}
	log := logger.CreateSvcLogger(logBody.ToHeader())
	logID, err := log.Begin(logBody.ToHeader(), "LogBodySvc.Insert")
	appErr.AddCollection(err)
	log.Params(logBody.ToHeader(), logID, logBody)

	appErr.AddCollection(logBody.ValidateBase())
	if !appErr.HasErrors() {
		id, errColl := funcsauto.InsertLogBody(log, logBody)
		appErr.AddCollection(errColl)
		if !appErr.HasErrors() {
			logBody.ID = id
		}
	}

	if appErr.HasErrors() {
		log.Error(logBody.ToHeader(), logID, appErr)
	} else {
		log.LogReturn(logBody.ToHeader(), logID, logBody)
	}
	return logBody, appErr
}

func (service LogBodySvc) Delete(logBody svcent.LogBodySvcEnt) app.ErrorCollection {
	appErr := app.ErrorCollection{}
	log := logger.CreateSvcLogger(logBody.ToHeader())
	logID, errs := log.Begin(logBody.ToHeader(), "LogBodySvc.Delete")
	appErr.AddCollection(errs)

	log.Params(logBody.ToHeader(), logID, logBody)
	appErr.AddCollection(logBody.ValidateBaseID())

	if !appErr.HasErrors() {
		errColl := funcsauto.DeleteLogBody(log, logBody)
		appErr.AddCollection(errColl)
	}

	if appErr.HasErrors() {
		log.Error(logBody.ToHeader(), logID, appErr)
	} else {
		log.LogReturn(logBody.ToHeader(), logID, "OK")
	}

	return appErr
}

func (service LogBodySvc) GetList(logBody svcent.LogBodySvcEnt) ([]svcent.LogBodySvcEnt, app.ErrorCollection) {
	logBodys := []svcent.LogBodySvcEnt{}
	appErr := app.ErrorCollection{}
	if !appErr.HasErrors() {
		logBodys = funcsauto.GetListLogBody(logBody)
	}
	return logBodys, appErr
}

func (service LogBodySvc) Search(logBody svcent.LogBodySvcEnt, q string) ([]svcent.LogBodySvcEnt, app.ErrorCollection) {
	logBodys := []svcent.LogBodySvcEnt{}
	appErr := app.ErrorCollection{}
	if !appErr.HasErrors() {
		logBodys1, errColl := funcsauto.SearchLogBody(logBody, q)
		appErr.AddCollection(errColl)
		logBodys = logBodys1
	}
	return logBodys, appErr
}

func (service LogBodySvc) GetOne(logBody svcent.LogBodySvcEnt) (svcent.LogBodySvcEnt, app.ErrorCollection) {
	var logBody1 svcent.LogBodySvcEnt
	appErr := app.ErrorCollection{}
	appErr.AddCollection(logBody.ValidateBaseID())
	if !appErr.HasErrors() {
		logBody2, errColl := funcsauto.GetOneLogBody(logBody)
		appErr.AddCollection(errColl)
		if !appErr.HasErrors() {
			logBody1 = logBody2
		}
	}
	return logBody1, appErr
}
