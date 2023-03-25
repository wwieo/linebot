package messagedao

import (
	"context"
	"linebot/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type QueryModel struct {
	UserID string
}

const collection = "Messages"

func setCollection(db *mongo.Database) *mongo.Collection {
	return db.Collection(collection)
}

func Gets(ctx context.Context, db *mongo.Database, querymodel *QueryModel) ([]*model.Message, error) {
	coll := setCollection(db)
	cursor, err := coll.Find(ctx, querymodel)

	messages := []*model.Message{}
	err = cursor.All(ctx, &messages)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func Insert(ctx context.Context, db *mongo.Database, message *model.Message) error {
	coll := setCollection(db)
	_, err := coll.InsertOne(ctx, message)
	if err != nil {
		return err
	}
	return nil
}
