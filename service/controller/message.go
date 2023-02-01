package controller

import (
	"context"
	"linebot/config"
	"linebot/model"
	"linebot/service/database/messagedao"
	"log"
)

type MessageController struct {
}

func NewMessageController() *MessageController {
	return &MessageController{}
}

func (messageController *MessageController) ReceiveMessage(ctx context.Context, message *model.Message) error {
	client := config.GetMongoClient()
	err := messagedao.Insert(ctx, client, message)
	if err != nil {
		log.Println("Insert message error:", err)
		return err
	}

	return nil
}

func (messageController *MessageController) GetMessages(ctx context.Context, userID string) ([]*model.Message, error) {
	client := config.GetMongoClient()
	queryModel := &messagedao.Querymodel{
		UserID: userID,
	}
	messages, err := messagedao.Gets(ctx, client, queryModel)
	if err != nil {
		log.Println("Gets message error:", err)
		return nil, err
	}

	return messages, nil
}
