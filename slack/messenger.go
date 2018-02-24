package slack

import (
	"github.com/nlopes/slack"
)

type slackRtm interface {
	SendMessage(msg *slack.OutgoingMessage)
	NewOutgoingMessage(text string, channel string) *slack.OutgoingMessage
}

// Messenger handle slack rtm
type Messenger struct {
	rtm slackRtm
}

// New returns instance
func New(rtm *slack.RTM) *Messenger {
	return &Messenger{rtm}
}

// SendChecking send "checking..." message
func (messenger *Messenger) SendChecking(ev *slack.MessageEvent) {
	messenger.sendMessage("checking...", ev)
}

// SendMessages send messages
func (messenger *Messenger) SendMessages(msgs []string, ev *slack.MessageEvent) {
	for _, msg := range msgs {
		messenger.sendMessage(msg, ev)
	}
}

func (messenger *Messenger) sendMessage(msg string, ev *slack.MessageEvent) {
	messenger.rtm.SendMessage(messenger.rtm.NewOutgoingMessage(msg, ev.Channel))
}
