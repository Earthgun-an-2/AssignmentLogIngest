package processor

type APILog struct {
	Timestamp string
	Message   string
	Level     string
}

type StructuredLog struct {
	ISO8601Time string
	Message     string
	Severity    string
}

// TransformLog standardizes the log format
func TransformLog(input APILog) StructuredLog {
	return StructuredLog{
		ISO8601Time: input.Timestamp, // assume already ISO8601 for simplicity
		Message:     input.Message,
		Severity:    input.Level,
	}
}
