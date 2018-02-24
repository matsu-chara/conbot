package message

// ResultMessage contains bot result
type ResultMessage []string

// ResultOrNotFound return result or notFound message
func (result ResultMessage) ResultOrNotFound() []string {
	if len(result) == 0 {
		return []string{"not found."}
	}
	return result
}
