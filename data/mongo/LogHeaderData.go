package mongo

import (
	"applib/conf"
	"applib/constant"
	"applib/db"
	"applib/entity"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type LogHeaderData struct {
	db.MongoData
	DbConfigFile  string
	AppConfigFile string
}

func (data LogHeaderData) table() *mongo.Collection {
	return data.Db.Collection(constant.TABLE_LOG_HEADER)
}

func (data *LogHeaderData) init() {
	data.CreateConnection(data.DbConfigFile, data.ConnectionID)
}

func (data LogHeaderData) Insert(accountID string, userID string, systemID string, appID string,
	dbInfo string) (string,
	error) {

	if !conf.IsEnableLog(data.AppConfigFile) {
		return "", nil
	}
	data.init()
	defer data.Close()

	ent := entity.LogHeader{}
	ent.ID = primitive.NewObjectID()
	ent.AccountID = accountID
	ent.UserID = userID
	ent.SystemID = systemID
	ent.AppID = appID
	ent.LogDate = time.Now().UTC()
	ent.DbInfo = dbInfo

	result, err := data.table().InsertOne(context.Background(), ent)
	if err != nil {
		return "", err
	}
	logID := result.InsertedID.(primitive.ObjectID)
	return logID.Hex(), nil

}

func (data LogHeaderData) GetList(accountID string, systemID string,
	fromDate time.Time, toDate time.Time) ([]entity.LogHeader, error) {

	var ents []entity.LogHeader
	if !conf.IsEnableLog(data.AppConfigFile) {
		return ents, nil
	}

	return ents, nil

}

func (data LogHeaderData) Search(accountID string, q string) ([]entity.LogHeader, error) {

	var ents []entity.LogHeader
	if !conf.IsEnableLog(data.AppConfigFile) {
		return ents, nil
	}

	return ents, nil
}
