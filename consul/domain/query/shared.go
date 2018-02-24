package query

import (
	"strings"

	"github.com/matsu-chara/conbot/libs"
)

func parse(arg string) []string {
	splitted := strings.Split(arg, ",")
	nonEmpty := libs.Filter(splitted, func(s string) bool { return s != "" })
	wildFilter := libs.Filter(nonEmpty, func(s string) bool { return s != "*" }) // specify *, getAll
	trimmed := libs.Map(wildFilter, func(s string) string { return strings.Trim(s, " ") })
	return trimmed
}
