package consul

import (
	"sort"
	"strings"

	"github.com/hashicorp/consul/api"
	"github.com/matsu-chara/conbot/tabler"
)

// ServiceNote Service Note Map
type ServiceNote struct {
	notesByService map[string][]string
}

// GetNote returns [Service, Note]
func (client *ConbotClient) GetNote(catalog *Catalog, tagExact string) (*ServiceNote, error) {
	consulMutex.Lock()
	defer consulMutex.Unlock()

	healthClient := client.consulClient.Health()
	var queryOptions *api.QueryOptions

	result := map[string][]string{}
	for service := range catalog.serviceAndTags {
		checks, _, err := healthClient.Checks(service, queryOptions)
		if err != nil {
			return nil, err
		}
		notes := []string{}
		for _, check := range checks {
			if tagExact != "" && !contains(check.ServiceTags, tagExact) {
				continue
			}
			notes = append(notes, check.Notes)
		}
		result[service] = notes
		sleepForConsul()
	}
	return &ServiceNote{result}, nil
}

// ToSlackFormat format
func (note ServiceNote) ToSlackFormat() string {
	data := [][]string{}

	services := []string{}
	for s := range note.notesByService {
		services = append(services, s)
	}
	sort.Strings(services)

	for _, service := range services {
		notes := note.notesByService[service]
		removeDuplicates(&notes)
		data = append(data, []string{service, strings.Join(notes, ", ")})
	}
	return tabler.ToSlackTable([]string{"Service", "Notes"}, data)
}
