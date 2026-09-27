package table

import "testing"

func TestFormatUuuid_NonUUID(t *testing.T) {
	result := formatUuuid("not-a-uuid")
	if result != "not-a-uuid" {
		t.Errorf("non-UUID should be returned as-is, got %v", result)
	}
}

func TestFormatUuuid_TruncatesUUID(t *testing.T) {
	result := formatUuuid("12345678-1234-1234-1234-123456789abc")
	if result != "1234...9abc" {
		t.Errorf("expected truncated UUID, got %v", result)
	}
}
