package main

import (
	"log"

	"github.com/ingestservice/fetcher"
	"github.com/ingestservice/processor"
	"github.com/ingestservice/storage"
)

func main() {
	s3 := storage.NewS3Client()
	// ticker := time.NewTicker(10 * time.Second)

	// for {
	// 	select {
	// 	case <-ticker.C:
	log.Println("Fetching logs...")
	entries, err := fetcher.FetchLogs()
	if err != nil {
		log.Println("Error fetching logs:", err)
	}

	for _, entry := range entries {
		proc := processor.Transform(entry)
		err := s3.Store(proc)
		if err != nil {
			log.Println("Error storing log:", err)
		}
	}
	// }
	// }
}
