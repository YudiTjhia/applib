package gorm

import (
	"applib/conf"
	"applib/constant"
	"applib/db"
	"applib/entity"
	"fmt"
	"github.com/google/uuid"
	gorm2 "github.com/jinzhu/gorm"
	"time"
)

type LogHeaderData struct {
	db.GormData
	DbConfigFile  string
	AppConfigFile string
}

func (data LogHeaderData) table() *gorm2.DB {
	return data.Db.Table(constant.TABLE_LOG_HEADER)
}

func (data *LogHeaderData) init() {
	data.CreateConnection(data.DbConfigFile, data.ConnectionID)
}

func (data LogHeaderData) Insert(accountID string, userID string, systemID string,
	appID string, dbInfo string) (string,
	error) {

	if !conf.IsEnableLog(data.AppConfigFile) {
		return "", nil
	}
	data.init()
	defer data.Db.Close()

	ent := entity.LogHeader{}
	ent.AccountID = accountID
	ent.LogID = uuid.New().String()
	ent.UserID = userID
	ent.SystemID = systemID
	ent.AppID = appID
	ent.LogDate = time.Now().UTC()
	ent.DbInfo = dbInfo

	fmt.Println(ent)

	query := data.table().Create(&ent)
	err := query.Error
	if err != nil {
		return "", err
	}

	return ent.LogID, nil
}

func (data LogHeaderData) GetList(accountID string, systemID string,
	fromDate time.Time, toDate time.Time) ([]entity.LogHeader, error) {

	var ents []entity.LogHeader
	if !conf.IsEnableLog(data.AppConfigFile) {
		return ents, nil
	}

	data.init()
	defer data.Db.Close()

	query := data.table().
		Where("account_id=?", accountID).
		Where("log_date>=? and log_date<=?", fromDate, toDate).
		Find(&ents)

	err := query.Error
	if err != nil {
		return ents, err
	}

	return ents, nil

}

func (data LogHeaderData) Search(accountID string, q string) ([]entity.LogHeader, error) {

	var ents []entity.LogHeader

	if !conf.IsEnableLog(data.AppConfigFile) {
		return ents, nil
	}
	data.init()
	defer data.Db.Close()

	query := data.table().
		Preload("UpdateUser").
		Where("account_id=?", accountID)

	if q != "" {
		likeQ := "%" + q + "%"
		query = query.Where("log_id like ?", likeQ)
	}

	query = query.Find(&ents)
	err := query.Error
	if err != nil {
		return ents, err
	}
	return ents, nil

}
