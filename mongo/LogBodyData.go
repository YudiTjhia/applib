package mongo

import (
	"applib/app"
	"applib/conf"
	"applib/constant"
	libconst "applib/constant"
	"applib/db"
	"applib/entity"
	"applib/logger"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LogBodyData struct {
	db.MongoData
}

func (data LogBodyData) table() *mongo.Collection {
	return data.Db.Collection(constant.TABLE_LOG_BODY)
}
func (data LogBodyData) createBodyLogger() logger.ILogBodyData2 {
	connIDConf := conf.GetConnectionID(constant.CONN_ID_CONFIG_FILE, constant.LOG_BODY_DATA)
	log := logger.CreateLogBodyData2(connIDConf.DataLayerType, connIDConf.ConnectionID,
		constant.DB_CONFIG_FILE, constant.APP_CONFIG_FILE)
	return log
}
func (data LogBodyData) Insert(header app.RequestHeader, pLogBodyData entity.LogBody) (string, error) {
	pLogBodyData.ObjectID = primitive.NewObjectID()
	pLogBodyData.AccountID = header.RequestAccount
	result, err := data.table().InsertOne(data.Ctx, pLogBodyData)
	if err != nil {
		return "", err
	}
	id := result.InsertedID.(primitive.ObjectID)
	log := data.createBodyLogger()
	log.InsertData(header,
		data.LogID,
		libconst.LOG_TAG_SQL_INSERT, pLogBodyData)
	return id.Hex(), err
}
func (data LogBodyData) Delete(header app.RequestHeader, pLogBody entity.LogBody) error {
	_id, _ := primitive.ObjectIDFromHex(pLogBody.ID)
	condition := bson.M{
		"_id":        _id,
		"account_id": header.RequestAccount,
	}
	_, err := data.table().DeleteOne(data.Ctx, condition)
	if err != nil {
		return err
	}
	log := data.createBodyLogger()
	log.InsertData(header, data.LogID, libconst.LOG_TAG_SQL_DELETE, condition)
	log.InsertData(header, data.LogID, libconst.LOG_TAG_SQLVARS, condition)
	return nil
}
func (data LogBodyData) GetList(header app.RequestHeader, pLogBody entity.LogBody) ([]entity.LogBody, error) {
	ents := []entity.LogBody{}

	condition := bson.M{
		"account_id": header.RequestAccount,
		"system_id":  pLogBody.SystemID,
	}

	findOpt := options.Find()
	findOpt.SetLimit(50)
	sort := bson.M{"log_date": -1}
	findOpt.SetSort(sort)

	cursor, err := data.table().Find(data.Ctx, condition, findOpt)

	if err != nil {
		return ents, err
	}
	for cursor.Next(data.Ctx) {
		ent := entity.LogBody{}
		err := cursor.Decode(&ent)
		if err != nil {
			return ents, err
		}
		ent.ID = ent.ObjectID.Hex()
		ents = append(ents, ent)
	}
	return ents, nil
}
func (data LogBodyData) Count(header app.RequestHeader, pLogBody entity.LogBody) (int64, error) {
	condition := bson.M{
		"account_id": header.RequestAccount,
	}
	return data.table().CountDocuments(data.Ctx, condition)
}
func (data LogBodyData) Search(header app.RequestHeader, pLogBody entity.LogBody, q string) ([]entity.LogBody, error) {
	ents := []entity.LogBody{}
	condition := bson.M{
		"account_id": header.RequestAccount,
	}
	cursor, err := data.table().Find(data.Ctx, condition)
	if err != nil {
		return ents, err
	}
	for cursor.Next(data.Ctx) {
		ent := entity.LogBody{}
		err := cursor.Decode(&ent)
		if err != nil {
			return ents, err
		}
		ent.ID = ent.ObjectID.Hex()
		ents = append(ents, ent)
	}
	return ents, nil
}
func (data LogBodyData) GetOne(header app.RequestHeader, pLogBody entity.LogBody) (entity.LogBody, error) {
	ent := entity.LogBody{}
	_id, _ := primitive.ObjectIDFromHex(pLogBody.ID)
	condition := bson.M{
		"_id":        _id,
		"account_id": header.RequestAccount,
	}
	result := data.table().FindOne(data.Ctx, condition)
	err := result.Decode(&ent)
	if err != nil {
		if err.Error() == constant.MONGO_NO_DOC {
			return ent, errors.New(constant.CANNOT_FIND_LOG_BODY)
		}
		return ent, err
	}
	ent.ID = ent.ObjectID.Hex()
	return ent, nil

}
