package messagedao

import (
	"context"
	"linebot/config"
	"linebot/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type Querymodel struct {
	UserID string
}

func Gets(client *mongo.Client, querymodel *Querymodel) ([]*model.Message, error) {
	ctx := context.TODO()
	mongoConfig := config.GetMongoConfig()
	coll := client.Database(mongoConfig.Database).Collection(mongoConfig.Collection)
	cursor, err := coll.Find(ctx, querymodel)

	messages := []*model.Message{}
	err = cursor.All(ctx, &messages)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func Insert(client *mongo.Client, message *model.Message) error {
	mongoConfig := config.GetMongoConfig()
	coll := client.Database(mongoConfig.Database).Collection(mongoConfig.Collection)
	_, err := coll.InsertOne(context.TODO(), message)
	if err != nil {
		return err
	}
	return nil
}
