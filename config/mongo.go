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

func init() {
	config := utilsController.GetConfig()
	mongoConfig = &database.MongoConfig{
		Url:        config.GetString("mongo.url"),
		Password:   config.GetString("mongo.password"),
		Database:   config.GetString("mongo.database"),
		Collection: config.GetString("mongo.collection"),
		Port:       config.GetInt("mongo.port"),
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
