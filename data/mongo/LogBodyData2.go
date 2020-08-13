package mongo

import (
	"applib/app"
	"applib/conf"
	"applib/constant"
	"applib/db"
	"applib/entity"
	"context"
	"encoding/json"
	"fmt"
	gorm2 "github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
	"time"
)

type LogBodyData2 struct {
	db.MongoData
	DbConfigFile  string
	AppConfigFile string
}

func (data LogBodyData2) table() *mongo.Collection {
	return data.Db.Collection(constant.TABLE_LOG_BODY)
}

func (data *LogBodyData2) init() {
	data.CreateConnection(data.DbConfigFile, data.ConnectionID)
}

func (data LogBodyData2) _insert(header app.RequestHeader,
	logID string, logType string, logTag string, message string) {

	if !conf.IsEnableLog(data.AppConfigFile) {
		return
	}

	data.init()
	ent := entity.LogBody{}
	ent.ObjectID = primitive.NewObjectID()
	ent.AccountID = header.RequestAccount
	ent.LogID = logID
	ent.SystemID = header.RequestSystem
	ent.AppID = header.RequestApp
	ent.Message = message
	ent.LogDate = time.Now().UTC()
	ent.LogType = logType
	ent.LogTag = logTag

	_, _ = data.table().InsertOne(context.Background(), ent)
	defer data.Close()

}

func (data LogBodyData2) Insert(header app.RequestHeader,
	logID string, logType string, logTag string, message string) {

	data._insert(header, logID, logType, logTag, message)

}

func (data LogBodyData2) InsertCount(header app.RequestHeader,
	logID string, count int) {

	logMessage := strconv.Itoa(count)

	data._insert(header,
		logID,
		constant.LOG_TYPE_INFO, constant.LOG_TAG_COUNT, logMessage)

}

func (data LogBodyData2) InsertInterface(header app.RequestHeader,
	logID string, logType string, logTag string, message interface{}) {

	msg, _ := json.Marshal(message)
	logMessage := string(msg)

	data._insert(header,
		logID, logType, logTag, logMessage)

}

func (data LogBodyData2) InsertData(header app.RequestHeader,
	logID string, logTag string, message interface{}) {

	msg, _ := json.Marshal(message)
	logMessage := string(msg)
	data._insert(header,
		logID, constant.LOG_TYPE_INFO, logTag, logMessage)

}

func (data LogBodyData2) InsertParams(header app.RequestHeader,
	logID string, message interface{}) {

	msg, _ := json.Marshal(message)
	logMessage := string(msg)

	data._insert(header,
		logID, constant.LOG_TYPE_INFO, constant.LOG_TAG_PARAMS, logMessage)

}

func (data LogBodyData2) InsertQuery(header app.RequestHeader,
	logID string, query *gorm2.DB) {

	if !conf.IsEnableLog(data.AppConfigFile) {
		return
	}
	data.init()
	sql, ok := query.Get("sql")
	if ok {
		ent := entity.LogBody{}
		ent.AccountID = header.RequestAccount
		ent.ObjectID = primitive.NewObjectID()
		ent.LogID = logID
		ent.SystemID = header.RequestSystem
		ent.AppID = header.RequestApp
		ent.LogDate = time.Now().UTC()
		ent.Message = sql.(string)
		ent.LogType = constant.LOG_TYPE_INFO
		ent.LogTag = constant.LOG_TAG_SQL
		_, err := data.table().InsertOne(context.Background(), ent)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	sqlVars, ok := query.Get("sqlVars")
	if ok {
		bSqlVars, _ := json.Marshal(sqlVars)
		ent := entity.LogBody{}
		ent.AccountID = header.RequestAccount
		ent.ObjectID = primitive.NewObjectID()
		ent.LogID = logID
		ent.SystemID = header.RequestSystem
		ent.AppID = header.RequestApp
		ent.LogDate = time.Now().UTC()
		ent.Message = string(bSqlVars)
		ent.LogType = constant.LOG_TYPE_INFO
		ent.LogTag = constant.LOG_TAG_SQLVARS
		_, err := data.table().InsertOne(context.Background(), ent)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	defer data.Close()

}

func (data LogBodyData2) GetList(header app.RequestHeader,
	fromDate time.Time, toDate time.Time,
	logType string) ([]entity.LogBody, error) {

	ents := []entity.LogBody{}
	condition := bson.M{
		"account_id": header.RequestAccount,
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
