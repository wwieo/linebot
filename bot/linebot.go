package bot

import (
	"context"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
)

type lineBot struct {
	*linebot.Client
}

// initialize initializes the MySQL dao handle.
func (bot *lineBot) initialize(ctx context.Context, cfg botConfig) {
	client, err := linebot.New(
		cfg.ChannelSecret,
		cfg.ChannelToken,
	)
	if err != nil {
		log.Fatal(err)
	}

	bot.Client = client
}

func (bot *lineBot) client() interface{} {
	return bot.Client
}
