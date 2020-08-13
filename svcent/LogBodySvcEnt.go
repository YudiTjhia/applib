package svcent

import (
	"applib/app"
	"applib/constant"
	"applib/entity"
	"applib/funcs"
	"applib/validator"
	libvalidator "applib/validator"
	"net/http"
)

type LogBodySvcEnt struct {
	app.BaseSvcEnt
	app.RequestHeader
	entity.LogBody
	objectValidator libvalidator.IObjectValidator
}

func (logBody *LogBodySvcEnt) DecodeQuery(r *http.Request) {
	logBody.ID = logBody.DecodeQueryKey(r, "id")
	logBody.SystemID = logBody.DecodeQueryKey(r, "systemID")
	logBody.AppID = logBody.DecodeQueryKey(r, "appID")
	logBody.LogID = logBody.DecodeQueryKey(r, "logID")
	logBody.LogType = logBody.DecodeQueryKey(r, "logType")
	logBody.LogTag = logBody.DecodeQueryKey(r, "logTag")
	logBody.Message = logBody.DecodeQueryKey(r, "message")

	logBody.UserID = logBody.DecodeQueryKey(r, "userID")

}

func (logBody LogBodySvcEnt) GetValue(name string) (interface{}, error) {
	switch name {
	case "ID":
		return logBody.ID, nil
	case "SystemID":
		return logBody.SystemID, nil
	case "AppID":
		return logBody.AppID, nil
	case "LogID":
		return logBody.LogID, nil
	case "LogType":
		return logBody.LogType, nil
	case "LogTag":
		return logBody.LogTag, nil
	case "Message":
		return logBody.Message, nil
	case "LogDate":
		return logBody.LogDate, nil
	case "UserID":
		return logBody.UserID, nil

	default:
		return nil, funcs.CannotFindProperty(name, constant.ENTITY_LOG_BODY)
	}
}

func (logBody *LogBodySvcEnt) InitValidator() {
	validator1 := validator.LogBodyValidator{}
	validator1.Init(constant.ENTITY_LOG_BODY)
	validator1.SetValidatorObserve(logBody)
	logBody.objectValidator = &validator1
}

func (logBody *LogBodySvcEnt) MapFromEnt(logBody1 entity.LogBody) {
	logBody.ID = logBody1.ID
	logBody.SystemID = logBody1.SystemID
	logBody.AppID = logBody1.AppID
	logBody.LogID = logBody1.LogID
	logBody.LogType = logBody1.LogType
	logBody.LogTag = logBody1.LogTag
	logBody.Message = logBody1.Message
	logBody.LogDate = logBody1.LogDate
	logBody.UserID = logBody1.UserID
}

func (logBody LogBodySvcEnt) ToEnt() entity.LogBody {

	logBody1 := entity.LogBody{}
	logBody1.ID = logBody.ID
	logBody1.SystemID = logBody.SystemID
	logBody1.AppID = logBody.AppID
	logBody1.LogID = logBody.LogID
	logBody1.LogType = logBody.LogType
	logBody1.LogTag = logBody.LogTag
	logBody1.Message = logBody.Message
	logBody1.LogDate = logBody.LogDate
	logBody1.UserID = logBody.UserID

	return logBody1
}

func (logBody LogBodySvcEnt) ValidateBaseID() app.ErrorCollection {
	logBody.InitValidator()
	fields := []string{"ID"}
	return logBody.objectValidator.Validate(fields)
}

func (logBody LogBodySvcEnt) ValidateBase() app.ErrorCollection {
	logBody.InitValidator()
	fields := []string{"SystemID", "AppID", "LogID", "LogType", "LogTag", "Message", "LogDate", "UserID"}
	return logBody.objectValidator.Validate(fields)
}
