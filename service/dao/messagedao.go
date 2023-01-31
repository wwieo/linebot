package dao

import (
	"context"
	"linebot/config"
	"linebot/model"

	"go.mongodb.org/mongo-driver/mongo"
)

func Insert(client *mongo.Client, message *model.Message) error {
	mongoConfig := config.GetMongoConfig()
	coll := client.Database(mongoConfig.Database).Collection(mongoConfig.Collection)
	_, err := coll.InsertOne(context.TODO(), message)
	if err != nil {
		return err
	}
	return nil
}
