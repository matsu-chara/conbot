package handler

import (
	"github.com/matsu-chara/conbot/consul/domain/query"
)

type catalogArg struct {
	servicePrefixes query.ServicePrefixes
	tagExacts       query.TagExacts
}

func parseCatalogArg(args []string) catalogArg {
	servicePrefixes := query.ParseServicePrefixes(args, 0)
	tagExacts := query.ParseTagExacts(args, 1)
	return catalogArg{servicePrefixes, tagExacts}
}

// HandleCatalog parse arg, call catalog api and format.
// command ex: catalog service
func (handler *CommandHandler) handleCatalog(args []string) []string {
	parsedArg := parseCatalogArg(args)
	catalog, err := handler.conbotClient.GetCatalog(parsedArg.servicePrefixes, parsedArg.tagExacts)
	if err != nil {
		return []string{err.Error()}
	}
	return []string{catalog.ToSlackFormat()}
}
