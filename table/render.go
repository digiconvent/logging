package table

import (
	"fmt"
	"strconv"
	"strings"
)

func (t *Table) Render() string {
	widths := make([]int, len(t.Headers))
	for i := range t.Headers {
		widths[i] = len(t.Headers[i]) + 2
	}
	for i := range t.Values {
		for j := range t.Values[i] {
			formatted := formatUuuid(t.Values[i][j])
			if widths[j] < len(fmt.Sprint(formatted))+2 {
				widths[j] = len(fmt.Sprint(formatted)) + 2
			}
		}
	}

	headers := make([]string, len(t.Headers))
	for i := range t.Headers {
		headers[i] = fmt.Sprintf(" %-"+strconv.Itoa(widths[i]-1)+"s", t.Headers[i])
	}

	rows := make([]string, len(t.Values))
	for i := range t.Values {
		row := make([]string, len(t.Headers))
		for j := range t.Values[i] {
			formatted := formatUuuid(t.Values[i][j])
			row[j] = fmt.Sprintf(" %-"+strconv.Itoa(widths[j]-1)+"v", formatted)
		}
		rows[i] = "┃" + strings.Join(row, "┃") + "┃\n"
	}

	table := t.Separator(widths, 0)
	table += "┃" + strings.Join(headers, "┃") + "┃\n"
	table += t.Separator(widths, 1)
	table += strings.Join(rows, t.Separator(widths, 1))
	table += t.Separator(widths, 2)
	return table
}
