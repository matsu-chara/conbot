package node

import (
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/matsu-chara/conbot/consul/domain/query"
	"github.com/matsu-chara/conbot/libs"
)

// NodesByTag [Tag, [IpAndPort]]
type NodesByTag map[string][]string

func buildNodesByTag(services []*api.CatalogService, tagExacts query.TagExacts) NodesByTag {
	nodesByTag := NodesByTag{}
	for _, service := range services {
		targetTags := libs.Filter(service.ServiceTags, func(t string) bool { return tagExacts.MatchesOnly(t) })
		for _, tag := range targetTags {
			addressAndPort := fmt.Sprintf("%s:%d", service.Address, service.ServicePort)
			nodesByTag[tag] = append(nodesByTag[tag], addressAndPort)
		}
	}
	return nodesByTag
}
