package consul

import (
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/consul/api"
	"github.com/matsu-chara/conbot/tabler"
)

// Catalog [Service, [Tag]].
type Catalog struct {
	serviceAndTags map[string][]string
}

// GetCatalog returns [Service, [Tag]]
func (client *ConbotClient) GetCatalog(servicePrefix string, tagExact string) (*Catalog, error) {
	consulMutex.Lock()
	defer consulMutex.Unlock()

	catalogClient := client.consulClient.Catalog()

	var queryOptions *api.QueryOptions
	serviceAndTagMap, _, err := catalogClient.Services(queryOptions)
	if err != nil {
		return nil, err
	}

	response := Catalog{
		serviceAndTagMap,
	}
	filtered := response.filterByServicePrefix(servicePrefix).filterByTagExact(tagExact)
	return filtered, nil
}

func (catalog *Catalog) filterByServicePrefix(servicePrefix string) *Catalog {
	if servicePrefix == "" {
		return catalog
	}

	result := map[string][]string{}
	for service, tags := range catalog.serviceAndTags {
		if strings.HasPrefix(service, servicePrefix) {
			result[service] = tags
		}
	}
	return &Catalog{
		serviceAndTags: result,
	}
}

func (catalog *Catalog) filterByTagExact(tagExact string) *Catalog {
	if tagExact == "" {
		return catalog
	}

	result := map[string][]string{}
	for service, tags := range catalog.serviceAndTags {
		for _, tag := range tags {
			if tag == tagExact {
				result[service] = []string{tag}
			}
		}
	}
	return &Catalog{
		serviceAndTags: result,
	}
}

// ToSlackFormat format
func (catalog Catalog) ToSlackFormat() string {
	data := [][]string{}

	services := []string{}
	for s := range catalog.serviceAndTags {
		services = append(services, s)
	}
	sort.Strings(services)

	for _, service := range services {
		tags := catalog.serviceAndTags[service]
		sort.Strings(tags)
		tagsString := fmt.Sprintf("%s", strings.Join(tags, ", "))
		data = append(data, []string{service, tagsString})
	}
	return tabler.ToSlackTable([]string{"Service", "Tags"}, data)
}
