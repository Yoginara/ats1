package db

import (
	"api/config"

	mongo "github.com/kamagasaki/go-utils/mongo"
)

var MongoString = config.Config("MONGOSTRCONNECT")

var DBATS = mongo.DBInfo{
	DBString: MongoString,
	DBName:   config.ATSYS,
}
