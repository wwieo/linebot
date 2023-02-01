package controller

import (
	"context"
	"fmt"
	"linebot/config"
	"linebot/model"
	"linebot/service/database/messagedao"
)

type MessageController struct {
}

func NewMessageController() *MessageController {
	return &MessageController{}
}

func (messageController *MessageController) ReceiveMessage(ctx context.Context, message *model.Message) error {
	client := config.GetMongoClient()
	err := messagedao.Insert(client, message)
	if err != nil {
		fmt.Println("Insert message error:", err)
		return err
	}

	return nil
}

func (messageController *MessageController) GetMessages(ctx context.Context, userID string) ([]*model.Message, error) {
	client := config.GetMongoClient()
	queryModel := &messagedao.Querymodel{
		UserID: userID,
	}
	fmt.Println("eeee")
	messages, err := messagedao.Gets(client, queryModel)

	if err != nil {
		fmt.Println("Gets message error:", err)
		return nil, err
	}

	return messages, nil
}
