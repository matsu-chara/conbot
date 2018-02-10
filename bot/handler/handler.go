package handler

import (
	"fmt"
	"strings"

	"github.com/matsu-chara/conbot/consul"
	"github.com/nlopes/slack"
)

// CommandHandler handle command
type CommandHandler struct {
	rtm          *slack.RTM
	conbotClient *consul.ConbotClient
	botID        string
	botName      string
}

// New create CommandHandler
func New(rtm *slack.RTM, conbotClient *consul.ConbotClient) *CommandHandler {
	return &CommandHandler{rtm, conbotClient, "", ""}
}

// SetBotInfo set bot info
func (commandHandler *CommandHandler) SetBotInfo(botID string, botName string) {
	commandHandler.botID = botID
	commandHandler.botName = botName
}

// HandleCommand handle command
func (commandHandler *CommandHandler) HandleCommand(ev *slack.MessageEvent) {
	words := strings.Split(ev.Text, " ")

	// accept only bot mention
	if words[0] != "<@"+commandHandler.botID+">" {
		return
	}

	// print help if message contains only mention
	if len(words) <= 1 {
		words = append(words, "help")
	}

	var resultMsgs []string
	switch command := words[1]; command {
	case "catalog", "catalogs":
		resultMsgs = commandHandler.handleCatalog(words[2:])
	case "note", "notes":
		commandHandler.rtm.SendMessage(commandHandler.rtm.NewOutgoingMessage("checking...", ev.Channel))
		resultMsgs = commandHandler.handleNote(words[2:])
	case "node", "nodes", "ip", "ips":
		commandHandler.rtm.SendMessage(commandHandler.rtm.NewOutgoingMessage("checking...", ev.Channel))
		resultMsgs = commandHandler.handleNode(words[2:])
	case "help":
		commands := []string{
			"catalog [servicePrefix] [tagExact]",
			"version [servicePrefix] [tagExact]",
			"node [servicePrefix] [tagExact]",
			"help",
		}
		resultMsgs = []string{
			fmt.Sprintf("@%s command [args]\ncommands:\n\t-%s", commandHandler.botName, strings.Join(commands, "\n\t")),
		}
	}

	// send result
	for _, msg := range resultMsgs {
		commandHandler.rtm.SendMessage(commandHandler.rtm.NewOutgoingMessage(msg, ev.Channel))
	}
	if len(resultMsgs) == 0 {
		commandHandler.rtm.SendMessage(commandHandler.rtm.NewOutgoingMessage("not found.", ev.Channel))
	}
}
