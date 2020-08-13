package searchent

import (
	"applib/app"
	entity2 "applib/entity"
	"applib/svcent"
)

type LogBodySearchEnt struct {
	app.RequestHeader
	entity2.SearchEnt
}

func (logBody LogBodySearchEnt) ToSvcEnt() svcent.LogBodySvcEnt {
	logBody1 := svcent.LogBodySvcEnt{}
	if logBody.Field == "id" {
		logBody1.ID = logBody.Value
	}

	if logBody.Field == "systemID" {
		logBody1.SystemID = logBody.Value
	}

	if logBody.Field == "appID" {
		logBody1.AppID = logBody.Value
	}

	if logBody.Field == "logID" {
		logBody1.LogID = logBody.Value
	}

	if logBody.Field == "logType" {
		logBody1.LogType = logBody.Value
	}

	if logBody.Field == "logTag" {
		logBody1.LogTag = logBody.Value
	}

	if logBody.Field == "message" {
		logBody1.Message = logBody.Value
	}

	if logBody.Field == "userID" {
		logBody1.UserID = logBody.Value
	}

	return logBody1
}
