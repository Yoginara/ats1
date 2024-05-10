package db

import (
	"api/config"
	"api/internal/model"
	"errors"
	"github.com/kamagasaki/go-utils/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertDataMatakuliah(requestData model.Matakuliah) error {
	db := mongo.MongoConnect(DBATS)
	insertedID := mongo.InsertOneDoc(db, config.MatakuliahTB, requestData)
	if insertedID == nil {
		return errors.New("couldn't insert data")
	}
	return nil
}

func GetDataMatakuliah(filter bson.M) ([]model.Matakuliah, error) {
	db := mongo.MongoConnect(DBATS)
	var datas []model.Matakuliah
	err := mongo.GetAllDocByFilter[model.Matakuliah](db, config.MatakuliahTB, filter, &datas)
	if err != nil {
		return nil, err
	}
	return datas, nil
}
