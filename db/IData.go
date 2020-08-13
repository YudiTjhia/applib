package db

import (
	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/mongo"
)

type IData interface {
	GetGormDB() *gorm.DB
	GetMongoDB() *mongo.Database
	SetLogID(logID string)
}
