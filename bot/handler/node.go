package handler

import "github.com/matsu-chara/conbot/consul/domain/query"

type nodeArg struct {
	servicePrefixes query.ServicePrefixes
	tagExacts       query.TagExacts
}

func parseNodeArg(args []string) nodeArg {
	servicePrefixes := query.ParseServicePrefixes(args, 0)
	tagExacts := query.ParseTagExacts(args, 1)
	return nodeArg{servicePrefixes, tagExacts}
}

// HandleNode parse arg, get nodes and format.
// command ex: node service
func (handler *CommandHandler) handleNode(args []string) []string {
	parsedArg := parseNodeArg(args)

	catalog, err := handler.conbotClient.GetCatalog(parsedArg.servicePrefixes, parsedArg.tagExacts)
	if err != nil {
		return []string{err.Error()}
	}

	node, err := handler.conbotClient.GetNode(catalog, parsedArg.tagExacts)
	if err != nil {
		return []string{err.Error()}
	}

	return node.ToSlackFormat()
}
