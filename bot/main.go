package main

import (
	"bot/config"
	"fmt"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func main() {
	bot, err := messaging_api.NewMessagingApiAPI(
		config.LINE_CHANNEL_TOKEN,
	)
	if err != nil {
		panic(err)
	}

	// Get the message quota
	res, _ := bot.GetMessageQuota()
	fmt.Print(res.Value)
}
