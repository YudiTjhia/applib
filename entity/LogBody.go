package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type LogBody struct {
	ObjectID  primitive.ObjectID `bson:"_id" json:"-" gorm:"-"`
	ID        string             `gorm:"Column:id;Type:varchar;primary_key;not null" json:"-" bson:"-"`
	AccountID string             `gorm:"Column:account_id;Type:varchar;not null" json:"accountID" bson:"account_id"`
	SystemID  string             `gorm:"Column:system_id;Type:varchar;not null" json:"systemID" bson:"system_id"`
	AppID     string             `gorm:"Column:app_id;Type:varchar;not null" json:"appID" bson:"app_id"`
	LogID     string             `gorm:"Column:log_id;Type:varchar;not null" json:"logID" bson:"log_id"`
	LogTag    string             `gorm:"Column:log_tag;Type:varchar;not null" json:"logTag" bson:"log_tag"`
	Message   string             `gorm:"Column:message;Type:text;not null" json:"message" bson:"message"`
	LogDate   time.Time          `gorm:"Column:log_date;Type:datetime;not null" json:"logDate" bson:"log_date"`
	LogType   string             `gorm:"Column:log_type;Type:varchar;not null" json:"logType" bson:"log_type"`
	UserID    string             `gorm:"Column:user_id;Type:datetime;not null" json:"userID" bson:"user_id"`
}
