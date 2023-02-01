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

func setCollection(client *mongo.Client) *mongo.Collection {
	mongoConfig := config.GetMongoConfig()
	return client.Database(mongoConfig.Database).Collection(mongoConfig.Collection)
}

func Gets(ctx context.Context, client *mongo.Client, querymodel *Querymodel) ([]*model.Message, error) {
	coll := setCollection(client)
	cursor, err := coll.Find(ctx, querymodel)

	messages := []*model.Message{}
	err = cursor.All(ctx, &messages)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func Insert(ctx context.Context, client *mongo.Client, message *model.Message) error {
	coll := setCollection(client)
	_, err := coll.InsertOne(ctx, message)
	if err != nil {
		return err
	}
	return nil
}
