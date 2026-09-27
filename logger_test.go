package logging

import (
	"bytes"
	"strings"
	"testing"
)

func TestNew_WritesToGivenOutput(t *testing.T) {
	var buf bytes.Buffer
	logger := New(&buf)

	logger.Info("fyi")

	if !strings.Contains(buf.String(), "fyi") {
		t.Errorf("expected New to write to the given output, got %q", buf.String())
	}
}

func TestLogger_ZeroValueWritesToStdout(t *testing.T) {
	var logger Logger

	if logger.writer() == nil {
		t.Error("expected zero-value Logger to default to a non-nil writer")
	}
}
