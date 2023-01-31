package view

import (
	"context"
	"fmt"
	"linebot/config"
	"linebot/model"
	"linebot/service/controller"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.Data(200, "text/plain", []byte("Hello world!"))
	})

	router.POST("/callback", insertMessage())

	return router
}

func insertMessage() gin.HandlerFunc {
	botClient := config.GetLinebotClient()
	return func(ginContext *gin.Context) {
		events, err := botClient.ParseRequest(ginContext.Request)
		if err != nil {
			fmt.Println("line bot parse request error:", err)
			return
		}

		for _, event := range events {
			profile, err := botClient.GetProfile(event.Source.UserID).Do()
			if err != nil {
				fmt.Println(err)
			}
			message := &model.Message{
				UserID:    profile.UserID,
				UserName:  profile.DisplayName,
				Timestamp: event.Timestamp,
			}

			if event.Type == linebot.EventTypeMessage {
				switch userMessage := event.Message.(type) {
				case *linebot.TextMessage:
					message.Text = userMessage.Text
					message.MessageID = userMessage.ID
					messageController := controller.NewMessageController()
					fmt.Println(message)
					err = messageController.InsertMessage(context.TODO(), message)
					if err != nil {
						fmt.Println("InsertMessage error:", err)
					}
				}
			}
		}

	}

}
