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

type LogResponseSvcEnt struct {
	app.BaseSvcEnt
	app.RequestHeader
	entity.LogResponse
	objectValidator libvalidator.IObjectValidator
}

func (logResponse *LogResponseSvcEnt) DecodeQuery(r *http.Request) {
	logResponse.ID = logResponse.DecodeQueryKey(r, "id")
	logResponse.SystemID = logResponse.DecodeQueryKey(r, "systemID")
	logResponse.AppID = logResponse.DecodeQueryKey(r, "appID")
	logResponse.ServiceName = logResponse.DecodeQueryKey(r, "serviceName")
}

func (logResponse LogResponseSvcEnt) GetValue(name string) (interface{}, error) {
	switch name {
	case "ID":
		return logResponse.ID, nil
	case "SystemID":
		return logResponse.SystemID, nil
	case "AppID":
		return logResponse.AppID, nil
	case "ServiceName":
		return logResponse.ServiceName, nil
	case "DataSize":
		return logResponse.DataSize, nil
	case "RequestTime":
		return logResponse.RequestTime, nil
	case "ResponseTime":
		return logResponse.ResponseTime, nil
	case "ResponseDuration":
		return logResponse.ResponseDuration, nil

	default:
		return nil, funcs.CannotFindProperty(name, constant.ENTITY_LOG_RESPONSE)
	}
}

func (logResponse *LogResponseSvcEnt) InitValidator() {
	validator1 := validator.LogResponseValidator{}
	validator1.Init(constant.ENTITY_LOG_RESPONSE)
	validator1.SetValidatorObserve(logResponse)
	logResponse.objectValidator = &validator1
}

func (logResponse *LogResponseSvcEnt) MapFromEnt(logResponse1 entity.LogResponse) {
	logResponse.ID = logResponse1.ID
	logResponse.SystemID = logResponse1.SystemID
	logResponse.AppID = logResponse1.AppID
	logResponse.ServiceName = logResponse1.ServiceName
	logResponse.DataSize = logResponse1.DataSize
	logResponse.RequestTime = logResponse1.RequestTime
	logResponse.ResponseTime = logResponse1.ResponseTime
	logResponse.ResponseDuration = logResponse1.ResponseDuration
}

func (logResponse LogResponseSvcEnt) ToEnt() entity.LogResponse {

	logResponse1 := entity.LogResponse{}
	logResponse1.ID = logResponse.ID
	logResponse1.SystemID = logResponse.SystemID
	logResponse1.AppID = logResponse.AppID
	logResponse1.ServiceName = logResponse.ServiceName
	logResponse1.DataSize = logResponse.DataSize
	logResponse1.RequestTime = logResponse.RequestTime
	logResponse1.ResponseTime = logResponse.ResponseTime
	logResponse1.ResponseDuration = logResponse.ResponseDuration

	return logResponse1
}

func (logResponse LogResponseSvcEnt) ValidateBaseID() app.ErrorCollection {
	logResponse.InitValidator()
	fields := []string{"ID"}
	return logResponse.objectValidator.Validate(fields)
}

func (logResponse LogResponseSvcEnt) ValidateBase() app.ErrorCollection {
	logResponse.InitValidator()
	fields := []string{"SystemID", "AppID", "ServiceName", "DataSize", "RequestTime", "ResponseTime", "ResponseDuration"}
	return logResponse.objectValidator.Validate(fields)
}
