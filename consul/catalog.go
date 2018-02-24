package consul

import (
	"github.com/hashicorp/consul/api"
	"github.com/matsu-chara/conbot/consul/domain"
	"github.com/matsu-chara/conbot/consul/domain/catalog"
)

// GetCatalog returns [Service, [Tag]]
func (client *ConbotClient) GetCatalog(servicePrefixes []string, tagExacts []string) (*catalog.Catalog, error) {
	consulMutex.Lock()
	defer consulMutex.Unlock()

	catalogClient := client.consulClient.Catalog()

	var queryOptions *api.QueryOptions
	result, _, err := catalogClient.Services(queryOptions)
	tagsByService := domain.TagsByServiceMap(result)
	if err != nil {
		return nil, err
	}

	// remove consul
	tagsByService.RemoveService("consul")

	filtered := tagsByService.FilterByServicePrefixes(servicePrefixes).FilterByTagExact(tagExacts)

	response := catalog.New(filtered)
	return &response, nil
}
