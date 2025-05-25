package processor

import (
	"encoding/json"

	"github.com/ingestservice/fetcher"
)

func Transform(entry fetcher.LogEntry) []byte {
	entry.Title = "[PROCESSED] " + entry.Title
	result, _ := json.Marshal(entry)
	return result
}
