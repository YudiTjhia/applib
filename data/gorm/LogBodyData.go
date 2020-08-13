package gorm

import (
	"applib/conf"
	"applib/constant"
	"applib/db"
	"applib/entity"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	gorm2 "github.com/jinzhu/gorm"
	"strconv"
	"time"
)

type LogBodyData struct {
	db.GormData
	DbConfigFile  string
	AppConfigFile string
}

func (data LogBodyData) table() *gorm2.DB {
	return data.Db.Table(constant.TABLE_LOG_BODY)
}

func (data *LogBodyData) init() {
	data.CreateConnection(data.DbConfigFile, data.ConnectionID)
}

func (data LogBodyData) Insert(accountID string, userID string, systemID string, appID string,
	logID string, logType string, logTag string, message string) {

	if !conf.IsEnableLog(data.AppConfigFile) {
		return
	}

	data.init()
	defer data.Db.Close()

	ent := entity.LogBody{}
	ent.AccountID = accountID
	ent.ID = uuid.New().String()
	ent.LogID = logID
	ent.SystemID = systemID
	ent.AppID = appID
	ent.Message = message
	ent.LogDate = time.Now().UTC()
	ent.LogType = logType
	ent.LogTag = logTag

	data.table().Create(&ent)

	fmt.Println(ent)

}

func (data LogBodyData) InsertCount(accountID string, userID string, systemID string, appID string,
	logID string, count int) {

	if !conf.IsEnableLog(data.AppConfigFile) {
		return
	}

	data.init()
	defer data.Db.Close()

	ent := entity.LogBody{}
	ent.AccountID = accountID
	ent.ID = uuid.New().String()
	ent.LogID = logID
	ent.SystemID = systemID
	ent.AppID = appID
	ent.Message = strconv.Itoa(count)
	ent.LogDate = time.Now().UTC()
	ent.LogType = constant.LOG_TYPE_INFO
	ent.LogTag = constant.LOG_TAG_COUNT

	data.table().Create(&ent)

	fmt.Println(ent)

}

func (data LogBodyData) InsertInterface(accountID string, userID string, systemID string, appID string,
	logID string, logType string, logTag string, message interface{}) {

	if !conf.IsEnableLog(data.AppConfigFile) {
		return
	}

	data.init()
	defer data.Db.Close()

	ent := entity.LogBody{}
	ent.AccountID = accountID
	ent.ID = uuid.New().String()
	ent.LogID = logID
	ent.SystemID = systemID
	ent.AppID = appID
	msg, _ := json.Marshal(message)
	ent.Message = string(msg)

	ent.LogDate = time.Now().UTC()
	ent.LogType = logType
	ent.LogTag = logTag

	data.table().Create(&ent)

	fmt.Println(ent)

}

func (data LogBodyData) InsertData(accountID string, userID string, systemID string, appID string,
	logID string, logTag string, message interface{}) {

	if !conf.IsEnableLog(data.AppConfigFile) {
		return
	}

	data.init()
	defer data.Db.Close()

	ent := entity.LogBody{}
	ent.AccountID = accountID
	ent.ID = uuid.New().String()
	ent.LogID = logID
	ent.SystemID = systemID
	ent.AppID = appID

	msg, _ := json.Marshal(message)
	ent.Message = string(msg)

	ent.LogDate = time.Now().UTC()
	ent.LogType = constant.LOG_TYPE_INFO
	ent.LogTag = logTag

	data.table().Create(&ent)

	fmt.Println(ent)

}

func (data LogBodyData) InsertParams(accountID string, userID string, systemID string, appID string,
	logID string, message interface{}) {

	if !conf.IsEnableLog(data.AppConfigFile) {
		return
	}

	data.init()
	defer data.Db.Close()

	ent := entity.LogBody{}
	ent.AccountID = accountID
	ent.ID = uuid.New().String()
	ent.LogID = logID
	ent.SystemID = systemID
	ent.AppID = appID

	msg, _ := json.Marshal(message)
	ent.Message = string(msg)

	ent.LogDate = time.Now().UTC()
	ent.LogType = constant.LOG_TYPE_INFO
	ent.LogTag = constant.LOG_TAG_PARAMS

	data.table().Create(&ent)

	fmt.Println(ent)

}

func (data LogBodyData) InsertQuery(accountID string, userID string, systemID string, appID string,
	logID string, query *gorm2.DB) {

	if !conf.IsEnableLog(data.AppConfigFile) {
		return
	}

	data.init()
	defer data.Db.Close()

	sql, ok := query.Get("sql")
	if ok {
		ent := entity.LogBody{}
		ent.AccountID = accountID
		ent.ID = uuid.New().String()
		ent.LogID = logID
		ent.SystemID = systemID
		ent.AppID = appID
		ent.LogDate = time.Now().UTC()
		ent.Message = sql.(string)
		ent.LogType = constant.LOG_TYPE_INFO
		ent.LogTag = constant.LOG_TAG_SQL
		data.table().Create(&ent)

		fmt.Println(ent)

	}

	sqlVars, ok := query.Get("sqlVars")
	if ok {

		bSqlVars, _ := json.Marshal(sqlVars)

		ent := entity.LogBody{}
		ent.AccountID = accountID
		ent.ID = uuid.New().String()
		ent.LogID = logID
		ent.SystemID = systemID
		ent.AppID = appID
		ent.LogDate = time.Now().UTC()
		ent.Message = string(bSqlVars)
		ent.LogType = constant.LOG_TYPE_INFO
		ent.LogTag = constant.LOG_TAG_SQLVARS
		data.table().Create(&ent)

		fmt.Println(ent)

	}

}

func (data LogBodyData) GetList(accountID string, fromDate time.Time, toDate time.Time,
	logType string) ([]entity.LogBody, error) {

	ents := []entity.LogBody{}
	query := data.table().
		Where("account_id=?", accountID).
		Where("log_date >= ?", fromDate).
		Where("log_date <= ?", toDate)

	if logType != "" {
		query = query.Where("log_type=?", logType)
	}

	query = query.Order("log_date desc")

	query.Find(&ents)

	return ents, nil

}
