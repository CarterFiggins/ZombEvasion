package main

import(
	"infection/mongo"
	"infection/bot"
)

func main() {
	mongo.Connect()
	defer mongo.Close()

	bot.Run()
}