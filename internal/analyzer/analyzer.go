package analyzer

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/21Hmzz/loganalyzer/internal/config"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func analyzeOne(e config.LogEntry) Result {
	// Vérif existence
	if info, err := os.Stat(e.Path); err != nil {
		if os.IsNotExist(err) {
			return Result{e.ID, e.Path, "FAILED", "Fichier introuvable", (&FileNotFoundErr{e.Path}).Error()}
		}
		return Result{e.ID, e.Path, "FAILED", "Erreur accès fichier", err.Error()}
	} else if info.IsDir() {
		return Result{e.ID, e.Path, "FAILED", "C’est un répertoire", fmt.Sprintf("%s est un dossier", e.Path)}
	}

	time.Sleep(time.Duration(50+rand.Intn(151)) * time.Millisecond)

	if rand.Float64() < 0.1 {
		perr := &ParseErr{e.ID}
		return Result{e.ID, e.Path, "FAILED", "Erreur de parsing", perr.Error()}
	}

	return Result{e.ID, e.Path, "OK", "Analyse réussie", ""}
}

type Result struct {
	LogID        string
	FilePath     string
	Status       string
	Message      string
	ErrorDetails string
}

func Run(entries []config.LogEntry, statusFilter string) []Result {
	ch := make(chan Result, len(entries))
	var wg sync.WaitGroup

	for _, e := range entries {
		wg.Add(1)
		go func(ent config.LogEntry) {
			defer wg.Done()
			ch <- analyzeOne(ent)
		}(e)
	}

	wg.Wait()
	close(ch)

	var results []Result
	for r := range ch {
		if statusFilter != "" && r.Status != statusFilter {
			continue
		}
		results = append(results, r)
	}
	return results
}
