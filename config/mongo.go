package config

import (
	"context"
	"fmt"
	"linebot/model/database"
	"linebot/utils"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	utilsController = utils.NewUtils()
	mongoConfig     = &database.MongoConfig{}
)

const (
	mongoURL        = "mongo.url"
	mongoPassword   = "mongo.password"
	mongoDatabase   = "mongo.database"
	mongoCollection = "mongo.collection"
	mongoPort       = "mongo.port"
)

func init() {
	config := utilsController.GetConfig()
	mongoConfig = &database.MongoConfig{
		Url:        config.GetString(mongoURL),
		Password:   config.GetString(mongoPassword),
		Database:   config.GetString(mongoDatabase),
		Collection: config.GetString(mongoCollection),
		Port:       config.GetInt(mongoPort),
	}
}

func GetMongoConfig() *database.MongoConfig {
	return mongoConfig
}

func GetMongoClient() *mongo.Client {
	URI := fmt.Sprintf("%s:%d", mongoConfig.Url, mongoConfig.Port)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatal(err)
	}

	return client
}
