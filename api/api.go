package api

import (
	"context"
	"linebot/bot"
	"linebot/model"
	"linebot/service/message"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.Data(200, "text/plain", []byte("Hello world!"))
	})

	router.GET("/messages", getMessages)
	router.POST("/pushMessage", pushMessage)
	router.POST("/callback", receiveMessage)

	return router
}

func getMessages(ctx *gin.Context) {
	m := &model.Message{}
	if err := ctx.ShouldBind(m); err != nil {
		ctx.JSON(http.StatusBadRequest, "getMessages bind error: "+err.Error())
		return
	}
	messageIntf := message.New()
	messages, err := messageIntf.GetMessages(context.TODO(), m.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "getMessages GetMessages error: "+err.Error())
		return
	}

	ctx.JSON(http.StatusOK, messages)
	return
}

func pushMessage(ctx *gin.Context) {
	botClient := bot.GetLineBotClient()
	m := &model.Message{}
	if err := ctx.ShouldBind(m); err != nil {
		ctx.JSON(http.StatusBadRequest, "pushMessage bind error: "+err.Error())
		return
	}

	if m.UserID == "" {
		ctx.JSON(http.StatusBadRequest, "userID shouldn't be none")
		return
	}

	messages := []linebot.SendingMessage{linebot.NewTextMessage(m.Text)}
	_, err := botClient.PushMessage(m.UserID, messages...).Do()
	if err != nil {
		log.Println("push message error:", err)
		return
	}
}

func receiveMessage(ctx *gin.Context) {
	botClient := bot.GetLineBotClient()
	messageIntf := message.New()

	events, err := botClient.ParseRequest(ctx.Request)
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
		m := &model.Message{
			UserID:    profile.UserID,
			Username:  profile.DisplayName,
			Timestamp: &event.Timestamp,
		}

		if event.Type == linebot.EventTypeMessage {
			switch userMessage := event.Message.(type) {
			case *linebot.TextMessage:
				m.Text = userMessage.Text
				m.MessageID = userMessage.ID
				err = messageIntf.ReceiveMessage(context.TODO(), m)
				if err != nil {
					log.Println("InsertMessage error:", err)
					return
				}
			}
		}
	}

}
