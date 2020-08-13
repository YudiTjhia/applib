package logger

import (
	"applib/app"
	"applib/conf"
	"applib/constant"
	"applib/util"
	"encoding/json"
	"strconv"
)

type LogSvc2 struct {
	app.BaseSvc
	LogID string
}

func CreateLogSvc2(dataLayerType string, connectionID string, dbConfigFile string,
	appConfigFile string) LogSvc2 {
	log := LogSvc2{}
	log.DataLayerType = dataLayerType
	log.ConnectionID = connectionID
	log.DbConfigFile = dbConfigFile
	log.AppConfigFile = appConfigFile
	return log
}

func (logSvc LogSvc2) validateAccountID(accountID string) error {
	return util.Required(accountID, constant.INVALID_ACCOUNT_ID)
}

func (logSvc LogSvc2) validateUserID(userID string) error {
	return util.Required(userID, constant.INVALID_USER_ID)
}

func (logSvc LogSvc2) validateAppID(appID string) error {
	return util.Required(appID, constant.INVALID_APP_ID)
}

func (logSvc LogSvc2) validateLogID(logID string) error {
	return util.Required(logID, constant.LOG_ID_IS_REQUIRED)
}

func (logSvc LogSvc2) validateSvcName(svcName string) error {
	return util.Required(svcName, constant.SERVICE_NAME_IS_REQUIRED)
}

func (logSvc LogSvc2) validateBaseAccount(accountID string, userID string, appID string) error {
	return nil
}

func (logSvc LogSvc2) validateBaseID(logID string) error {
	return nil
}

func (logSvc *LogSvc2) Begin(header app.RequestHeader, serviceName string) (string, app.ErrorCollection) {
	appErr := app.ErrorCollection{}
	if !appErr.HasErrors() {
		iLogHeaderData := CreateLogHeaderData2(logSvc.DataLayerType, logSvc.ConnectionID,
			logSvc.DbConfigFile, logSvc.AppConfigFile)
		dbConf := conf.GetDbConfig(logSvc.DbConfigFile, logSvc.ConnectionID)
		dbConf.Passwd = ""
		dbInfo := conf.SerializeDbConf(dbConf)
		logID, err := iLogHeaderData.Insert(header, dbInfo)
		appErr.Add(err)
		if !appErr.HasErrors() {
			logSvc.LogID = logID
			logBodyData := CreateLogBodyData2(logSvc.DataLayerType,
				logSvc.ConnectionID, logSvc.DbConfigFile, logSvc.AppConfigFile)
			logBodyData.Insert(header, logID, constant.LOG_TYPE_INFO,
				constant.LOG_TAG_BEGIN, serviceName)
			return logID, appErr
		} else {
			return "", appErr
		}
	} else {
		return "", appErr
	}
}

func (logSvc LogSvc2) Message(header app.RequestHeader,
	logID string, logTag string, message string) error {

	err := util.Required(message, "Message is required")
	if err != nil {
		return err
	}

	logBodyData := CreateLogBodyData2(logSvc.DataLayerType, logSvc.ConnectionID,
		logSvc.DbConfigFile, logSvc.AppConfigFile)

	logBodyData.Insert(header, logID, constant.LOG_TYPE_INFO, logTag, message)

	return nil
}

func (logSvc LogSvc2) Count(header app.RequestHeader,
	logID string, count int) {

	logBodyData := CreateLogBodyData2(logSvc.DataLayerType, logSvc.ConnectionID,
		logSvc.DbConfigFile, logSvc.AppConfigFile)

	logBodyData.Insert(header, logID, constant.LOG_TYPE_INFO,
		constant.LOG_TAG_COUNT, strconv.Itoa(count))

}

func (logSvc LogSvc2) Data(header app.RequestHeader, logID string, logTag string, pData interface{}) {

	logBodyData := CreateLogBodyData2(logSvc.DataLayerType, logSvc.ConnectionID,
		logSvc.DbConfigFile, logSvc.AppConfigFile)
	b, _ := json.Marshal(pData)
	logBodyData.Insert(header, logID, constant.LOG_TYPE_INFO, logTag, string(b))
}

func (logSvc LogSvc2) Params(header app.RequestHeader, logID string, pData interface{}) {

	logBodyData := CreateLogBodyData2(logSvc.DataLayerType, logSvc.ConnectionID,
		logSvc.DbConfigFile, logSvc.AppConfigFile)

	b, _ := json.Marshal(pData)

	logBodyData.Insert(header, logID, constant.LOG_TYPE_INFO, constant.LOG_TAG_PARAMS, string(b))

}

func (logSvc LogSvc2) LogReturn(header app.RequestHeader, logID string, message interface{}) {

	logBodyData := CreateLogBodyData2(logSvc.DataLayerType, logSvc.ConnectionID,
		logSvc.DbConfigFile, logSvc.AppConfigFile)
	b, _ := json.Marshal(message)
	logBodyData.Insert(header, logID, constant.LOG_TYPE_INFO, constant.LOG_TAG_RETURN, string(b))
}

func (logSvc LogSvc2) Error(header app.RequestHeader, logID string, message interface{}) {

	logBodyData := CreateLogBodyData2(logSvc.DataLayerType, logSvc.ConnectionID,
		logSvc.DbConfigFile, logSvc.AppConfigFile)

	logBodyData.InsertInterface(header, logID, constant.LOG_TYPE_ERROR, "", message)

}

func (logSvc LogSvc2) End(header app.RequestHeader, logID string) {

	logBodyData := CreateLogBodyData2(logSvc.DataLayerType, logSvc.ConnectionID,
		logSvc.DbConfigFile, logSvc.AppConfigFile)

	logBodyData.Insert(header, logID, constant.LOG_TYPE_INFO, constant.LOG_TAG_END, "")

}
