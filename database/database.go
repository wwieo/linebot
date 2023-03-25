package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"linebot/utils"
)

type DB interface {
	initialize(context.Context, dbConfig)
	db() interface{}
}

const (
	mongoURL        = "mongo.url"
	mongoPassword   = "mongo.password"
	mongoDatabase   = "mongo.dao"
	mongoCollection = "mongo.collection"
	mongoPort       = "mongo.port"
)

type dbConfig struct {
	Url        string
	Password   string
	Database   string
	Collection string
	Port       int
}

var dbIntf DB

func Initialize(ctx context.Context) {
	config := utils.GetConfig()
	mongoConfig := dbConfig{
		Url:        config.GetString(mongoURL),
		Password:   config.GetString(mongoPassword),
		Database:   config.GetString(mongoDatabase),
		Collection: config.GetString(mongoCollection),
		Port:       config.GetInt(mongoPort),
	}

	dbIntf = &mongoDB{}
	dbIntf.initialize(ctx, mongoConfig)
}

func GetMongoDB() *mongo.Database {
	return dbIntf.db().(*mongo.Database)
}
