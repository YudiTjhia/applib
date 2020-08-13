package logger

import (
	"applib/app"
	"applib/conf"
	"applib/constant"
	"applib/util"
	"encoding/json"
	"strconv"
)

type LogSvc struct {
	app.BaseSvc
}

func CreateLogSvc(dataLayerType string, connectionID string, dbConfigFile string,
	appConfigFile string) LogSvc {

	log := LogSvc{}
	log.DataLayerType = dataLayerType
	log.ConnectionID = connectionID
	log.DbConfigFile = dbConfigFile
	log.AppConfigFile = appConfigFile

	return log
}

func (service LogSvc) validateAccountID(accountID string) error {
	return util.Required(accountID, constant.INVALID_ACCOUNT_ID)
}

func (service LogSvc) validateUserID(userID string) error {
	return util.Required(userID, constant.INVALID_USER_ID)
}

func (service LogSvc) validateAppID(appID string) error {
	return util.Required(appID, constant.INVALID_APP_ID)
}

func (service LogSvc) validateLogID(logID string) error {
	return util.Required(logID, constant.LOG_ID_IS_REQUIRED)
}

func (service LogSvc) validateSvcName(svcName string) error {
	return util.Required(svcName, constant.SERVICE_NAME_IS_REQUIRED)
}

func (service LogSvc) validateBaseAccount(accountID string, userID string, appID string) error {
	err := service.validateAccountID(accountID)
	if err != nil {
		return err
	}

	err = service.validateUserID(userID)
	if err != nil {
		return err
	}

	err = service.validateAppID(appID)
	if err != nil {
		return err
	}

	return nil

}

func (service LogSvc) validateBaseID(logID string) error {

	err := service.validateLogID(logID)
	if err != nil {
		return err
	}

	return nil

}

func (service LogSvc) Begin(accountID string, requestUser string,
	systemID string, requestApp string, serviceName string) (string, app.ErrorCollection) {

	appErr := app.ErrorCollection{}

	if !appErr.HasErrors() {

		iLogHeaderData := CreateLogHeaderData(service.DataLayerType, service.ConnectionID,
			service.DbConfigFile, service.AppConfigFile)

		dbConf := conf.GetDbConfig(service.DbConfigFile, service.ConnectionID)

		dbConf.Passwd = ""
		dbInfo := conf.SerializeDbConf(dbConf)
		logID, err := iLogHeaderData.Insert(accountID,
			requestUser, systemID, requestApp, dbInfo)

		appErr.Add(err)

		if !appErr.HasErrors() {

			logBodyData := CreateLogBodyData(service.DataLayerType,
				service.ConnectionID, service.DbConfigFile, service.AppConfigFile)

			logBodyData.Insert(accountID, requestUser,
				systemID, requestApp, logID, constant.LOG_TYPE_INFO,
				constant.LOG_TAG_BEGIN, serviceName)

			return logID, appErr

		} else {
			return "", appErr

		}

	} else {
		return "", appErr

	}
}

func (service LogSvc) Message(accountID string, userID string, systemID string, appID string,
	logID string, logTag string, message string) error {

	err := util.Required(message, "Message is required")
	if err != nil {
		return err
	}

	logBodyData := CreateLogBodyData(service.DataLayerType, service.ConnectionID,
		service.DbConfigFile, service.AppConfigFile)

	logBodyData.Insert(accountID, userID, systemID, appID, logID, constant.LOG_TYPE_INFO, logTag, message)

	return nil
}

func (service LogSvc) Count(accountID string, userID string, systemID string, appID string, logID string, count int) {

	logBodyData := CreateLogBodyData(service.DataLayerType, service.ConnectionID,
		service.DbConfigFile, service.AppConfigFile)

	logBodyData.Insert(accountID, userID, systemID, appID, logID, constant.LOG_TYPE_INFO,
		constant.LOG_TAG_COUNT, strconv.Itoa(count))

}

func (service LogSvc) Data(accountID string, userID string, systemID string, appID string, logID string, logTag string, pData interface{}) {

	logBodyData := CreateLogBodyData(service.DataLayerType, service.ConnectionID,
		service.DbConfigFile, service.AppConfigFile)

	b, _ := json.Marshal(pData)

	logBodyData.Insert(accountID, userID, systemID, appID, logID, constant.LOG_TYPE_INFO, logTag, string(b))

}

func (service LogSvc) Params(accountID string, requestUser string,
	systemID string, appID string, logID string, pData interface{}) {

	logBodyData := CreateLogBodyData(service.DataLayerType, service.ConnectionID,
		service.DbConfigFile, service.AppConfigFile)

	b, _ := json.Marshal(pData)

	logBodyData.Insert(accountID, requestUser, systemID,
		appID, logID, constant.LOG_TYPE_INFO, constant.LOG_TAG_PARAMS, string(b))

}

func (service LogSvc) LogReturn(accountID string, userID string, systemID string, appID string, logID string, message interface{}) {

	logBodyData := CreateLogBodyData(service.DataLayerType, service.ConnectionID,
		service.DbConfigFile, service.AppConfigFile)

	b, _ := json.Marshal(message)

	logBodyData.Insert(accountID, userID, systemID, appID, logID, constant.LOG_TYPE_INFO, constant.LOG_TAG_RETURN, string(b))
}

func (service LogSvc) Error(accountID string, userID string, systemID string, appID string, logID string, message interface{}) {

	logBodyData := CreateLogBodyData(service.DataLayerType, service.ConnectionID,
		service.DbConfigFile, service.AppConfigFile)

	logBodyData.InsertInterface(accountID, userID, systemID, appID, logID, constant.LOG_TYPE_ERROR, "", message)

}

func (service LogSvc) End(accountID string, userID string, systemID string, appID string, logID string) {

	logBodyData := CreateLogBodyData(service.DataLayerType, service.ConnectionID,
		service.DbConfigFile, service.AppConfigFile)

	logBodyData.Insert(accountID, userID, systemID, appID, logID, constant.LOG_TYPE_INFO, constant.LOG_TAG_END, "")

}
