package main

import (
	"context"
	"linebot/api"
	"linebot/bot"
	"linebot/database"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	database.Initialize(ctx)
	bot.Initialize(ctx)

	router := api.InitRouter()
	router.Run(":8000")
}
