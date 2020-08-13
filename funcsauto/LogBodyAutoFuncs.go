package funcsauto

import (
	"applib/app"
	"applib/constant"
	"applib/datafactory"
	"applib/db"
	"applib/entity"
	"applib/logger"
	"applib/svcent"
)

func InsertLogBody(logger logger.LogSvc2,
	logBody svcent.LogBodySvcEnt) (string, app.ErrorCollection) {

	appErr := app.ErrorCollection{}
	connection := db.CreateDataConnection2(logBody.ToHeader(), constant.LOG_BODY_DATA)
	iLogBodyData := datafactory.CreateLogBodyData(connection, logger.LogID)
	id, err := iLogBodyData.Insert(logBody.ToHeader(),
		logBody.ToEnt())

	defer connection.Close()
	appErr.Add(err)
	if !appErr.HasErrors() {
		return id, appErr
	}

	return "", appErr
}

func DeleteLogBody(logger logger.LogSvc2, logBody svcent.LogBodySvcEnt) app.ErrorCollection {
	appErr := app.ErrorCollection{}

	connection := db.CreateDataConnection2(logBody.ToHeader(), constant.LOG_BODY_DATA)
	iLogBodyData := datafactory.CreateLogBodyData(connection, logger.LogID)
	err := iLogBodyData.Delete(logBody.ToHeader(), logBody.ToEnt())

	defer connection.Close()
	appErr.Add(err)
	return appErr
}

func GetListLogBody(logBody svcent.LogBodySvcEnt) []svcent.LogBodySvcEnt {
	appErr := app.ErrorCollection{}
	var logBodys []svcent.LogBodySvcEnt

	connection := db.CreateDataConnection2(logBody.ToHeader(), constant.LOG_BODY_DATA)
	iLogBodyData := datafactory.CreateLogBodyData(connection, "")
	logBodys1, err := iLogBodyData.GetList(logBody.ToHeader(),
		logBody.ToEnt())

	defer connection.Close()
	appErr.Add(err)

	if !appErr.HasErrors() {
		logBodys = mapLogBodys(logBodys1)
	}
	return logBodys
}

func SearchLogBody(logBody svcent.LogBodySvcEnt,
	q string) ([]svcent.LogBodySvcEnt, app.ErrorCollection) {
	appErr := app.ErrorCollection{}
	var logBodys []svcent.LogBodySvcEnt

	connection := db.CreateDataConnection2(logBody.ToHeader(), constant.LOG_BODY_DATA)
	iLogBodyData := datafactory.CreateLogBodyData(connection, "")
	logBodys1, err := iLogBodyData.Search(logBody.ToHeader(),
		logBody.ToEnt(), q)

	defer connection.Close()
	appErr.Add(err)

	if !appErr.HasErrors() {

		logBodys = mapLogBodys(logBodys1)

	}

	return logBodys, appErr
}

func GetOneLogBody(logBody svcent.LogBodySvcEnt) (svcent.LogBodySvcEnt, app.ErrorCollection) {
	appErr := app.ErrorCollection{}
	logBody1 := svcent.LogBodySvcEnt{}

	connection := db.CreateDataConnection2(logBody.ToHeader(), constant.LOG_BODY_DATA)
	iLogBodyData := datafactory.CreateLogBodyData(connection, "")

	logBody2, err := iLogBodyData.GetOne(logBody.ToHeader(), logBody.ToEnt())

	defer connection.Close()
	appErr.Add(err)

	if !appErr.HasErrors() {
		logBody1.MapFromEnt(logBody2)
	}

	return logBody1, appErr
}

func mapLogBodys(logBodys []entity.LogBody) []svcent.LogBodySvcEnt {
	logBodys1 := []svcent.LogBodySvcEnt{}
	for i := 0; i < len(logBodys); i++ {
		logBody := svcent.LogBodySvcEnt{}
		logBody.MapFromEnt(logBodys[i])
		logBodys1 = append(logBodys1, logBody)
	}
	return logBodys1
}
