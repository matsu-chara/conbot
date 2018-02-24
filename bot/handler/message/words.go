package message

import (
	"fmt"
	"strings"
)

// SlackWords holds slack word spplitted by space
type SlackWords []string

// NewSlackWords create SlackWords
func NewSlackWords(message string) SlackWords {
	return strings.Split(message, " ")
}

// HasMentionAtFirst checks words has mention at first
func (words SlackWords) HasMentionAtFirst(userName string) bool {
	return words[0] == fmt.Sprintf("<@%s>", userName)
}

// Command return command
func (words SlackWords) Command() string {
	// print help if message does not have command
	if len(words) <= 1 {
		return "help"
	}
	return words[1]
}

// CommandArgs returns args
func (words SlackWords) CommandArgs() []string {
	if len(words) <= 2 {
		return []string{}
	}
	return words[2:]
}
