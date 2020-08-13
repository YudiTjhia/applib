package funcsauto

import (
	"applib/app"
	"applib/constant"
	"applib/datafactory"
	"applib/db"
	"applib/entity"
	"applib/svcent"
)

func InsertLogResponse(logResponse svcent.LogResponseSvcEnt) (string, app.ErrorCollection) {

	appErr := app.ErrorCollection{}
	connection := db.CreateDataConnection2(logResponse.ToHeader(), constant.LOG_RESPONSE_DATA)
	iLogResponseData := datafactory.CreateLogResponseData(connection, "")
	id, err := iLogResponseData.Insert(logResponse.ToHeader(),
		logResponse.ToEnt())

	defer connection.Close()
	appErr.Add(err)
	if !appErr.HasErrors() {
		return id, appErr
	}

	return "", appErr
}

func DeleteLogResponse(logResponse svcent.LogResponseSvcEnt) app.ErrorCollection {
	appErr := app.ErrorCollection{}
	connection := db.CreateDataConnection2(logResponse.ToHeader(), constant.LOG_RESPONSE_DATA)
	iLogResponseData := datafactory.CreateLogResponseData(connection, "")
	err := iLogResponseData.Delete(logResponse.ToHeader(), logResponse.ToEnt())
	defer connection.Close()
	appErr.Add(err)
	return appErr
}

func GetListLogResponse(logResponse svcent.LogResponseSvcEnt) []svcent.LogResponseSvcEnt {
	appErr := app.ErrorCollection{}
	var logResponses []svcent.LogResponseSvcEnt
	connection := db.CreateDataConnection2(logResponse.ToHeader(), constant.LOG_RESPONSE_DATA)
	iLogResponseData := datafactory.CreateLogResponseData(connection, "")
	logResponses1, err := iLogResponseData.GetList(logResponse.ToHeader(),
		logResponse.ToEnt())

	defer connection.Close()
	appErr.Add(err)

	if !appErr.HasErrors() {
		logResponses = mapLogResponses(logResponses1)
	}
	return logResponses
}

func SearchLogResponse(logResponse svcent.LogResponseSvcEnt,
	q string) ([]svcent.LogResponseSvcEnt, app.ErrorCollection) {
	appErr := app.ErrorCollection{}
	var logResponses []svcent.LogResponseSvcEnt

	connection := db.CreateDataConnection2(logResponse.ToHeader(), constant.LOG_RESPONSE_DATA)
	iLogResponseData := datafactory.CreateLogResponseData(connection, "")
	logResponses1, err := iLogResponseData.Search(logResponse.ToHeader(),
		logResponse.ToEnt(), q)

	defer connection.Close()
	appErr.Add(err)

	if !appErr.HasErrors() {

		logResponses = mapLogResponses(logResponses1)

	}

	return logResponses, appErr
}

func GetOneLogResponse(logResponse svcent.LogResponseSvcEnt) (svcent.LogResponseSvcEnt, app.ErrorCollection) {
	appErr := app.ErrorCollection{}
	logResponse1 := svcent.LogResponseSvcEnt{}

	connection := db.CreateDataConnection2(logResponse.ToHeader(), constant.LOG_RESPONSE_DATA)
	iLogResponseData := datafactory.CreateLogResponseData(connection, "")

	logResponse2, err := iLogResponseData.GetOne(logResponse.ToHeader(), logResponse.ToEnt())

	defer connection.Close()
	appErr.Add(err)

	if !appErr.HasErrors() {
		logResponse1.MapFromEnt(logResponse2)
	}

	return logResponse1, appErr
}

func mapLogResponses(logResponses []entity.LogResponse) []svcent.LogResponseSvcEnt {
	logResponses1 := []svcent.LogResponseSvcEnt{}
	for i := 0; i < len(logResponses); i++ {
		logResponse := svcent.LogResponseSvcEnt{}
		logResponse.MapFromEnt(logResponses[i])
		logResponses1 = append(logResponses1, logResponse)
	}
	return logResponses1
}
