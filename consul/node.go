package consul

import (
	"github.com/hashicorp/consul/api"
	"github.com/matsu-chara/conbot/consul/domain/catalog"
	"github.com/matsu-chara/conbot/consul/domain/node"
)

// GetNode returns [Service, Tag, IpAndPort]
func (client *ConbotClient) GetNode(catalog *catalog.Catalog, tagExacts []string) (*node.ServiceNode, error) {
	consulMutex.Lock()
	defer consulMutex.Unlock()

	catalogClient := client.consulClient.Catalog()

	var queryOptions *api.QueryOptions
	result := map[string][]*api.CatalogService{}
	for service := range catalog.TagsByServiceMap {
		services, _, err := catalogClient.Service(service, "", queryOptions)
		if err != nil {
			return nil, err
		}
		result[service] = services
		sleepForConsul()
	}

	serviceNode := node.BuildServiceNode(result, tagExacts)
	return &serviceNode, nil
}
