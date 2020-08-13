package gorm

import (
	"applib/app"
	"applib/conf"
	"applib/constant"
	"applib/db"
	"applib/entity"
	"fmt"
	"github.com/google/uuid"
	gorm2 "github.com/jinzhu/gorm"
	"time"
)

type LogHeaderData2 struct {
	db.GormData
	DbConfigFile  string
	AppConfigFile string
}

func (data LogHeaderData2) table() *gorm2.DB {
	return data.Db.Table(constant.TABLE_LOG_HEADER)
}
func (data *LogHeaderData2) init() {
	data.CreateConnection(data.DbConfigFile, data.ConnectionID)
}
func (data LogHeaderData2) Insert(header app.RequestHeader, dbInfo string) (string,
	error) {
	if !conf.IsEnableLog(data.AppConfigFile) {
		return "", nil
	}
	data.init()
	ent := entity.LogHeader{}
	ent.AccountID = header.RequestAccount
	ent.LogID = uuid.New().String()
	ent.UserID = header.RequestUser
	ent.SystemID = header.RequestSystem
	ent.AppID = header.RequestApp
	ent.LogDate = time.Now().UTC()
	ent.DbInfo = dbInfo
	fmt.Println(ent)
	query := data.table().Create(&ent)
	err := query.Error
	defer data.Db.Close()
	if err != nil {
		return "", err
	}
	return ent.LogID, nil
}
func (data LogHeaderData2) GetList(header app.RequestHeader,
	fromDate time.Time, toDate time.Time) ([]entity.LogHeader, error) {
	var ents []entity.LogHeader
	if !conf.IsEnableLog(data.AppConfigFile) {
		return ents, nil
	}
	data.init()
	query := data.table().
		Where("account_id=?", header.RequestAccount).
		Where("log_date>=? and log_date<=?", fromDate, toDate).
		Find(&ents)
	err := query.Error
	defer data.Db.Close()
	if err != nil {
		return ents, err
	}
	return ents, nil
}
func (data LogHeaderData2) Search(header app.RequestHeader, q string) ([]entity.LogHeader, error) {
	var ents []entity.LogHeader
	if !conf.IsEnableLog(data.AppConfigFile) {
		return ents, nil
	}
	data.init()
	query := data.table().
		Preload("UpdateUser").
		Where("account_id=?", header.RequestAccount)
	if q != "" {
		likeQ := "%" + q + "%"
		query = query.Where("log_id like ?", likeQ)
	}
	query = query.Find(&ents)
	err := query.Error
	defer data.Db.Close()
	if err != nil {
		return ents, err
	}
	return ents, nil
}
