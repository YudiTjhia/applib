package mongo

import (
	"applib/conf"
	"applib/constant"
	"applib/db"
	"applib/entity"
	"context"
	"encoding/json"
	gorm2 "github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
	"time"
)

type LogBodyData struct {
	db.MongoData
	DbConfigFile  string
	AppConfigFile string
}

func (data LogBodyData) table() *mongo.Collection {
	return data.Db.Collection(constant.TABLE_LOG_BODY)
}

func (data *LogBodyData) init() {
	data.CreateConnection(data.DbConfigFile, data.ConnectionID)
}

func (data LogBodyData) _insert(accountID string, userID string, systemID string, appID string,
	logID string, logType string, logTag string, message string) {

	if !conf.IsEnableLog(data.AppConfigFile) {
		return
	}

	data.init()
	defer data.Close()

	ent := entity.LogBody{}
	ent.ObjectID = primitive.NewObjectID()
	ent.AccountID = accountID
	ent.LogID = logID
	ent.SystemID = systemID
	ent.AppID = appID
	ent.Message = message
	ent.LogDate = time.Now().UTC()
	ent.LogType = logType
	ent.LogTag = logTag

	_, _ = data.table().InsertOne(context.Background(), ent)

}

func (data LogBodyData) Insert(accountID string, userID string, systemID string, appID string,
	logID string, logType string, logTag string, message string) {

	data._insert(accountID, userID, systemID, appID, logID, logType, logTag, message)

}

func (data LogBodyData) InsertCount(accountID string, userID string, systemID string, appID string,
	logID string, count int) {

	logMessage := strconv.Itoa(count)

	data._insert(accountID, userID, systemID, appID, logID,
		constant.LOG_TYPE_INFO, constant.LOG_TAG_COUNT, logMessage)

}

func (data LogBodyData) InsertInterface(accountID string, userID string, systemID string, appID string,
	logID string, logType string, logTag string, message interface{}) {

	msg, _ := json.Marshal(message)
	logMessage := string(msg)

	data._insert(accountID, userID, systemID, appID, logID, logType, logTag, logMessage)

}

func (data LogBodyData) InsertData(accountID string, userID string, systemID string, appID string,
	logID string, logTag string, message interface{}) {

	msg, _ := json.Marshal(message)
	logMessage := string(msg)
	data._insert(accountID, userID, systemID, appID,
		logID, constant.LOG_TYPE_INFO, logTag, logMessage)

}

func (data LogBodyData) InsertParams(accountID string, userID string, systemID string, appID string,
	logID string, message interface{}) {

	msg, _ := json.Marshal(message)
	logMessage := string(msg)

	data._insert(accountID, userID, systemID, appID,
		logID, constant.LOG_TYPE_INFO, constant.LOG_TAG_PARAMS, logMessage)

}

func (data LogBodyData) InsertQuery(accountID string, userID string, systemID string, appID string,
	logID string, query *gorm2.DB) {

}

func (data LogBodyData) GetList(accountID string, fromDate time.Time, toDate time.Time,
	logType string) ([]entity.LogBody, error) {

	ents := []entity.LogBody{}
	condition := bson.M{
		"account_id": accountID,
		"log_date": bson.M{
			"$gte": fromDate,
			"$lte": toDate,
		},
	}

	if logType != "" {
		condition["log_type"] = logType
	}

	cursor, err := data.table().Find(data.Ctx, condition)
	if err != nil {
		return ents, err
	}

	for cursor.Next(data.Ctx) {

		ent := entity.LogBody{}
		err := cursor.Decode(&ent)
		if err != nil {
			return ents, err
		}
		ent.ID = ent.ObjectID.Hex()
		ents = append(ents, ent)
	}

	return ents, nil

}
