package main

import (
	"fmt"
	view "linebot/view"
)

func main() {
	fmt.Println("hello world!")
	router := view.InitRouter()
	router.Run(":8000")
}
