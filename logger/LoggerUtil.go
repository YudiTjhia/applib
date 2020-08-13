package logger

import (
	"applib/app"
	"applib/conf"
)

func CreateBodyLogger(header app.RequestHeader) ILogBodyData2 {
	connIDConf := conf.GetConnectionID(header.AppConfigFile.ConnIDConfigFile,
		LOG_BODY_DATA)

	log := CreateLogBodyData2(connIDConf.DataLayerType, connIDConf.ConnectionID,
		header.AppConfigFile.DBConfigFile,
		header.AppConfigFile.AppConfigFile)

	return log
}

func CreateSvcLogger(header app.RequestHeader) LogSvc2 {
	logConnID := conf.GetConnectionID(header.AppConfigFile.ConnIDConfigFile, LOG_SVC)
	log := CreateLogSvc2(logConnID.DataLayerType,
		logConnID.ConnectionID,
		header.AppConfigFile.DBConfigFile,
		header.AppConfigFile.AppConfigFile)
	return log
}
