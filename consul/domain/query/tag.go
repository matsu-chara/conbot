package query

import "github.com/matsu-chara/conbot/libs"

// TagExacts has parsed tags
type TagExacts []string

// ParseTagExacts return tags
func ParseTagExacts(arg []string, index int) TagExacts {
	if len(arg) > index {
		return parse(arg[index])
	}
	return []string{}
}

// Matches check service match prefix condition
func (exacts TagExacts) Matches(tag []string) ([]string, bool) {
	if len(exacts) == 0 {
		return tag, true // if no conditions was defined, then no filter
	}

	exaxtSet := libs.FromSlice(exacts)
	for _, t := range tag {
		_, ok := exaxtSet[t]
		if ok {
			return []string{t}, true
		}
	}
	return []string{}, false
}

// MatchesOnly check service match prefix condition for 1 tag
func (exacts TagExacts) MatchesOnly(tag string) bool {
	_, ok := exacts.Matches([]string{tag})
	return ok
}
