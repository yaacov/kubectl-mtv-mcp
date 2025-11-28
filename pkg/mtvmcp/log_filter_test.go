package mtvmcp

import (
	"strings"
	"testing"
)

func TestFilterLogs(t *testing.T) {
	testLogs := `2024-01-15 10:00:00 INFO: Starting application
2024-01-15 10:00:01 INFO: Loading configuration
2024-01-15 10:00:02 ERROR: Failed to connect to database
2024-01-15 10:00:03 INFO: Retrying connection
2024-01-15 10:00:04 ERROR: Connection timeout
2024-01-15 10:00:05 INFO: Using fallback configuration
2024-01-15 10:00:06 WARNING: Performance degraded
2024-01-15 10:00:07 INFO: System operational`

	tests := []struct {
		name           string
		filter         LogFilter
		expectedLines  int
		expectedGrep   int
		shouldTruncate bool
	}{
		{
			name:          "No filtering",
			filter:        LogFilter{},
			expectedLines: 8,
		},
		{
			name: "Grep for ERROR",
			filter: LogFilter{
				GrepPattern: "ERROR",
			},
			expectedLines: 2,
			expectedGrep:  2,
		},
		{
			name: "Grep for ERROR|WARNING (case insensitive)",
			filter: LogFilter{
				GrepPattern:    "error|warning",
				GrepIgnoreCase: true,
			},
			expectedLines: 3,
			expectedGrep:  3,
		},
		{
			name: "Tail 3 lines",
			filter: LogFilter{
				TailLines: 3,
			},
			expectedLines: 3,
		},
		{
			name: "Head 3 lines",
			filter: LogFilter{
				HeadLines: 3,
			},
			expectedLines: 3,
		},
		{
			name: "Grep with context",
			filter: LogFilter{
				GrepPattern: "ERROR",
				GrepContext: 1,
			},
			expectedLines: 5, // 2 errors + 1 line before each + 1 line after each, but overlapping
			expectedGrep:  2,
		},
		{
			name: "Max lines limit",
			filter: LogFilter{
				MaxLines: 3,
			},
			expectedLines:  3,
			shouldTruncate: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filtered, metadata := FilterLogs(testLogs, tt.filter)

			lines := strings.Split(strings.TrimSpace(filtered), "\n")
			actualLines := len(lines)

			if filtered == "" {
				actualLines = 0
			}

			if actualLines != tt.expectedLines {
				t.Errorf("Expected %d lines, got %d", tt.expectedLines, actualLines)
			}

			if tt.expectedGrep > 0 && metadata.GrepMatches != tt.expectedGrep {
				t.Errorf("Expected %d grep matches, got %d", tt.expectedGrep, metadata.GrepMatches)
			}

			if tt.shouldTruncate && !metadata.Truncated {
				t.Errorf("Expected truncation but got none")
			}

			if metadata.TotalLines != 8 {
				t.Errorf("Expected total lines to be 8, got %d", metadata.TotalLines)
			}
		})
	}
}

func TestGrepInvert(t *testing.T) {
	testLogs := `INFO: line 1
ERROR: line 2
INFO: line 3
ERROR: line 4`

	filter := LogFilter{
		GrepPattern: "ERROR",
		GrepInvert:  true,
	}

	filtered, metadata := FilterLogs(testLogs, filter)
	lines := strings.Split(strings.TrimSpace(filtered), "\n")

	// Should get only INFO lines (2 lines)
	if len(lines) != 2 {
		t.Errorf("Expected 2 lines with inverted grep, got %d", len(lines))
	}

	if !strings.Contains(filtered, "INFO") {
		t.Errorf("Expected INFO lines in output")
	}

	if strings.Contains(filtered, "ERROR") {
		t.Errorf("Did not expect ERROR lines in inverted grep output")
	}

	// Invert counts actual non-matches in the original pattern
	if metadata.GrepMatches != 2 {
		t.Errorf("Expected 2 grep matches (ERROR lines), got %d", metadata.GrepMatches)
	}
}

func TestMaxBytes(t *testing.T) {
	// Create a log with known size
	testLogs := strings.Repeat("1234567890", 10) // 100 bytes

	filter := LogFilter{
		MaxBytes: 50,
	}

	filtered, metadata := FilterLogs(testLogs, filter)

	if len(filtered) > 50 {
		t.Errorf("Expected max 50 bytes, got %d", len(filtered))
	}

	if !metadata.Truncated {
		t.Errorf("Expected truncation flag to be set")
	}

	if metadata.TruncationReason != "max_bytes_exceeded" {
		t.Errorf("Expected truncation reason 'max_bytes_exceeded', got '%s'", metadata.TruncationReason)
	}
}
