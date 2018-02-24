package note

import (
	"sort"
	"strings"

	"github.com/hashicorp/consul/api"
	"github.com/matsu-chara/conbot/libs"
	"github.com/matsu-chara/conbot/tabler"
)

type notePerService []string

func buildFromChecksPerService(checks api.HealthChecks) notePerService {
	note := []string{}
	for _, check := range checks {
		note = append(note, check.Notes)
	}
	return note
}

// ServiceNote Service Note Map
type ServiceNote struct {
	notesByService map[string][]string
}

// BuildFromChecks build notePerService from check
func BuildFromChecks(checksByService map[string]api.HealthChecks) ServiceNote {
	notes := map[string][]string{}
	for service, check := range checksByService {
		notes[service] = buildFromChecksPerService(check)
	}
	return ServiceNote{notes}
}

func (note ServiceNote) sortedService() []string {
	result := []string{}
	for service := range note.notesByService {
		result = append(result, service)
	}
	sort.Strings(result)
	return result
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
		libs.RemoveDuplicates(&notes)
		data = append(data, []string{service, strings.Join(notes, ", ")})
	}
	return tabler.ToSlackTable([]string{"Service", "Notes"}, data)
}
