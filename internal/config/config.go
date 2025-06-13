package config

import (
	"encoding/json"
	"io/ioutil"
)

type LogEntry struct {
    ID   string `json:"id"`
    Path string `json:"path"`
    Type string `json:"type"`
}

func Load(path string) ([]LogEntry, error) {
    data, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, err
    }
    var entries []LogEntry
    if err := json.Unmarshal(data, &entries); err != nil {
        return nil, err
    }
    return entries, nil
}
