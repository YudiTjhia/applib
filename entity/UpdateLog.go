package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type OrderBy struct {
	OrderField string `json:"orderField"`
	SortType   string `json:"sortType"`
}

type OrderByClause struct {
	OrderBys []OrderBy `json:"orderBys"`
}

type UpdateLog struct {
	AccountID  string    `gorm:"Column:account_id;Type:varchar;not null" bson:"account_id" json:"accountID"`
	UpdateDate time.Time `gorm:"Column:update_date;Type:datetime;not null" bson:"update_date" json:"updateDate"`
	UpdateBy   string    `gorm:"Column:update_by;Type:varchar;not null" bson:"update_by" json:"updateBy"`
	UpdateApp  string    `gorm:"Column:update_app;Type:varchar;not null" bson:"update_app" json:"updateApp"`

	UpdateByMaster  UserMasterBase `gorm:"ForeignKey:UpdateBy;ASSOCIATION_FOREIGNKEY:ID" json:"updateByMaster"`
	UpdateAppMaster AppMasterBase  `gorm:"ForeignKey:UpdateApp;ASSOCIATION_FOREIGNKEY:AppID" json:"updateAppMaster"`
}

type AppMasterBase struct {
	ObjectID primitive.ObjectID `bson:"_id" json:"-" gorm:"-"`
	AppID    string             `gorm:"Column:app_id;Type:varchar;primary_key;not null" bson:"app_id" json:"appID"`
	SystemID string             `gorm:"Column:system_id;Type:varchar;not null" bson:"system_id" json:"systemID"`
	AppName  string             `gorm:"Column:app_name;Type:varchar;null" bson:"app_name" json:"appName"`
	IsActive bool               `gorm:"Column:is_active;Type:bit;not null" bson:"is_active" json:"isActive"`
}

func (ent AppMasterBase) TableName() string {
	return "app_master"
}

type UserMasterBase struct {
	ObjectID    primitive.ObjectID `bson:"_id" json:"-" gorm:"-"`
	ID          string             `gorm:"Column:id;Type:varchar;primary_key;not null" bson:"id" json:"id"`
	UserID      string             `gorm:"Column:user_id;Type:varchar;not null" bson:"user_id" json:"userID"`
	MobilePhone string             `gorm:"Column:mobile_phone;Type:varchar;not null" bson:"mobile_phone" json:"mobilePhone"`
	UserName    string             `gorm:"Column:user_name;Type:varchar;not null" bson:"user_name" json:"userName"`
	IsActive    bool               `gorm:"Column:is_active;Type:bit;not null" bson:"is_active" json:"isActive"`
	IsSuperUser bool               `gorm:"Column:is_super_user;Type:bit;not null" bson:"is_super_user" json:"isSuperUser"`
}

func (ent UserMasterBase) TableName() string {
	return "user_master"
}
