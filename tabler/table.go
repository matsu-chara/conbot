package tabler

import (
	"bytes"

	"github.com/olekukonko/tablewriter"
)

// ToSlackTable format data
func ToSlackTable(titles []string, data [][]string) string {
	buf := new(bytes.Buffer)
	table := tablewriter.NewWriter(buf)
	table.SetHeader(titles)
	table.AppendBulk(data)
	buf.Write([]byte("```\n"))
	table.Render()
	buf.Write([]byte("```"))
	s := buf.String()
	return s
}
