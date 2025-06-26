package reporter

import (
	"encoding/json"
	"github.com/21Hmzz/loganalyzer/internal/analyzer"
	"os"
	"path/filepath"
)

func Write(path string, results []analyzer.Result) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
