package banner

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	err := os.MkdirAll("banners", 0755)
	if err != nil {
		t.Fatalf("failed to create test directory: %v", err)
	}

	err = os.WriteFile(
		"banners/test.txt",
		[]byte("line1\nline2\nline3"),
		0644,
	)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	defer os.RemoveAll("banners")

	lines, err := Load("test.txt")
	if err != nil {
		t.Fatalf("expected nil, got error: %v", err)
	}

	if len(lines) != 3 {
		t.Errorf("expected 3 lines, got %d", len(lines))
	}

	_, err = Load("banana.txt")
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}