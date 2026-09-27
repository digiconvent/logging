package logging

import (
	"strings"
	"testing"

	"github.com/digiconvent/logging/internal/color"
)

func TestRed_WrapsInColour(t *testing.T) {
	logger := &Logger{}
	result := logger.Red("x")
	if !strings.Contains(result, "x") || !strings.Contains(result, color.Red) {
		t.Errorf("expected coloured output containing x, got %q", result)
	}
}

func TestRed_NoColorReturnsPlainString(t *testing.T) {
	logger := &Logger{NoColor: true}
	result := logger.Red("x")
	if result != "x" {
		t.Errorf("expected plain string with NoColor set, got %q", result)
	}
}

func TestGreen_WrapsInColour(t *testing.T) {
	logger := &Logger{}
	result := logger.Green("x")
	if !strings.Contains(result, "x") || !strings.Contains(result, color.Green) {
		t.Errorf("expected coloured output containing x, got %q", result)
	}
}

func TestGreen_NoColorReturnsPlainString(t *testing.T) {
	logger := &Logger{NoColor: true}
	result := logger.Green("x")
	if result != "x" {
		t.Errorf("expected plain string with NoColor set, got %q", result)
	}
}

func TestDisabled_WrapsInColour(t *testing.T) {
	logger := &Logger{}
	result := logger.Disabled("x")
	if !strings.Contains(result, "x") || !strings.Contains(result, color.Gray) {
		t.Errorf("expected coloured output containing x, got %q", result)
	}
}

func TestDisabled_NoColorReturnsPlainString(t *testing.T) {
	logger := &Logger{NoColor: true}
	result := logger.Disabled("x")
	if result != "x" {
		t.Errorf("expected plain string with NoColor set, got %q", result)
	}
}
