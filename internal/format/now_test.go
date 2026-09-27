package format

import "testing"

func TestNow_MatchesExpectedLength(t *testing.T) {
	ts := now()
	if len(ts) != len("2006.01.02T15:04:05") {
		t.Errorf("unexpected time format length: %q", ts)
	}
}
