package utils_test

import (
	"regexp"
	"testing"

	"github.com/Deathfireofdoom/excel-client-go/pkg/utils" // Ensure to use your actual import path
)

func TestGenerateUUID(t *testing.T) {
	uuid, err := utils.GenerateUUID()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if len(uuid) != 32 {
		t.Fatalf("expected length of 32, got: %d", len(uuid))
	}

	// Ensure all characters are hexadecimal
	for _, c := range uuid {
		if (c < '0' || c > '9') && (c < 'a' || c > 'f') {
			t.Fatalf("expected hexadecimal character, got: %c", c)
		}
	}

	// Test version and variant
	version := uuid[12:13]
	if version != "4" {
		t.Fatalf("expected version 4, got: %s", version)
	}

	variant := uuid[16:17]
	if !(variant >= "8" && variant <= "b") {
		t.Fatalf("expected variant character to be between 8 and b (inclusive), got: %s", variant)
	}
}

// Additional test to check the uniqueness of generated UUIDs.
func TestGenerateUUIDUniqueness(t *testing.T) {
	uuidSet := make(map[string]struct{})

	for i := 0; i < 1000; i++ {
		uuid, err := utils.GenerateUUID()
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}

		if _, exists := uuidSet[uuid]; exists {
			t.Fatalf("duplicate UUID generated: %s", uuid)
		}

		uuidSet[uuid] = struct{}{}
	}
}

// Additional test to check the format of generated UUIDs.
func TestGenerateUUIDFormat(t *testing.T) {
	uuid, err := utils.GenerateUUID()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	// Using regex to check UUID format (ensure it's a valid format)
	matched, err := regexp.MatchString("^[0-9a-f]{32}$", uuid)
	if err != nil {
		t.Fatalf("unexpected error during regex match: %v", err)
	}

	if !matched {
		t.Fatalf("UUID does not match expected format: %s", uuid)
	}
}
