package handler

// HandleNote parse arg, get nodes and format.
// command ex: note stg account
func (handler *CommandHandler) handleNote(args []string) []string {
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

	note, err := handler.conbotClient.GetNote(catalog, tagExact)
	if err != nil {
		return []string{err.Error()}
	}

	return []string{note.ToSlackFormat()}
}
