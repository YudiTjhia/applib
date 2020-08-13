package gormauto

import (
	"applib/app"
	"applib/constant"
	"applib/db"
	"applib/entity"
	libfuncs "applib/funcs"
	"github.com/google/uuid"
	gorm2 "github.com/jinzhu/gorm"
)

type LogResponseAutoData struct {
	db.GormData
}

func (data LogResponseAutoData) Table() *gorm2.DB {
	return data.Db.Table(constant.TABLE_LOG_RESPONSE)
}
func (data LogResponseAutoData) Insert(header app.RequestHeader, logResponse entity.LogResponse) (string, error) {
	logResponse.AccountID = header.RequestAccount
	logResponse.ID = uuid.New().String()
	query := data.Table().Create(&logResponse)
	if query.Error != nil {
		return "", query.Error
	}
	return logResponse.ID, nil
}
func (data LogResponseAutoData) Update(header app.RequestHeader, logResponse entity.LogResponse) error {
	query := data.Table().Where("id=?", logResponse.ID).Updates(logResponse)
	return query.Error
}
func (data LogResponseAutoData) Delete(header app.RequestHeader, logResponse entity.LogResponse) error {
	query := data.Table().Where("account_id=?", header.RequestAccount)
	query = data.filterCondition(logResponse, query).Delete(entity.LogResponse{})
	return query.Error
}
func (data LogResponseAutoData) GetList(header app.RequestHeader, logResponse entity.LogResponse) ([]entity.LogResponse, error) {
	logResponses := []entity.LogResponse{}
	query := data.Table()
	query = data.preload(query)
	query = query.Where("account_id=?", header.RequestAccount)
	query = data.filterCondition(logResponse, query)
	query.Find(&logResponses)
	return logResponses, nil
}
func (data LogResponseAutoData) Count(header app.RequestHeader, logResponse entity.LogResponse) (int64, error) {
	logResponses := []entity.LogResponse{}
	query := data.Table().Where("account_id=?", header.RequestAccount)
	query = data.filterCondition(logResponse, query)
	cnt := int64(0)
	query.Find(&logResponses).Count(&cnt)
	return cnt, nil
}
func (data LogResponseAutoData) Search(header app.RequestHeader, logResponse entity.LogResponse, q string) ([]entity.LogResponse, error) {
	logResponses := []entity.LogResponse{}
	query := data.Table()
	query = data.preload(query)
	query = query.Where("account_id=?", header.RequestAccount)
	query = data.filterSearch(logResponse, query)
	query.Find(&logResponses)
	return logResponses, nil
}
func (data LogResponseAutoData) GetOne(header app.RequestHeader, logResponse entity.LogResponse) (entity.LogResponse, error) {
	logResponse1 := entity.LogResponse{}
	query := data.Table()
	query = data.preload(query)
	query = query.Where("account_id=?", header.RequestAccount)
	query = data.filterCondition(logResponse, query)
	query = query.Find(&logResponse1)
	if query.RecordNotFound() {
		return logResponse1, libfuncs.CannotFind(constant.ENTITY_LOG_RESPONSE)
	}
	return logResponse1, query.Error
}
func (data LogResponseAutoData) preload(query *gorm2.DB) *gorm2.DB {
	return query
}
func (data LogResponseAutoData) filterCondition(logResponse entity.LogResponse,
	query *gorm2.DB) *gorm2.DB {
	if logResponse.ID != "" {
		query = query.Where("id=?", logResponse.ID)
	}
	if logResponse.SystemID != "" {
		query = query.Where("system_id=?", logResponse.SystemID)
	}
	if logResponse.AppID != "" {
		query = query.Where("app_id=?", logResponse.AppID)
	}
	if logResponse.ServiceName != "" {
		query = query.Where("service_name=?", logResponse.ServiceName)
	}
	return query
}
func (data LogResponseAutoData) filterSearch(logResponse entity.LogResponse,
	query *gorm2.DB) *gorm2.DB {
	if logResponse.ID != "" {
		query = query.Where("id like ?", "%"+logResponse.ID+"%")
	}
	if logResponse.SystemID != "" {
		query = query.Where("system_id like ?", "%"+logResponse.SystemID+"%")
	}
	if logResponse.AppID != "" {
		query = query.Where("app_id like ?", "%"+logResponse.AppID+"%")
	}
	if logResponse.ServiceName != "" {
		query = query.Where("service_name like ?", "%"+logResponse.ServiceName+"%")
	}
	return query
}
