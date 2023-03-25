package bot

import (
	"context"
	"github.com/line/line-bot-sdk-go/linebot"
	"linebot/utils"
)

type Bot interface {
	initialize(context.Context, botConfig)
	client() interface{}
}

type botConfig struct {
	ChannelSecret string
	ChannelToken  string
}

const (
	channelSecret = "linebot.channelSecret"
	channelToken  = "linebot.channelToken"
)

var botIntf Bot

func Initialize(ctx context.Context) {
	config := utils.GetConfig()
	lineBotConfig := botConfig{
		ChannelSecret: config.GetString(channelSecret),
		ChannelToken:  config.GetString(channelToken),
	}

	botIntf.initialize(ctx, lineBotConfig)
}

func GetLineBotClient() *linebot.Client {
	return botIntf.client().(*linebot.Client)
}
