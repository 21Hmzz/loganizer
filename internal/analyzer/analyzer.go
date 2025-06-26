package analyzer

import "github.com/21Hmzz/loganalyzer/internal/config"

type Result struct {
	LogID        string
	FilePath     string
	Status       string
	Message      string
	ErrorDetails string
}

func Run(entries []config.LogEntry) []Result {
	// TODO
	return nil

}
