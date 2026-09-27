package table

func (t *Table) AddRow(values ...any) {
	valuesCopy := make([]any, len(values))
	copy(valuesCopy, values)
	t.Values = append(t.Values, valuesCopy)
}
