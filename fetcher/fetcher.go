package fetcher

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type LogEntry struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

func FetchLogs() ([]LogEntry, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var logs []LogEntry
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &logs)
	return logs, err
}
