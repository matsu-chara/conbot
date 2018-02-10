package handler

// HandleCatalog parse arg, call catalog api and format.
// command ex: catalog stg account
func (handler *CommandHandler) handleCatalog(args []string) []string {
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
	return []string{catalog.ToSlackFormat()}
}
