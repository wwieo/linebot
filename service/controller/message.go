package controller

import (
	"context"
	"fmt"
	"linebot/config"
	"linebot/model"
	"linebot/service/dao"
)

type MessageController struct {
}

func NewMessageController() *MessageController {
	return &MessageController{}
}

func (messageController *MessageController) InsertMessage(context context.Context, message *model.Message) error {
	client := config.GetMongoClient()
	err := dao.Insert(client, message)
	if err != nil {
		fmt.Println("insert message error:", err)
		return err
	}

	return nil
}
