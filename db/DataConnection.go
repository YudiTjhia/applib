package db

import (
	"applib/app"
	"applib/conf"
	libconst "applib/constant"
	"context"
	"errors"

	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/mongo"
)

type DataConnection struct {
	GormDB        *gorm.DB
	MongoDB       *mongo.Database
	MongoClient   *mongo.Client
	CApi          CApi
	DataLayerType string

	WithTransaction bool
	GormTransaction *gorm.DB
	MongoSession    mongo.Session
}

var (
	_connection *gorm.DB
)

func CreateDBConnection(header app.RequestHeader, connectionID string) (DataConnection, error) {

	connectionFile := header.AppConfigFile.ConnIDConfigFile
	configFile := header.AppConfigFile.DBConfigFile

	dataConn := DataConnection{}
	connectionIDConf, err := conf.GetConnectionID2(connectionFile, connectionID)
	if err!=nil { return dataConn, err }

	dataConn.DataLayerType = connectionIDConf.DataLayerType
	if connectionIDConf.DataLayerType == libconst.DATA_LAYER_SQL {
		if _connection == nil {
			_connection, err := CreateGormConnection2(configFile, connectionIDConf.ConnectionID)
			if err!=nil {
				return DataConnection{}, err
			}

			dataConn.GormDB = _connection
		} else {
			dataConn.GormDB = _connection
		}

	} else if connectionIDConf.DataLayerType == libconst.DATA_LAYER_NOSQL {
		dataConn.MongoClient, dataConn.MongoDB = CreateMongoConnection(configFile, connectionIDConf.ConnectionID)

	} else if connectionIDConf.DataLayerType == libconst.DATA_LAYER_API {
		dataConn.CApi = CreateCApiConnection(configFile, connectionIDConf.ConnectionID)
	} else {
		//panic("Invalid DataLayerType=" + connectionIDConf.DataLayerType)
		return DataConnection{}, errors.New("Invalid DataLayerType=" + connectionIDConf.DataLayerType)
	}

	return dataConn, nil
}

func CreateDataConnection3(appConfig app.AppConfigFile,	connectionID string) DataConnection {

	connectionFile := appConfig.ConnIDConfigFile
	configFile := appConfig.DBConfigFile

	dataConn := DataConnection{}
	connectionIDConf := conf.GetConnectionID(connectionFile, connectionID)
	dataConn.DataLayerType = connectionIDConf.DataLayerType

	if connectionIDConf.DataLayerType == libconst.DATA_LAYER_SQL {
		if _connection == nil {
			_connection = CreateGormConnection(configFile, connectionIDConf.ConnectionID)
			dataConn.GormDB = _connection
		} else {
			dataConn.GormDB = _connection
		}

	} else if connectionIDConf.DataLayerType == libconst.DATA_LAYER_NOSQL {
		dataConn.MongoClient, dataConn.MongoDB = CreateMongoConnection(configFile, connectionIDConf.ConnectionID)

	} else if connectionIDConf.DataLayerType == libconst.DATA_LAYER_API {
		dataConn.CApi = CreateCApiConnection(configFile, connectionIDConf.ConnectionID)
	} else {
		panic("Invalid DataLayerType=" + connectionIDConf.DataLayerType)
	}	
	return dataConn
}


func CreateDataConnection2(header app.RequestHeader, connectionID string) DataConnection {

	connectionFile := header.AppConfigFile.ConnIDConfigFile
	configFile := header.AppConfigFile.DBConfigFile

	dataConn := DataConnection{}
	connectionIDConf := conf.GetConnectionID(connectionFile, connectionID)
	dataConn.DataLayerType = connectionIDConf.DataLayerType

	if connectionIDConf.DataLayerType == libconst.DATA_LAYER_SQL {
		if _connection == nil {
			_connection = CreateGormConnection(configFile, connectionIDConf.ConnectionID)
			dataConn.GormDB = _connection
		} else {
			dataConn.GormDB = _connection
		}
		//dataConn.GormDB = CreateGormConnection(configFile, connectionIDConf.ConnectionID)

	} else if connectionIDConf.DataLayerType == libconst.DATA_LAYER_NOSQL {
		dataConn.MongoClient, dataConn.MongoDB = CreateMongoConnection(configFile, connectionIDConf.ConnectionID)

	} else if connectionIDConf.DataLayerType == libconst.DATA_LAYER_API {
		dataConn.CApi = CreateCApiConnection(configFile, connectionIDConf.ConnectionID)
	} else {
		panic("Invalid DataLayerType=" + connectionIDConf.DataLayerType)
	}

	return dataConn
}

func CreateDataConnection(connectionFile string, connectionID string,
	configFile string) DataConnection {

	dataConn := DataConnection{}
	connectionIDConf := conf.GetConnectionID(connectionFile, connectionID)
	dataConn.DataLayerType = connectionIDConf.DataLayerType

	if connectionIDConf.DataLayerType == libconst.DATA_LAYER_SQL {
		dataConn.GormDB = CreateGormConnection(configFile, connectionIDConf.ConnectionID)

	} else if connectionIDConf.DataLayerType == libconst.DATA_LAYER_NOSQL {
		dataConn.MongoClient, dataConn.MongoDB = CreateMongoConnection(configFile, connectionIDConf.ConnectionID)

	} else if connectionIDConf.DataLayerType == libconst.DATA_LAYER_API {
		dataConn.CApi = CreateCApiConnection(configFile, connectionIDConf.ConnectionID)
	} else {
		panic("Invalid DataLayerType=" + connectionIDConf.DataLayerType)
	}

	return dataConn

}

func (dataConn *DataConnection) BeginTransaction() {

	if dataConn.DataLayerType == libconst.DATA_LAYER_SQL {

		dataConn.GormTransaction = dataConn.GormDB.Begin()
		dataConn.WithTransaction = true

	} else if dataConn.DataLayerType == libconst.DATA_LAYER_NOSQL {

		session, err := dataConn.MongoClient.StartSession()
		if err != nil {
			return
		}

		dataConn.MongoSession = session
		err = dataConn.MongoSession.StartTransaction()
		if err != nil {
			return
		}
		dataConn.WithTransaction = true

	}
}

func (dataConn *DataConnection) Commit() {
	if dataConn.DataLayerType == libconst.DATA_LAYER_SQL {
		dataConn.WithTransaction = false
		dataConn.GormTransaction.Commit()

	} else if dataConn.DataLayerType == libconst.DATA_LAYER_NOSQL {
		dataConn.WithTransaction = false
		_ = dataConn.MongoSession.CommitTransaction(context.Background())
		dataConn.MongoSession.EndSession(context.Background())
	}
}

func (dataConn *DataConnection) Close() {

	if dataConn.DataLayerType == libconst.DATA_LAYER_SQL {
		if dataConn.WithTransaction {
			dataConn.GormTransaction.Close()
		}
		//dataConn.GormDB.Close()

	} else if dataConn.DataLayerType == libconst.DATA_LAYER_NOSQL {
		_ = dataConn.MongoClient.Disconnect(context.Background())

	}

}

func (dataConn *DataConnection) Rollback() {

	if dataConn.DataLayerType == libconst.DATA_LAYER_SQL {

		if dataConn.WithTransaction {
			dataConn.GormTransaction.Rollback()
		}
		dataConn.WithTransaction = false

	} else if dataConn.DataLayerType == libconst.DATA_LAYER_NOSQL {

		if dataConn.WithTransaction {
			_ = dataConn.MongoSession.AbortTransaction(context.Background())
			dataConn.MongoSession.EndSession(context.Background())
		}
		dataConn.WithTransaction = false

	}
}
