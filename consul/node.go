package consul

import (
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/consul/api"
	"github.com/matsu-chara/conbot/tabler"
)

// ServiceNode [Service, [Tag [IpAndPort]]]
type ServiceNode struct {
	nodesByTagAndService map[string]NodesByTag
}

// NodesByTag [Tag, [IpAndPort]]
type NodesByTag struct {
	nodesByTag map[string][]string
}

// GetNode returns [Service, Tag, IpAndPort]
func (client *ConbotClient) GetNode(catalog *Catalog, tagExact string) (*ServiceNode, error) {
	consulMutex.Lock()
	defer consulMutex.Unlock()

	catalogClient := client.consulClient.Catalog()

	var queryOptions *api.QueryOptions
	tag := tagExact

	result := map[string]NodesByTag{}
	for service := range catalog.serviceAndTags {
		services, _, err := catalogClient.Service(service, tag, queryOptions)
		if err != nil {
			return nil, err
		}
		nodesByTag := map[string][]string{}
		for _, service := range services {
			if len(service.ServiceTags) != 1 {
				return nil, fmt.Errorf("unexpected service tags count. tags = %v", strings.Join(service.ServiceTags, ", "))
			}
			tag := service.ServiceTags[0]
			addressAndPort := fmt.Sprintf("%s:%d", service.Address, service.ServicePort)
			nodesByTag[tag] = append(nodesByTag[tag], addressAndPort)
		}
		result[service] = NodesByTag{nodesByTag}
		sleepForConsul()
	}
	return &ServiceNode{result}, nil
}

// ToSlackFormat format
func (node ServiceNode) ToSlackFormat() []string {
	// TODO: refactor
	services := []string{}
	for s := range node.nodesByTagAndService {
		services = append(services, s)
	}
	sort.Strings(services)

	groupedServices := [][]string{}
	for i := 0; i < len(services); i += 5 {
		to := i + 5
		if to > len(services) {
			to = len(services)
		}
		groupedServices = append(groupedServices, services[i:to])
	}

	result := []string{}
	for _, partOfServices := range groupedServices {
		result = append(result, node.servicesToSlackTable(partOfServices))
	}
	return result
}

func (node ServiceNode) servicesToSlackTable(services []string) string {
	// TODO: refactor
	data := [][]string{}
	for _, service := range services {
		nodesByTag := node.nodesByTagAndService[service]

		tags := []string{}
		for t := range node.nodesByTagAndService[service].nodesByTag {
			tags = append(tags, t)
		}
		sort.Strings(tags)

		for _, tag := range tags {
			nodes := nodesByTag.nodesByTag[tag]
			sort.Strings(nodes)
			data = append(data, []string{service, tag, strings.Join(nodes, " ")})
		}
	}
	return tabler.ToSlackTable([]string{"Service", "Tag", "Node"}, data)
}
