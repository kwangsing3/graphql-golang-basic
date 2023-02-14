package dbhandler

import (
	"context"

	"github.com/kwangsing3/graphql-golang-basic/graph/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DATABASE = "DEBUG"
var COLLECTION = "STOCK"

var DB, _ = 
NewDBHandler("mongodb+srv://username:password@cluster0.jm9ahx2.mongodb.net/test")

type DBHandler struct {
	client *mongo.Client
	db     *mongo.Database
	coll   *mongo.Collection
}

func NewDBHandler(srv string) (*DBHandler, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(srv))
	if err != nil {
		return nil, err
	}
	err = client.Connect(context.TODO())
	if err != nil {
		return nil, err
	}
	db := client.Database(DATABASE)
	coll := db.Collection(COLLECTION)
	return &DBHandler{client: client, db: db, coll: coll}, nil
}

func DisConnect() {
	DB.client.Disconnect(context.TODO())
}

func (r *DBHandler) InsertStock(stock model.NewStock) (*model.Stock, error) {
	res := model.Stock{
		Name: stock.Name,
		Code: stock.Code,
	}

	_, err := r.coll.InsertOne(context.TODO(), stock)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *DBHandler) GetStockByCode(code string) (*model.Stock, error) {
	var res model.Stock
	err := r.coll.FindOne(context.TODO(), bson.M{"code": code}).Decode(&res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *DBHandler) InsertRecord(code string, record model.DailyRecord) (bool, error) {
	filter := bson.M{"code": code}
	update := bson.M{
		"$push": bson.M{
			"HistoricalRecord": record,
		},
	}
	_, err := r.coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *DBHandler) UpdateRecord(code string, record model.NewRecord) error {
	filter := bson.M{"code": code}
	update := bson.M{
		"$set": bson.M{
			"HistoricalRecord": record,
		},
	}
	_, err := r.coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
