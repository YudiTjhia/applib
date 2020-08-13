package logger

import (
	"applib/constant"
	"applib/data/gorm"
	"applib/data/mongo"
	"applib/db"
)

func CreateLogHeaderData(dataLayerType string, connectionID string,
	dbConfigFile string, appConfigFile string) ILogHeaderData {

	var _idata ILogHeaderData

	if dataLayerType == constant.DATA_LAYER_SQL {
		_data := gorm.LogHeaderData{}
		_data.ConnectionID = connectionID
		_data.DbConfigFile = dbConfigFile
		_data.AppConfigFile = appConfigFile
		_idata = _data

	} else if dataLayerType == constant.DATA_LAYER_NOSQL {
		_data := mongo.LogHeaderData{}
		_data.ConnectionID = connectionID
		_data.DbConfigFile = dbConfigFile
		_data.AppConfigFile = appConfigFile
		_idata = _data
	} else {
		panic("Unknown dataLayerType=" + dataLayerType)
	}

	return _idata

}

func CreateLogHeaderData2(dataLayerType string, connectionID string,
	dbConfigFile string, appConfigFile string) ILogHeaderData2 {

	var _idata ILogHeaderData2

	if dataLayerType == constant.DATA_LAYER_SQL {
		_data := gorm.LogHeaderData2{}
		_data.ConnectionID = connectionID
		_data.DbConfigFile = dbConfigFile
		_data.AppConfigFile = appConfigFile
		_idata = _data

	} else if dataLayerType == constant.DATA_LAYER_NOSQL {
		_data := mongo.LogHeaderData2{}
		_data.ConnectionID = connectionID
		_data.DbConfigFile = dbConfigFile
		_data.AppConfigFile = appConfigFile
		_idata = _data
	} else {
		panic("Unknown dataLayerType=" + dataLayerType)
	}

	return _idata

}

func CreateLogBodyData(dataLayerType string, connectionID string,
	dbConfigFile string, appConfigFile string) ILogBodyData {

	var _idata ILogBodyData

	if dataLayerType == constant.DATA_LAYER_SQL {
		_data := gorm.LogBodyData{}
		_data.ConnectionID = connectionID
		_data.DbConfigFile = dbConfigFile
		_data.AppConfigFile = appConfigFile
		_idata = _data

	} else if dataLayerType == constant.DATA_LAYER_NOSQL {
		_data := mongo.LogBodyData{}
		_data.ConnectionID = connectionID
		_data.DbConfigFile = dbConfigFile
		_data.AppConfigFile = appConfigFile
		_idata = _data

	} else {
		panic("Unknown dataLayerType=" + dataLayerType)

	}

	return _idata

}

func CreateLogBodyData2(dataLayerType string, connectionID string,
	dbConfigFile string, appConfigFile string) ILogBodyData2 {

	var _idata ILogBodyData2

	if dataLayerType == constant.DATA_LAYER_SQL {
		_data := gorm.LogBodyData2{}
		_data.ConnectionID = connectionID
		_data.DbConfigFile = dbConfigFile
		_data.AppConfigFile = appConfigFile
		_idata = _data

	} else if dataLayerType == constant.DATA_LAYER_NOSQL {
		_data := mongo.LogBodyData2{}
		_data.ConnectionID = connectionID
		_data.DbConfigFile = dbConfigFile
		_data.AppConfigFile = appConfigFile
		_idata = _data

	} else {
		panic("Unknown dataLayerType=" + dataLayerType)

	}

	return _idata

}

func CreateLogBodyDataByConnection(connection db.DataConnection) ILogBodyData {

	var _idata ILogBodyData
	if connection.DataLayerType == constant.DATA_LAYER_SQL {
		_data := &gorm.LogBodyData{}
		_data.SetDataConnection(connection)
		_idata = _data

	} else if connection.DataLayerType == constant.DATA_LAYER_NOSQL {
		_data := &mongo.LogBodyData{}
		_data.SetDataConnection(connection)
		_idata = _data

	} else {
		panic("Unknown dataLayerType=" + connection.DataLayerType)
	}

	return _idata

}

func CreateLogBodyData2ByConnection(connection db.DataConnection) ILogBodyData2 {

	var _idata ILogBodyData2
	if connection.DataLayerType == constant.DATA_LAYER_SQL {
		_data := &gorm.LogBodyData2{}
		_data.SetDataConnection(connection)
		_idata = _data

	} else if connection.DataLayerType == constant.DATA_LAYER_NOSQL {
		_data := &mongo.LogBodyData2{}
		_data.SetDataConnection(connection)
		_idata = _data

	} else {
		panic("Unknown dataLayerType=" + connection.DataLayerType)
	}

	return _idata

}
