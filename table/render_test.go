package table

import (
	"strings"
	"testing"
)

func TestRender_IncludesHeadersAndValues(t *testing.T) {
	table := NewTable([]string{"name", "age"})
	table.AddRow("alice", 30)

	result := table.Render()

	for _, want := range []string{"name", "age", "alice", "30"} {
		if !strings.Contains(result, want) {
			t.Errorf("expected rendered table to contain %q, got:\n%s", want, result)
		}
	}
}

func TestRender_TruncatesUUIDValues(t *testing.T) {
	table := NewTable([]string{"id"})
	table.AddRow("12345678-1234-1234-1234-123456789abc")

	result := table.Render()

	if !strings.Contains(result, "1234...9abc") {
		t.Errorf("expected UUID value to be truncated, got:\n%s", result)
	}
	if strings.Contains(result, "12345678-1234-1234-1234-123456789abc") {
		t.Errorf("expected full UUID not to appear, got:\n%s", result)
	}
}
