package message

import (
	"context"
	"linebot/database"
	"linebot/model"
	"linebot/service/dao/messagedao"
	"log"
)

type messageIntf interface {
	ReceiveMessage(ctx context.Context, message *model.Message) error
	GetMessages(ctx context.Context, userID string) ([]*model.Message, error)
}

type messageImpl struct {
}

func New() messageIntf {
	return &messageImpl{}
}

func (impl *messageImpl) ReceiveMessage(ctx context.Context, message *model.Message) error {
	db := database.GetMongoDB()
	err := messagedao.Insert(ctx, db, message)
	if err != nil {
		log.Println("Insert message error:", err)
		return err
	}

	return nil
}

func (impl *messageImpl) GetMessages(ctx context.Context, userID string) ([]*model.Message, error) {
	db := database.GetMongoDB()
	queryModel := &messagedao.QueryModel{
		UserID: userID,
	}
	messages, err := messagedao.Gets(ctx, db, queryModel)
	if err != nil {
		log.Println("Gets message error:", err)
		return nil, err
	}

	return messages, nil
}
