package table

type Table struct {
	Headers []string
	Values  [][]any
}

func NewTable(headers []string) Table {
	return Table{
		Headers: headers,
		Values:  [][]any{},
	}
}
