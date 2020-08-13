package db

import (
	"applib/conf"
	"applib/constant"
	"errors"
	"fmt"
	salestrackingconst "salestracking/constant"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"go.mongodb.org/mongo-driver/mongo"
)

type GormData struct {
	Db           *gorm.DB
	LogID        string
	ConnectionID string
	DataConnection DataConnection
}

func (data *GormData) SetDataConnection(connection DataConnection) {
	if connection.WithTransaction {
		data.Db = connection.GormTransaction
	} else {
		data.Db = connection.GormDB
	}
	data.DataConnection = connection
}


func (data *GormData) SetConnectionID(connectionID string) {
	data.ConnectionID = connectionID
}

func (data *GormData) CreateConnection(dbConfigFile string, connectionID string) {
	data.Db = CreateGormConnection(dbConfigFile, connectionID)
}

func (data GormData) GetGormDB() *gorm.DB {
	return data.Db
}

func (data GormData) GetMongoDB() *mongo.Database {
	return nil
}

func (data *GormData) SetLogID(logID string) {
	data.LogID = logID
}

func createGormConnString(dbConf conf.DbConf) string {

	if dbConf.DbType == constant.DB_TYPE_MSSQL {

		connString := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", dbConf.User, dbConf.Passwd,
			dbConf.Server, dbConf.Port, dbConf.DbName)

		return connString

	} else if dbConf.DbType == constant.DB_TYPE_MYSQL {

		connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			dbConf.User, dbConf.Passwd, dbConf.Server, dbConf.Port, dbConf.DbName)

		return connString

	} else if dbConf.DbType == constant.DB_TYPE_POSTGRES {
		connString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
			dbConf.Server, dbConf.Port, dbConf.User, dbConf.DbName, dbConf.Passwd)

		return connString

	} else {
		panic("Unknown DbType=" + dbConf.DbType)
	}

}


func createGormConnString2(dbConf conf.DbConf) (string, error) {

	if dbConf.DbType == constant.DB_TYPE_MSSQL {

		connString := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", dbConf.User, dbConf.Passwd,
			dbConf.Server, dbConf.Port, dbConf.DbName)

		return connString, nil

	} else if dbConf.DbType == constant.DB_TYPE_MYSQL {

		connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			dbConf.User, dbConf.Passwd, dbConf.Server, dbConf.Port, dbConf.DbName)

		return connString, nil

	} else if dbConf.DbType == constant.DB_TYPE_POSTGRES {
		connString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
			dbConf.Server, dbConf.Port, dbConf.User, dbConf.DbName, dbConf.Passwd)

		return connString, nil

	} else {
		//panic("Unknown DbType=" + dbConf.DbType)
		return "", errors.New("Unknown DbType=" + dbConf.DbType)
	}

}

func scopeStoreSqlAndVars(scope *gorm.Scope) {
	scope.DB().InstantSet("sql", scope.SQL)
	scope.DB().InstantSet("sqlVars", scope.SQLVars)
}

func registerCallback(connection *gorm.DB) {
	//connection.Callback().Query().Register("query_scope_sql_vars", scopeStoreSqlAndVars)
	//connection.Callback().Create().Register("query_scope_sql_vars", scopeStoreSqlAndVars)
	//connection.Callback().Update().Register("query_scope_sql_vars", scopeStoreSqlAndVars)
	//connection.Callback().Delete().Register("query_scope_sql_vars", scopeStoreSqlAndVars)

}


func _createGormConnection(dbConf conf.DbConf, connString string) *gorm.DB {

	dialect := ""

	if dbConf.DbType == constant.DB_TYPE_MSSQL {
		dialect = "mssql"

	} else if dbConf.DbType == constant.DB_TYPE_MYSQL {
		dialect = "mysql"

	} else if dbConf.DbType == constant.DB_TYPE_POSTGRES {
		dialect = "postgres"

	} else {
		panic("Unknown DbType=" + dbConf.DbType)
	}


	_connection, err := gorm.Open(dialect, connString)
	if err != nil {
		panic(err.Error())
	}
	if salestrackingconst.IS_DEBUG {
		_connection.LogMode(true)
		//registerCallback(_connection)
	}
	_connection.SingularTable(true)

	return _connection

}


func _createGormConnection2(dbConf conf.DbConf, connString string) (*gorm.DB, error) {

	dialect := ""
	if dbConf.DbType == constant.DB_TYPE_MSSQL {
		dialect = "mssql"

	} else if dbConf.DbType == constant.DB_TYPE_MYSQL {
		dialect = "mysql"

	} else if dbConf.DbType == constant.DB_TYPE_POSTGRES {
		dialect = "postgres"

	} else {
		panic("Unknown DbType=" + dbConf.DbType)
	}

	_connection, err := gorm.Open(dialect, connString)
	if err != nil {
		//panic(err.Error())
		return nil, err
	}

	if salestrackingconst.IS_DEBUG {
		_connection.LogMode(true)
		//registerCallback(_connection)
	}

	_connection.SingularTable(true)

	return _connection, nil

}


func CreateGormConnection(configFile string, connectionID string) *gorm.DB {

	dbConfs := conf.GetDbConfigs(configFile)
	dbConf, e := dbConfs[connectionID]
	if !e {
		panic(fmt.Sprintf("Cannot find connectionID=%s on DbConfs.", connectionID))
	}

	connString := createGormConnString(dbConf)
	connection := _createGormConnection(dbConf, connString)

	return connection
}


func CreateGormConnection2(configFile string, connectionID string) (*gorm.DB, error) {

	dbConfs,err := conf.GetDbConfigs2(configFile)
	if err!=nil {
		return nil, err
	}

	dbConf, e := dbConfs[connectionID]
	if !e {
		//panic(fmt.Sprintf("Cannot find connectionID=%s on DbConfs.", connectionID))
		return nil, errors.New(fmt.Sprintf("Cannot find connectionID=%s on DbConfs.", connectionID))
	}

	connString, err := createGormConnString2(dbConf)
	if err!=nil {
		return nil, err
	}

	connection, err := _createGormConnection2(dbConf, connString)
	if err!=nil {
		return nil, err
	}

	return connection, nil
}


func CreateGormConnectionAndConfig(configFile string, connectionID string) (*gorm.DB, conf.DbConf) {

	dbConfs := conf.GetDbConfigs(configFile)
	dbConf, e := dbConfs[connectionID]
	if !e {
		panic(fmt.Sprintf("Cannot find connectionID=%s on DbConfs.", connectionID))
	}

	connString := createGormConnString(dbConf)
	connection := _createGormConnection(dbConf, connString)

	return connection, dbConf

}
