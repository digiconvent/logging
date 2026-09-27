package table

import "strings"

func (t *Table) Separator(widths []int, position int) string {
	var start, mid, end string
	switch position {
	case 0:
		start, mid, end = "┏", "┳", "┓"
	case 1:
		start, mid, end = "┣", "╋", "┫"
	default:
		start, mid, end = "┗", "┻", "┛"
	}

	segments := make([]string, len(widths))
	for i := range widths {
		segments[i] = strings.Repeat("━", widths[i])
	}

	return start + strings.Join(segments, mid) + end + "\n"
}
