package db

import (
	"applib/conf"
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoData struct {
	Ctx      context.Context
	DbClient *mongo.Client
	Db       *mongo.Database

	LogID        string
	ConnectionID string
}

func (data *MongoData) SetDataConnection(connection DataConnection) {
	data.Db = connection.MongoDB
	data.DbClient = connection.MongoClient
	data.Ctx = context.TODO()

}

func (data *MongoData) Close() {
	if data.DbClient != nil {
		err := data.DbClient.Disconnect(context.Background())
		fmt.Println(err)
	}
}

func (data *MongoData) SetConnectionID(connectionID string) {
	data.ConnectionID = connectionID
}

func (data *MongoData) CreateConnection(dbConfigFile string, connectionID string) {

	dbConfs := conf.GetDbConfigs(dbConfigFile)
	dbConf, e := dbConfs[connectionID]
	if !e {
		panic(fmt.Sprintf("Cannot find connectionID=%s on DbConfs.", connectionID))
	}

	connString := createConnString(dbConf)
	data.DbClient, data.Db = _createConnection(dbConf, connString)
}

func (data MongoData) GetGormDB() *gorm.DB {
	return nil
}

func (data MongoData) GetMongoDB() *mongo.Database {
	return data.Db
}

func (data *MongoData) SetLogID(logID string) {
	data.LogID = logID
}

func createConnString(conf conf.DbConf) string {
	//uri := "mongodb://{user}:{password}@{host}:{port}/{db}"
	//uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s",
	//	conf.User, conf.Passwd,  conf.Server, conf.Port, conf.DbName)

	uri := fmt.Sprintf("mongodb://%s:%d",
		conf.Server, conf.Port)
	return uri
}

func _createConnection(conf conf.DbConf, connString string) (*mongo.Client, *mongo.Database) {

	//client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	client, err := mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	db := client.Database(conf.DbName)

	return client, db
}

func CreateMongoConnection(dbConfigFile string, connectionID string) (*mongo.Client, *mongo.Database) {

	dbConfs := conf.GetDbConfigs(dbConfigFile)
	dbConf, e := dbConfs[connectionID]
	if !e {
		panic(fmt.Sprintf("Cannot find connectionID=%s on DbConfs.", connectionID))
	}

	connString := createConnString(dbConf)
	return _createConnection(dbConf, connString)
}
