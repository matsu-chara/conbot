package bot

import (
	"log"

	"github.com/matsu-chara/conbot/bot/handler"
	"github.com/matsu-chara/conbot/consul"
	"github.com/nlopes/slack"
)

// Run activate bot routine
func Run(api *slack.Client, conbot *consul.ConbotClient) int {
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	// initialized when connected
	var commandHandler *handler.CommandHandler

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				log.Print("hello")
			case *slack.ConnectedEvent:
				log.Print("connected")
				commandHandler = handler.New(
					rtm,
					conbot,
					ev.Info.User.ID,
					ev.Info.User.Name,
				)
			case *slack.InvalidAuthEvent:
				log.Printf("Invalid credentials. %v", ev)
				return 1
			case *slack.MessageEvent:
				commandHandler.HandleCommand(ev)
			}
		}
	}
}
