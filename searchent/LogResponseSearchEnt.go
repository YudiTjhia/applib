package searchent

import (
	"applib/app"
	entity2 "applib/entity"
	"applib/svcent"
)

type LogResponseSearchEnt struct {
	app.RequestHeader
	entity2.SearchEnt
}

func (logResponse LogResponseSearchEnt) ToSvcEnt() svcent.LogResponseSvcEnt {
	logResponse1 := svcent.LogResponseSvcEnt{}
	if logResponse.Field == "id" {
		logResponse1.ID = logResponse.Value
	}

	if logResponse.Field == "systemID" {
		logResponse1.SystemID = logResponse.Value
	}

	if logResponse.Field == "appID" {
		logResponse1.AppID = logResponse.Value
	}

	if logResponse.Field == "serviceName" {
		logResponse1.ServiceName = logResponse.Value
	}

	return logResponse1
}
