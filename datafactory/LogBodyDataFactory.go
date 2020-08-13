package datafactory

import (
	libconst "applib/constant"
	"applib/db"
	"applib/gorm"
	"applib/idata"
	"applib/mongo"
)

func CreateLogBodyData(connection db.DataConnection, logID string) idata.ILogBodyData {
	var _idata idata.ILogBodyData
	if connection.DataLayerType == libconst.DATA_LAYER_SQL {
		_data := &gorm.LogBodyData{}
		_data.SetDataConnection(connection)
		_data.SetLogID(logID)
		_idata = _data
	} else if connection.DataLayerType == libconst.DATA_LAYER_NOSQL {
		_data := &mongo.LogBodyData{}
		_data.SetDataConnection(connection)
		_data.SetLogID(logID)
		_idata = _data
	} else {
		panic("Unknown dataLayerType=" + connection.DataLayerType)
	}
	return _idata
}
