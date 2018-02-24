package domain

import "strings"

// Tags represent consul tags
type Tags []string

func (ts Tags) String() string {
	return strings.Join(ts, ", ")
}
