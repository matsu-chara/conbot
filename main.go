package main

import (
	"os"

	"github.com/matsu-chara/conbot/bot"
	"github.com/matsu-chara/conbot/consul"
	"github.com/nlopes/slack"
)

func main() {
	api := slack.New(os.Getenv("CONBOT_SLACK_BOT_TOKEN"))
	consulClient, err := consul.Connect(os.Getenv("CONBOT_CONSUL"))
	if err != nil {
		panic(err)
	}
	os.Exit(bot.Run(api, consulClient))
}
