package catalog

import (
	"github.com/matsu-chara/conbot/consul/domain"
	"github.com/matsu-chara/conbot/tabler"
)

// Catalog [Service, [Tag]].
type Catalog struct {
	TagsByServiceMap domain.TagsByServiceMap
}

// New return Catalog
func New(tags domain.TagsByServiceMap) Catalog {
	return Catalog{tags}
}

// ToSlackFormat format
func (catalog Catalog) ToSlackFormat() string {
	data := [][]string{}

	for _, service := range catalog.TagsByServiceMap.SortedServices() {
		soretedTags := catalog.TagsByServiceMap.SortedTags(service)
		data = append(data, []string{service, soretedTags.String()}) // {service, "tag1, tag2, tag3"}
	}
	return tabler.ToSlackTable([]string{"Service", "Tags"}, data)
}
