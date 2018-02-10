package handler

// HandleNode parse arg, get nodes and format.
// command ex: node stg account
func (handler *CommandHandler) handleNode(args []string) []string {
	servicePrefix := ""
	tagExact := ""
	if len(args) >= 1 {
		servicePrefix = args[0]
	}
	if len(args) >= 2 {
		tagExact = args[1]
	}

	catalog, err := handler.conbotClient.GetCatalog(servicePrefix, tagExact)
	if err != nil {
		return []string{err.Error()}
	}

	node, err := handler.conbotClient.GetNode(catalog, tagExact)
	if err != nil {
		return []string{err.Error()}
	}

	return node.ToSlackFormat()
}
