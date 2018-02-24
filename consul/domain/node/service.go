package node

import (
	"sort"
	"strings"

	"github.com/hashicorp/consul/api"
	"github.com/matsu-chara/conbot/consul/domain/query"
	"github.com/matsu-chara/conbot/libs"
	"github.com/matsu-chara/conbot/tabler"
)

// ServiceNode [Service, [Tag [IpAndPort]]]
type ServiceNode struct {
	nodesByTagAndService map[string]NodesByTag
}

// BuildServiceNode returns our struct
func BuildServiceNode(serviceByServiceNames map[string][]*api.CatalogService, tagExacts query.TagExacts) ServiceNode {
	result := map[string]NodesByTag{}
	for serviceName, services := range serviceByServiceNames {
		result[serviceName] = buildNodesByTag(services, tagExacts)
	}
	return ServiceNode{result}
}

func (node ServiceNode) sortedService() []string {
	services := []string{}
	for s := range node.nodesByTagAndService {
		services = append(services, s)
	}
	sort.Strings(services)
	return services
}

func (node ServiceNode) sortedTags(service string) []string {
	tags := []string{}
	for t := range node.nodesByTagAndService[service] {
		tags = append(tags, t)
	}
	sort.Strings(tags)
	return tags
}

// ToSlackFormat format
func (node ServiceNode) ToSlackFormat() []string {
	services := node.sortedService()
	groupedServices := libs.Grouped(services, 5) // avoid too long response

	result := []string{}
	for _, partOfServices := range groupedServices {
		result = append(result, node.servicesToSlackTable(partOfServices))
	}
	return result
}

func (node ServiceNode) servicesToSlackTable(services []string) string {
	data := [][]string{}
	for _, service := range services {
		data = append(data, node.oneServiceToSlackRows(service)...)
	}
	return tabler.ToSlackTable([]string{"Service", "Tag", "Node"}, data)
}

func (node ServiceNode) oneServiceToSlackRows(service string) [][]string {
	nodesByTag := node.nodesByTagAndService[service]
	tags := node.sortedTags(service)

	serviceNode := [][]string{}
	for _, tag := range tags {
		nodes := nodesByTag[tag]
		sort.Strings(nodes)
		serviceNode = append(serviceNode, []string{service, tag, strings.Join(nodes, " ")})
	}
	return serviceNode
}
