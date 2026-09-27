package table

import "testing"

func TestAddRow_AppendsRow(t *testing.T) {
	table := NewTable([]string{"a", "b"})
	table.AddRow(1, 2)
	if len(table.Values) != 1 || table.Values[0][0] != 1 || table.Values[0][1] != 2 {
		t.Errorf("expected row to be appended as given, got %v", table.Values)
	}
}

func TestAddRow_CopiesValues(t *testing.T) {
	table := NewTable([]string{"a"})
	values := []any{1}
	table.AddRow(values...)
	values[0] = 2

	if table.Values[0][0] != 1 {
		t.Errorf("expected stored row to be unaffected by later mutation of caller's slice, got %v", table.Values[0])
	}
}
