package logging

import (
	"bytes"
	"strings"
	"testing"
)

func TestMute_SuppressesEverything(t *testing.T) {
	var buf bytes.Buffer
	logger := New(&buf)
	logger.Mute()

	logger.Error("boom")
	logger.Warning("careful")
	logger.Info("fyi")
	logger.Success("done")

	if buf.Len() != 0 {
		t.Errorf("expected no output while muted, got %q", buf.String())
	}
}

func TestUnmute_RestoresOutput(t *testing.T) {
	var buf bytes.Buffer
	logger := New(&buf)
	logger.Mute()
	logger.Unmute()

	logger.Info("fyi")

	if !strings.Contains(buf.String(), "fyi") {
		t.Errorf("expected output to resume after Unmute, got %q", buf.String())
	}
}

func TestShowFrom_KeepsMessagesAtOrAboveThreshold(t *testing.T) {
	var buf bytes.Buffer
	logger := New(&buf)
	logger.ShowFrom(LevelWarning)

	logger.Warning("careful")
	logger.Error("boom")

	output := buf.String()
	if !strings.Contains(output, "careful") || !strings.Contains(output, "boom") {
		t.Errorf("expected Warning and Error to survive a Warning threshold, got %q", output)
	}
}
