package mongo

import (
	"applib/app"
	"applib/conf"
	"applib/constant"
	"applib/db"
	"applib/entity"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type LogHeaderData2 struct {
	db.MongoData
	DbConfigFile  string
	AppConfigFile string
}

func (data LogHeaderData2) table() *mongo.Collection {
	return data.Db.Collection(constant.TABLE_LOG_HEADER)
}
func (data *LogHeaderData2) init() {
	data.CreateConnection(data.DbConfigFile, data.ConnectionID)
}
func (data LogHeaderData2) Insert(
	header app.RequestHeader,
	dbInfo string) (string,
	error) {

	if !conf.IsEnableLog(data.AppConfigFile) {
		return "", nil
	}
	data.init()
	ent := entity.LogHeader{}
	ent.ID = primitive.NewObjectID()
	ent.AccountID = header.RequestAccount
	ent.UserID = header.RequestUser
	ent.SystemID = header.RequestSystem
	ent.AppID = header.RequestApp
	ent.LogDate = time.Now().UTC()
	ent.DbInfo = dbInfo

	result, err := data.table().InsertOne(context.Background(), ent)
	if err != nil {
		return "", err
	}
	logID := result.InsertedID.(primitive.ObjectID)
	defer data.Close()
	return logID.Hex(), nil
}

func (data LogHeaderData2) GetList(header app.RequestHeader,
	fromDate time.Time, toDate time.Time) ([]entity.LogHeader, error) {

	var ents []entity.LogHeader
	if !conf.IsEnableLog(data.AppConfigFile) {
		return ents, nil
	}

	return ents, nil

}

func (data LogHeaderData2) Search(header app.RequestHeader, q string) ([]entity.LogHeader, error) {

	var ents []entity.LogHeader
	if !conf.IsEnableLog(data.AppConfigFile) {
		return ents, nil
	}

	return ents, nil
}
