package datafactory

import (
	libconst "applib/constant"
	"applib/db"
	"applib/gorm"
	"applib/idata"
	"applib/mongo"
)

func CreateLogResponseData(connection db.DataConnection, logID string) idata.ILogResponseData {
	var _idata idata.ILogResponseData
	if connection.DataLayerType == libconst.DATA_LAYER_SQL {
		_data := &gorm.LogResponseData{}
		_data.SetDataConnection(connection)
		_data.SetLogID(logID)
		_idata = _data
	} else if connection.DataLayerType == libconst.DATA_LAYER_NOSQL {
		_data := &mongo.LogResponseData{}
		_data.SetDataConnection(connection)
		_data.SetLogID(logID)
		_idata = _data
	} else {
		panic("Unknown dataLayerType=" + connection.DataLayerType)
	}
	return _idata
}
