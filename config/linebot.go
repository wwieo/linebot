package config

import (
	"linebot/model"
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
)

var (
	lineBotConfig = &model.Linebot{}
)

const (
	channelSecret = "linebot.channelSecret"
	channelToken  = "linebot.channelToken"
)

func init() {
	config := utilsController.GetConfig()
	lineBotConfig = &model.Linebot{
		ChannelSecret: config.GetString(channelSecret),
		ChannelToken:  config.GetString(channelToken),
	}
}

func GetLineBotConfig() *model.Linebot {
	return lineBotConfig
}

func GetLinebotClient() *linebot.Client {
	bot, err := linebot.New(
		lineBotConfig.ChannelSecret,
		lineBotConfig.ChannelToken,
	)
	if err != nil {
		log.Fatal(err)
	}
	return bot
}
