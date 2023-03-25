package main

import (
	"context"
	"fmt"
	"linebot/api"
	"linebot/database"
)

func main() {
	fmt.Println("hello world!")

	database.Initialize(context.Background())

	router := api.InitRouter()
	router.Run(":8000")
}
