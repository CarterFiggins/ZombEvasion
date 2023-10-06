package main

import(
	"log"
	"os"

	"infection/bot"
)

func main() {
	botToken, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		log.Fatal("Must set Discord token as env variable: BOT_TOKEN")
	}

	bot.BotToken = botToken
	bot.Run()
}