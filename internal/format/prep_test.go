package format

import (
	"strings"
	"testing"
)

func TestPrep_SingleLine(t *testing.T) {
	result := prep("hello")
	if result != "hello" {
		t.Errorf("unexpected: %q", result)
	}
}

func TestPrep_MultiLine(t *testing.T) {
	result := prep("line1\nline2\nline3")
	if !strings.HasPrefix(result, "line1\n") {
		t.Errorf("first line not preserved: %q", result)
	}
	indent := strings.Repeat(" ", 21)
	if !strings.Contains(result, indent+"line2") {
		t.Errorf("second line not indented: %q", result)
	}
}

func TestPrep_MultiLineEmptySegments(t *testing.T) {
	result := prep("line1\n\nline3")
	if strings.Contains(result, "\n\n") {
		t.Errorf("empty segments should be dropped: %q", result)
	}
}
