package gorm

import (
	"applib/app"
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

type LogBodyData2 struct {
	db.GormData
	DbConfigFile  string
	AppConfigFile string
}

func (data LogBodyData2) table() *gorm2.DB {
	return data.Db.Table(constant.TABLE_LOG_BODY)
}
func (data *LogBodyData2) init() {
	data.CreateConnection(data.DbConfigFile, data.ConnectionID)
}
func (data LogBodyData2) Insert(header app.RequestHeader,
	logID string, logType string, logTag string, message string) {
	if !conf.IsEnableLog(data.AppConfigFile) {
		return
	}
	data.init()
	ent := entity.LogBody{}
	ent.AccountID = header.RequestAccount
	ent.ID = uuid.New().String()
	ent.LogID = logID
	ent.SystemID = header.RequestSystem
	ent.AppID = header.RequestApp
	ent.Message = message
	ent.LogDate = time.Now().UTC()
	ent.LogType = logType
	ent.LogTag = logTag
	data.table().Create(&ent)
	defer data.Db.Close()
	fmt.Println(ent)
}
func (data LogBodyData2) InsertCount(header app.RequestHeader,
	logID string, count int) {
	if !conf.IsEnableLog(data.AppConfigFile) {
		return
	}
	data.init()
	ent := entity.LogBody{}
	ent.AccountID = header.RequestAccount
	ent.ID = uuid.New().String()
	ent.LogID = logID
	ent.SystemID = header.RequestSystem
	ent.AppID = header.RequestApp
	ent.Message = strconv.Itoa(count)
	ent.LogDate = time.Now().UTC()
	ent.LogType = constant.LOG_TYPE_INFO
	ent.LogTag = constant.LOG_TAG_COUNT
	data.table().Create(&ent)
	defer data.Db.Close()
	fmt.Println(ent)
}
func (data LogBodyData2) InsertInterface(header app.RequestHeader,
	logID string, logType string, logTag string, message interface{}) {
	if !conf.IsEnableLog(data.AppConfigFile) {
		return
	}
	data.init()
	ent := entity.LogBody{}
	ent.AccountID = header.RequestAccount
	ent.ID = uuid.New().String()
	ent.LogID = logID
	ent.SystemID = header.RequestSystem
	ent.AppID = header.RequestApp
	msg, _ := json.Marshal(message)
	ent.Message = string(msg)
	ent.LogDate = time.Now().UTC()
	ent.LogType = logType
	ent.LogTag = logTag
	data.table().Create(&ent)
	defer data.Db.Close()
	fmt.Println(ent)
}
func (data LogBodyData2) InsertData(header app.RequestHeader,
	logID string, logTag string, message interface{}) {
	if !conf.IsEnableLog(data.AppConfigFile) {
		return
	}
	data.init()
	ent := entity.LogBody{}
	ent.AccountID = header.RequestAccount
	ent.ID = uuid.New().String()
	ent.LogID = logID
	ent.SystemID = header.RequestSystem
	ent.AppID = header.RequestApp
	msg, _ := json.Marshal(message)
	ent.Message = string(msg)
	ent.LogDate = time.Now().UTC()
	ent.LogType = constant.LOG_TYPE_INFO
	ent.LogTag = logTag
	data.table().Create(&ent)
	defer data.Db.Close()
	fmt.Println(ent)
}
func (data LogBodyData2) InsertParams(header app.RequestHeader,
	logID string, message interface{}) {
	if !conf.IsEnableLog(data.AppConfigFile) {
		return
	}
	data.init()
	ent := entity.LogBody{}
	ent.AccountID = header.RequestAccount
	ent.ID = uuid.New().String()
	ent.LogID = logID
	ent.SystemID = header.RequestSystem
	ent.AppID = header.RequestApp
	msg, _ := json.Marshal(message)
	ent.Message = string(msg)
	ent.LogDate = time.Now().UTC()
	ent.LogType = constant.LOG_TYPE_INFO
	ent.LogTag = constant.LOG_TAG_PARAMS
	data.table().Create(&ent)
	defer data.Db.Close()
	fmt.Println(ent)
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
		ent.ID = uuid.New().String()
		ent.LogID = logID
		ent.SystemID = header.RequestSystem
		ent.AppID = header.RequestApp
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
		ent.AccountID = header.RequestAccount
		ent.ID = uuid.New().String()
		ent.LogID = logID
		ent.SystemID = header.RequestSystem
		ent.AppID = header.RequestApp
		ent.LogDate = time.Now().UTC()
		ent.Message = string(bSqlVars)
		ent.LogType = constant.LOG_TYPE_INFO
		ent.LogTag = constant.LOG_TAG_SQLVARS
		data.table().Create(&ent)
		fmt.Println(ent)
	}
	defer data.Db.Close()
}
func (data LogBodyData2) GetList(header app.RequestHeader,
	fromDate time.Time, toDate time.Time,
	logType string) ([]entity.LogBody, error) {
	ents := []entity.LogBody{}
	query := data.table().
		Where("account_id=?", header.RequestAccount).
		Where("log_date >= ?", fromDate).
		Where("log_date <= ?", toDate)
	if logType != "" {
		query = query.Where("log_type=?", logType)
	}
	query = query.Order("log_date desc")
	query.Find(&ents)
	defer data.Db.Close()
	return ents, nil
}
