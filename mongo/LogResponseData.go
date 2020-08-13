package mongo

import (
	"applib/app"
	"applib/constant"
	"applib/db"
	"applib/entity"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LogResponseData struct {
	db.MongoData
}

func (data LogResponseData) table() *mongo.Collection {
	return data.Db.Collection(constant.TABLE_LOG_RESPONSE)
}
func (data LogResponseData) Insert(header app.RequestHeader, pLogResponse entity.LogResponse) (string, error) {
	pLogResponse.ObjectID = primitive.NewObjectID()
	pLogResponse.AccountID = header.RequestAccount
	result, err := data.table().InsertOne(data.Ctx, pLogResponse)
	if err != nil {
		return "", err
	}
	id := result.InsertedID.(primitive.ObjectID)
	return id.Hex(), err
}
func (data LogResponseData) Delete(header app.RequestHeader, pLogResponse entity.LogResponse) error {
	_id, _ := primitive.ObjectIDFromHex(pLogResponse.ID)
	condition := bson.M{
		"_id":        _id,
		"account_id": header.RequestAccount,
	}
	_, err := data.table().DeleteOne(data.Ctx, condition)
	if err != nil {
		return err
	}
	return nil
}
func (data LogResponseData) GetList(header app.RequestHeader, pLogResponse entity.LogResponse) ([]entity.LogResponse, error) {
	ents := []entity.LogResponse{}

	condition := bson.M{
		"account_id": header.RequestAccount,
	}

	findOpt := options.Find()
	findOpt.SetLimit(50)
	sort := bson.M{"response_time": -1}
	findOpt.SetSort(sort)

	cursor, err := data.table().Find(data.Ctx, condition, findOpt)
	if err != nil {
		return ents, err
	}
	for cursor.Next(data.Ctx) {
		ent := entity.LogResponse{}
		err := cursor.Decode(&ent)
		if err != nil {
			return ents, err
		}
		ent.ID = ent.ObjectID.Hex()
		ents = append(ents, ent)
	}
	return ents, nil
}

func (data LogResponseData) Count(header app.RequestHeader, pLogResponse entity.LogResponse) (int64, error) {
	condition := bson.M{
		"account_id": header.RequestAccount,
	}
	return data.table().CountDocuments(data.Ctx, condition)

}
func (data LogResponseData) Search(header app.RequestHeader, pLogResponse entity.LogResponse, q string) ([]entity.LogResponse, error) {
	ents := []entity.LogResponse{}
	condition := bson.M{
		"account_id": header.RequestAccount,
	}
	cursor, err := data.table().Find(data.Ctx, condition)
	if err != nil {
		return ents, err
	}
	for cursor.Next(data.Ctx) {
		ent := entity.LogResponse{}
		err := cursor.Decode(&ent)
		if err != nil {
			return ents, err
		}
		ent.ID = ent.ObjectID.Hex()
		ents = append(ents, ent)

	}
	return ents, nil
}

func (data LogResponseData) GetOne(header app.RequestHeader, pLogResponse entity.LogResponse) (entity.LogResponse, error) {
	ent := entity.LogResponse{}
	_id, _ := primitive.ObjectIDFromHex(pLogResponse.ID)
	condition := bson.M{
		"_id":        _id,
		"account_id": header.RequestAccount,
	}
	result := data.table().FindOne(data.Ctx, condition)
	err := result.Decode(&ent)
	if err != nil {
		if err.Error() == constant.MONGO_NO_DOC {
			return ent, errors.New(constant.CANNOT_FIND_LOG_RESPONSE)
		}
		return ent, err
	}
	ent.ID = ent.ObjectID.Hex()
	return ent, nil
}
