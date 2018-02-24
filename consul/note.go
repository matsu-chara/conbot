package consul

import (
	"github.com/hashicorp/consul/api"
	"github.com/matsu-chara/conbot/consul/domain/catalog"
	"github.com/matsu-chara/conbot/consul/domain/note"
)

// GetNote returns [Service, Note]
func (client *ConbotClient) GetNote(catalog *catalog.Catalog, tagExacts []string) (*note.ServiceNote, error) {
	consulMutex.Lock()
	defer consulMutex.Unlock()

	healthClient := client.consulClient.Health()
	var queryOptions *api.QueryOptions

	result := map[string]api.HealthChecks{}
	for service := range catalog.TagsByServiceMap {
		checks, _, err := healthClient.Checks(service, queryOptions)
		if err != nil {
			return nil, err
		}
		result[service] = checks
		sleepForConsul()
	}

	serviceNote := note.BuildFromChecks(result)
	return &serviceNote, nil
}
