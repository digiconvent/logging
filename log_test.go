package logging

import (
	"bytes"
	"strings"
	"sync"
	"testing"

	"github.com/digiconvent/logging/internal/color"
)

func TestError_WritesRedMessage(t *testing.T) {
	var buf bytes.Buffer
	logger := New(&buf)

	logger.Error("boom")

	if !strings.Contains(buf.String(), "boom") {
		t.Errorf("expected output to contain message, got %q", buf.String())
	}
	if !strings.Contains(buf.String(), color.Red) {
		t.Errorf("expected output to contain red colour code, got %q", buf.String())
	}
}

func TestError_WritesEvenWhenShowingOnlyErrors(t *testing.T) {
	var buf bytes.Buffer
	logger := New(&buf)
	logger.ShowFrom(LevelError)

	logger.Error("boom")

	if !strings.Contains(buf.String(), "boom") {
		t.Errorf("expected Error to survive an Error threshold, got %q", buf.String())
	}
}

func TestWarning_WritesYellowMessage(t *testing.T) {
	var buf bytes.Buffer
	logger := New(&buf)

	logger.Warning("careful")

	if !strings.Contains(buf.String(), "careful") {
		t.Errorf("expected output to contain message, got %q", buf.String())
	}
	if !strings.Contains(buf.String(), color.Yellow) {
		t.Errorf("expected output to contain yellow colour code, got %q", buf.String())
	}
}

func TestWarning_SuppressedWhenShowingOnlyErrors(t *testing.T) {
	var buf bytes.Buffer
	logger := New(&buf)
	logger.ShowFrom(LevelError)

	logger.Warning("careful")

	if buf.Len() != 0 {
		t.Errorf("expected Warning to be suppressed at an Error threshold, got %q", buf.String())
	}
}

func TestInfo_WritesCyanMessage(t *testing.T) {
	var buf bytes.Buffer
	logger := New(&buf)

	logger.Info("fyi")

	if !strings.Contains(buf.String(), "fyi") {
		t.Errorf("expected output to contain message, got %q", buf.String())
	}
	if !strings.Contains(buf.String(), color.Cyan) {
		t.Errorf("expected output to contain cyan colour code, got %q", buf.String())
	}
}

func TestInfo_SuppressedWhenShowingOnlyWarnings(t *testing.T) {
	var buf bytes.Buffer
	logger := New(&buf)
	logger.ShowFrom(LevelWarning)

	logger.Info("fyi")

	if buf.Len() != 0 {
		t.Errorf("expected Info to be suppressed at a Warning threshold, got %q", buf.String())
	}
}

func TestSuccess_WritesGreenMessage(t *testing.T) {
	var buf bytes.Buffer
	logger := New(&buf)

	logger.Success("done")

	if !strings.Contains(buf.String(), "done") {
		t.Errorf("expected output to contain message, got %q", buf.String())
	}
	if !strings.Contains(buf.String(), color.Green) {
		t.Errorf("expected output to contain green colour code, got %q", buf.String())
	}
}

func TestSuccess_SuppressedWhenShowingOnlyWarnings(t *testing.T) {
	var buf bytes.Buffer
	logger := New(&buf)
	logger.ShowFrom(LevelWarning)

	logger.Success("done")

	if buf.Len() != 0 {
		t.Errorf("expected Success to be suppressed at a Warning threshold, got %q", buf.String())
	}
}

func TestWrite_NoColorOmitsEscapeCodes(t *testing.T) {
	var buf bytes.Buffer
	logger := New(&buf)
	logger.NoColor = true

	logger.Error("boom")

	if strings.Contains(buf.String(), "\033") {
		t.Errorf("expected no ANSI escape codes with NoColor set, got %q", buf.String())
	}
	if !strings.Contains(buf.String(), "boom") {
		t.Errorf("expected message to still be written, got %q", buf.String())
	}
}

func TestWrite_ConcurrentCallsDoNotRace(t *testing.T) {
	var buf bytes.Buffer
	logger := New(&buf)

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			logger.Info("concurrent")
		}()
	}
	wg.Wait()
}
