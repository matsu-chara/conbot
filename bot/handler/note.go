package handler

import (
	"github.com/matsu-chara/conbot/consul/domain/query"
)

type noteArg struct {
	servicePrefixes query.ServicePrefixes
	tagExacts       query.TagExacts
}

func parseNoteArg(args []string) noteArg {
	servicePrefixes := query.ParseServicePrefixes(args, 0)
	tagExacts := query.ParseTagExacts(args, 1)
	return noteArg{servicePrefixes, tagExacts}
}

// HandleNote parse arg, get notes and format.
// command ex: note service
func (handler *CommandHandler) handleNote(args []string) []string {
	parsedArg := parseNoteArg(args)

	catalog, err := handler.conbotClient.GetCatalog(parsedArg.servicePrefixes, parsedArg.tagExacts)
	if err != nil {
		return []string{err.Error()}
	}

	note, err := handler.conbotClient.GetNote(catalog, parsedArg.tagExacts)
	if err != nil {
		return []string{err.Error()}
	}

	return []string{note.ToSlackFormat()}
}
