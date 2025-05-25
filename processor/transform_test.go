package processor

import (
	"testing"
)

func TestTransformLog(t *testing.T) {
	input := APILog{
		Timestamp: "2024-01-01T15:04:05Z",
		Message:   "API request failed",
		Level:     "ERROR",
	}

	expected := StructuredLog{
		ISO8601Time: "2024-01-01T15:04:05Z",
		Message:     "API request failed",
		Severity:    "ERROR",
	}

	result := TransformLog(input)

	if result != expected {
		t.Errorf("TransformLog() = %v; want %v", result, expected)
	}
}
