package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type LogResponse struct {
	ObjectID         primitive.ObjectID `bson:"_id" json:"-" gorm:"-"`
	ID               string             `gorm:"Column:id;Type:varchar;primary_key;not null" bson:"-" json:"id"`
	AccountID        string             `gorm:"Column:account_id;Type:varchar;not null" bson:"account_id" json:"accountID"`
	SystemID         string             `gorm:"Column:system_id;Type:varchar;not null" bson:"system_id" json:"systemID"`
	AppID            string             `gorm:"Column:app_id;Type:varchar;not null" bson:"app_id" json:"appID"`
	ServiceName      string             `gorm:"Column:service_name;Type:varchar;not null" bson:"service_name" json:"serviceName"`
	DataSize         int                `gorm:"Column:data_size;Type:int;not null" bson:"data_size" json:"dataSize"`
	RequestTime      time.Time          `gorm:"Column:request_time;Type:datetime;not null" bson:"request_time" json:"requestTime"`
	ResponseTime     time.Time          `gorm:"Column:response_time;Type:datetime;not null" bson:"response_time" json:"responseTime"`
	ResponseDuration float64            `gorm:"Column:response_duration;Type:float;not null" bson:"response_duration" json:"responseDuration"`
}
