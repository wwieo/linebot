package view

import (
	"context"
	"fmt"
	"linebot/config"
	"linebot/model"
	"linebot/service/controller"
	"linebot/utils"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.Data(200, "text/plain", []byte("Hello world!"))
	})

	router.GET("/messages", getMessages())
	router.POST("/callback", receiveMessage())
	router.POST("/pushMessage", pushMessage())

	return router
}

func receiveMessage() gin.HandlerFunc {
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
					err = messageController.ReceiveMessage(context.TODO(), message)
					if err != nil {
						fmt.Println("InsertMessage error:", err)
						return
					}
				}
			}
		}

	}
}

func pushMessage() gin.HandlerFunc {
	botClient := config.GetLinebotClient()
	return func(context *gin.Context) {
		userID := context.PostForm("userID")
		text := context.PostForm("text")
		messages := []linebot.SendingMessage{linebot.NewTextMessage(text)}
		_, err := botClient.PushMessage(userID, messages...).Do()
		if err != nil {
			fmt.Println("push message error:", err)
		}
	}
}

func getMessages() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		userID := ginContext.Query("userID")
		messageController := controller.NewMessageController()
		utils := utils.NewUtils()

		messages, err := messageController.GetMessages(context.TODO(), userID)
		if err != nil {
			utils.ReturnAPIResult(ginContext, false, err)
			return
		}
		utils.ReturnAPIResult(ginContext, true, messages)
		return
	}
}
