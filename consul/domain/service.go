package domain

import (
	"sort"

	"github.com/matsu-chara/conbot/consul/domain/query"
	"github.com/matsu-chara/conbot/libs"
)

// Services is slice of service name
type Services []string

// TagsByServiceMap is [servicename, tags]
type TagsByServiceMap map[string][]string

// SortedServices return services
func (tsm TagsByServiceMap) SortedServices() Services {
	result := []string{}
	for s := range tsm {
		result = append(result, s)
	}
	sort.Strings(result)
	return result
}

// SortedTags returns tags
func (tsm TagsByServiceMap) SortedTags(service string) Tags {
	result := tsm[service]
	sort.Strings(result)
	return result
}

// RemoveService remove service
func (tsm TagsByServiceMap) RemoveService(service string) {
	_, ok := tsm[service]
	if ok {
		delete(tsm, service)
	}
}

// FilterByServicePrefixes filters
func (tsm TagsByServiceMap) FilterByServicePrefixes(servicePrefixes query.ServicePrefixes) TagsByServiceMap {
	return libs.FilterByKey(tsm, func(service string) bool {
		return servicePrefixes.Matches(service)
	})
}

// FilterByTagExact filters
func (tsm TagsByServiceMap) FilterByTagExact(tagExacts query.TagExacts) TagsByServiceMap {
	return libs.FilterByValues(tsm, func(tags []string) ([]string, bool) {
		return tagExacts.Matches(tags)
	})
}
