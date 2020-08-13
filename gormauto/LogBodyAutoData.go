package gormauto

import (
	"applib/app"
	"applib/constant"
	"applib/db"
	"applib/entity"
	libfuncs "applib/funcs"
	"applib/logger"
	"github.com/google/uuid"
	gorm2 "github.com/jinzhu/gorm"
)

type LogBodyAutoData struct {
	db.GormData
}

func (data LogBodyAutoData) Table() *gorm2.DB {
	return data.Db.Table(constant.TABLE_LOG_BODY)
}
func (data LogBodyAutoData) Insert(header app.RequestHeader, logBody entity.LogBody) (string, error) {
	logBody.AccountID = header.RequestAccount
	logBody.ID = uuid.New().String()
	query := data.Table().Create(&logBody)
	log := logger.CreateBodyLogger(header)
	log.InsertQuery(header, data.LogID, query)
	if query.Error != nil {
		return "", query.Error
	}
	return logBody.ID, nil
}
func (data LogBodyAutoData) Update(header app.RequestHeader, logBody entity.LogBody) error {
	query := data.Table().Where("id=?", logBody.ID).Updates(logBody)

	log := logger.CreateBodyLogger(header)
	log.InsertQuery(header, data.LogID, query)
	return query.Error
}

func (data LogBodyAutoData) Delete(header app.RequestHeader, logBody entity.LogBody) error {
	query := data.Table().Where("account_id=?", header.RequestAccount)
	query = data.filterCondition(logBody, query).Delete(entity.LogBody{})

	log := logger.CreateBodyLogger(header)
	log.InsertQuery(header, data.LogID, query)
	return query.Error
}
func (data LogBodyAutoData) GetList(header app.RequestHeader, logBody entity.LogBody) ([]entity.LogBody, error) {
	logBodys := []entity.LogBody{}
	query := data.Table()
	query = data.preload(query)
	query = query.Where("account_id=?", header.RequestAccount)
	query = data.filterCondition(logBody, query).Order("log_date desc").Limit(100)
	query.Find(&logBodys)
	return logBodys, nil
}
func (data LogBodyAutoData) Count(header app.RequestHeader, logBody entity.LogBody) (int64, error) {
	logBodys := []entity.LogBody{}
	query := data.Table().Where("account_id=?", header.RequestAccount)
	query = data.filterCondition(logBody, query)
	cnt := int64(0)
	query.Find(&logBodys).Count(&cnt)
	return cnt, nil
}
func (data LogBodyAutoData) Search(header app.RequestHeader, logBody entity.LogBody, q string) ([]entity.LogBody, error) {
	logBodys := []entity.LogBody{}
	query := data.Table()
	query = data.preload(query)
	query = query.Where("account_id=?", header.RequestAccount)
	query = data.filterSearch(logBody, query).Order("log_date desc").Limit(100)
	query.Find(&logBodys)
	return logBodys, nil
}
func (data LogBodyAutoData) GetOne(header app.RequestHeader, logBody entity.LogBody) (entity.LogBody, error) {
	logBody1 := entity.LogBody{}
	query := data.Table()
	query = data.preload(query)
	query = query.Where("account_id=?", header.RequestAccount)
	query = data.filterCondition(logBody, query)
	query = query.Find(&logBody1)
	if query.RecordNotFound() {
		return logBody1, libfuncs.CannotFind(constant.ENTITY_LOG_BODY)
	}
	return logBody1, query.Error
}
func (data LogBodyAutoData) preload(query *gorm2.DB) *gorm2.DB {
	return query
}
func (data LogBodyAutoData) filterCondition(logBody entity.LogBody,
	query *gorm2.DB) *gorm2.DB {
	if logBody.ID != "" {
		query = query.Where("id=?", logBody.ID)
	}
	if logBody.SystemID != "" {
		query = query.Where("system_id=?", logBody.SystemID)
	}
	if logBody.AppID != "" {
		query = query.Where("app_id=?", logBody.AppID)
	}
	if logBody.LogID != "" {
		query = query.Where("log_id=?", logBody.LogID)
	}
	if logBody.LogType != "" {
		query = query.Where("log_type=?", logBody.LogType)
	}
	if logBody.LogTag != "" {
		query = query.Where("log_tag=?", logBody.LogTag)
	}
	if logBody.Message != "" {
		query = query.Where("message=?", logBody.Message)
	}
	if logBody.UserID != "" {
		query = query.Where("user_id=?", logBody.UserID)
	}

	return query
}
func (data LogBodyAutoData) filterSearch(logBody entity.LogBody,
	query *gorm2.DB) *gorm2.DB {
	if logBody.ID != "" {
		query = query.Where("id like ?", "%"+logBody.ID+"%")
	}
	if logBody.SystemID != "" {
		query = query.Where("system_id like ?", "%"+logBody.SystemID+"%")
	}
	if logBody.AppID != "" {
		query = query.Where("app_id like ?", "%"+logBody.AppID+"%")
	}
	if logBody.LogID != "" {
		query = query.Where("log_id like ?", "%"+logBody.LogID+"%")
	}
	if logBody.LogType != "" {
		query = query.Where("log_type like ?", "%"+logBody.LogType+"%")
	}
	if logBody.LogTag != "" {
		query = query.Where("log_tag like ?", "%"+logBody.LogTag+"%")
	}
	if logBody.Message != "" {
		query = query.Where("message like ?", "%"+logBody.Message+"%")
	}
	if logBody.UserID != "" {
		query = query.Where("user_id like ?", "%"+logBody.UserID+"%")
	}

	return query
}
