package handler

import (
	"fmt"
	"strings"

	"github.com/matsu-chara/conbot/bot/handler/message"
	"github.com/matsu-chara/conbot/consul"
	myslack "github.com/matsu-chara/conbot/slack"
	"github.com/nlopes/slack"
)

// CommandHandler handle command
type CommandHandler struct {
	slackMessenger *myslack.Messenger
	conbotClient   *consul.ConbotClient
	botID          string
	botName        string
}

// New create CommandHandler
func New(rtm *slack.RTM, conbotClient *consul.ConbotClient, botID string, botName string) *CommandHandler {
	return &CommandHandler{myslack.New(rtm), conbotClient, botID, botName}
}

// HandleCommand handle command
func (handler *CommandHandler) HandleCommand(ev *slack.MessageEvent) {
	words := message.NewSlackWords(ev.Text)

	// accept only bot mention
	if !words.HasMentionAtFirst(handler.botID) {
		return // ignore message
	}

	command := words.Command()
	commandArgs := words.CommandArgs()

	var resultMsgs message.ResultMessage
	switch command {
	case "catalog", "catalogs":
		resultMsgs = handler.handleCatalog(commandArgs)
	case "note", "notes":
		handler.slackMessenger.SendChecking(ev)
		resultMsgs = handler.handleNote(commandArgs)
	case "node", "nodes", "ip", "ips":
		handler.slackMessenger.SendChecking(ev)
		resultMsgs = handler.handleNode(commandArgs)
	case "help":
		commands := []string{
			"catalog [servicePrefix] [tagExact]",
			"version [servicePrefix] [tagExact]",
			"node [servicePrefix] [tagExact]",
			"help",
		}
		resultMsgs = []string{
			fmt.Sprintf("@%s command [args]\ncommands:\n\t-%s", handler.botName, strings.Join(commands, "\n\t")),
		}
	}

	// send result
	handler.slackMessenger.SendMessages(resultMsgs.ResultOrNotFound(), ev)
}
