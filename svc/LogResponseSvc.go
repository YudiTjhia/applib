package svc

import (
	"applib/app"
	"applib/funcsauto"
	"applib/svcent"
	"encoding/json"
	"time"
)

type LogResponseSvc struct {
	app.BaseSvc
	logRespSvcEnt svcent.LogResponseSvcEnt
}

func (service *LogResponseSvc) Start(header app.RequestHeader, serviceName string) {
	service.logRespSvcEnt = svcent.LogResponseSvcEnt{}
	service.logRespSvcEnt.FromHeader(header)
	service.logRespSvcEnt.SystemID = header.RequestSystem
	service.logRespSvcEnt.AppID = header.RequestApp
	service.logRespSvcEnt.ServiceName = serviceName
	service.logRespSvcEnt.RequestTime = time.Now()
}

func (service *LogResponseSvc) End(data interface{}) {
	b, err := json.Marshal(data)
	if err == nil {
		service.logRespSvcEnt.DataSize = len(b)
	}
	service.logRespSvcEnt.ResponseTime = time.Now()
	service.logRespSvcEnt.ResponseDuration = service.logRespSvcEnt.ResponseTime.Sub(service.logRespSvcEnt.RequestTime).Seconds() * 1000
	service.Insert(service.logRespSvcEnt)
}

func (service LogResponseSvc) Insert(logResponse svcent.LogResponseSvcEnt) (svcent.LogResponseSvcEnt, app.ErrorCollection) {

	appErr := app.ErrorCollection{}
	appErr.AddCollection(logResponse.ValidateBase())
	if !appErr.HasErrors() {
		id, errColl := funcsauto.InsertLogResponse(logResponse)
		appErr.AddCollection(errColl)
		if !appErr.HasErrors() {
			logResponse.ID = id
		}
	}
	return logResponse, appErr
}

func (service LogResponseSvc) Delete(logResponse svcent.LogResponseSvcEnt) app.ErrorCollection {
	appErr := app.ErrorCollection{}
	appErr.AddCollection(logResponse.ValidateBaseID())
	if !appErr.HasErrors() {
		errColl := funcsauto.DeleteLogResponse(logResponse)
		appErr.AddCollection(errColl)
	}
	return appErr
}

func (service LogResponseSvc) GetList(logResponse svcent.LogResponseSvcEnt) ([]svcent.LogResponseSvcEnt, app.ErrorCollection) {
	logResponses := []svcent.LogResponseSvcEnt{}
	appErr := app.ErrorCollection{}
	if !appErr.HasErrors() {
		logResponses = funcsauto.GetListLogResponse(logResponse)
	}
	return logResponses, appErr
}

func (service LogResponseSvc) Search(logResponse svcent.LogResponseSvcEnt, q string) ([]svcent.LogResponseSvcEnt, app.ErrorCollection) {
	logResponses := []svcent.LogResponseSvcEnt{}
	appErr := app.ErrorCollection{}
	if !appErr.HasErrors() {
		logResponses1, errColl := funcsauto.SearchLogResponse(logResponse, q)
		appErr.AddCollection(errColl)
		logResponses = logResponses1
	}
	return logResponses, appErr
}

func (service LogResponseSvc) GetOne(logResponse svcent.LogResponseSvcEnt) (svcent.LogResponseSvcEnt, app.ErrorCollection) {
	var logResponse1 svcent.LogResponseSvcEnt
	appErr := app.ErrorCollection{}
	appErr.AddCollection(logResponse.ValidateBaseID())
	if !appErr.HasErrors() {
		logResponse2, errColl := funcsauto.GetOneLogResponse(logResponse)
		appErr.AddCollection(errColl)
		if !appErr.HasErrors() {
			logResponse1 = logResponse2
		}
	}
	return logResponse1, appErr
}
