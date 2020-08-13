package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type LogHeader struct {
	ID        primitive.ObjectID `bson:"_id" json:"-" gorm:"-"`
	LogID     string             `gorm:"Column:log_id;Type:varchar;primary_key;not null" bson:"-" json:"logID"`
	AccountID string             `gorm:"Column:account_id;Type:varchar;not null" bson:"account_id" json:"accountID"`
	LogDate   time.Time          `gorm:"Column:log_date;Type:datetime;not null" bson:"log_date" json:"logDate"`
	UserID    string             `gorm:"Column:user_id;Type:varchar;not null" bson:"user_id" json:"userID"`
	SystemID  string             `gorm:"Column:system_id;Type:varchar;not null" bson:"system_id" json:"systemID"`
	AppID     string             `gorm:"Column:app_id;Type:varchar;not null" json:"appID" bson:"app_id"`
	DbInfo    string             `gorm:"Column:db_info;Type:text;not null" bson:"db_info" json:"dbInfo"`
}
