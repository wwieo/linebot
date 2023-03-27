package bot

import (
	"context"
	"github.com/line/line-bot-sdk-go/linebot"
	"os"
)

type Bot interface {
	initialize(context.Context, botConfig)
	client() interface{}
}

type botConfig struct {
	ChannelSecret string
	ChannelToken  string
}

var botIntf Bot

const (
	channelSecret = "channelSecret"
	channelToken  = "channelToken"
)

func Initialize(ctx context.Context) {
	lineBotConfig := botConfig{
		ChannelSecret: os.Getenv(channelSecret),
		ChannelToken:  os.Getenv(channelToken),
	}

	botIntf = &lineBot{}
	botIntf.initialize(ctx, lineBotConfig)
}

func GetLineBotClient() *linebot.Client {
	return botIntf.client().(*linebot.Client)
}
