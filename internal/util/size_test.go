package util

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFormatSize(t *testing.T) {
	tests := []struct {
		bytes    int64
		expected string
	}{
		{0, "0 B"},
		{1023, "1023 B"},
		{1024, "1.0 KB"},
		{1536, "1.5 KB"},
		{1048576, "1.0 MB"},
		{1073741824, "1.0 GB"},
	}

	for _, tt := range tests {
		result := FormatSize(tt.bytes)
		if result != tt.expected {
			t.Errorf("FormatSize(%d) = %s, want %s", tt.bytes, result, tt.expected)
		}
	}
}

func TestCalculateSize(t *testing.T) {
	tmpDir := t.TempDir()

	file1 := filepath.Join(tmpDir, "file1.txt")
	if err := os.WriteFile(file1, make([]byte, 1024), 0644); err != nil {
		t.Fatal(err)
	}

	file2 := filepath.Join(tmpDir, "file2.txt")
	if err := os.WriteFile(file2, make([]byte, 512), 0644); err != nil {
		t.Fatal(err)
	}

	size := CalculateSize(tmpDir)
	expected := int64(1024 + 512)

	if size != expected {
		t.Errorf("CalculateSize() = %d, want %d", size, expected)
	}
}
