package format

import (
	"strings"
	"testing"

	"github.com/digiconvent/logging/internal/color"
)

func TestLine_IncludesMessageAndColour(t *testing.T) {
	result := Line(color.Red, "boom", false)
	if !strings.Contains(result, "boom") {
		t.Errorf("expected message in line, got %q", result)
	}
	if !strings.Contains(result, color.Red) {
		t.Errorf("expected severity colour in line, got %q", result)
	}
}

func TestLine_NoColorOmitsEscapeCodes(t *testing.T) {
	result := Line(color.Red, "boom", true)
	if strings.Contains(result, "\033") {
		t.Errorf("expected no ANSI escape codes with noColor set, got %q", result)
	}
	if !strings.Contains(result, "boom") {
		t.Errorf("expected message to still be present, got %q", result)
	}
}
