package util

import "testing"

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
