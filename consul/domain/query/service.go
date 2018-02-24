package query

import "strings"

// ServicePrefixes has parsed service prefixes
type ServicePrefixes []string

// ParseServicePrefixes return ServicePrefixes
func ParseServicePrefixes(arg []string, index int) ServicePrefixes {
	if len(arg) > index {
		return parse(arg[index])
	}
	return []string{}
}

// Matches check service match prefix condition
func (prefixes ServicePrefixes) Matches(service string) bool {
	if len(prefixes) == 0 {
		return true // if no conditions was defined, then no filter
	}

	for _, servicePrefix := range prefixes {
		if strings.HasPrefix(service, servicePrefix) {
			return true
		}
	}
	return false
}
