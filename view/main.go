package view

import (
	"context"
	"linebot/config"
	"linebot/model"
	"linebot/service/controller"
	"linebot/utils"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

var (
	messageController *controller.MessageController
	util              *utils.Utils
)

func init() {
	messageController = controller.NewMessageController()
	util = utils.NewUtils()
}

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
	return func(c *gin.Context) {
		events, err := botClient.ParseRequest(c.Request)
		if err != nil {
			log.Println("line bot parse request error:", err)
			return
		}

		for _, event := range events {
			profile, err := botClient.GetProfile(event.Source.UserID).Do()

			if err != nil {
				log.Println(err)
				continue
			}
			message := &model.Message{
				UserID:    profile.UserID,
				Username:  profile.DisplayName,
				Timestamp: &event.Timestamp,
			}

			if event.Type == linebot.EventTypeMessage {
				switch userMessage := event.Message.(type) {
				case *linebot.TextMessage:
					message.Text = userMessage.Text
					message.MessageID = userMessage.ID
					messageController := controller.NewMessageController()
					err = messageController.ReceiveMessage(context.TODO(), message)
					if err != nil {
						log.Println("InsertMessage error:", err)
						return
					}
				}
			}
		}

	}
}

func pushMessage() gin.HandlerFunc {
	botClient := config.GetLinebotClient()
	return func(c *gin.Context) {
		message := &model.Message{}
		c.ShouldBind(message)

		if message.UserID == "" {
			util.ReturnAPIResult(c, false, "userID shouldn't be none")
			return
		}

		messages := []linebot.SendingMessage{linebot.NewTextMessage(message.Text)}
		_, err := botClient.PushMessage(message.UserID, messages...).Do()
		if err != nil {
			log.Println("push message error:", err)
			return
		}
	}
}

func getMessages() gin.HandlerFunc {
	return func(c *gin.Context) {
		message := &model.Message{}
		c.ShouldBind(message)
		messages, err := messageController.GetMessages(context.TODO(), message.UserID)
		if err != nil {
			util.ReturnAPIResult(c, false, err)
			return
		}
		util.ReturnAPIResult(c, true, messages)
		return
	}
}
