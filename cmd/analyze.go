package cmd

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/21Hmzz/loganalyzer/internal/analyzer"
	"github.com/21Hmzz/loganalyzer/internal/config"
	"github.com/21Hmzz/loganalyzer/internal/reporter"
	"github.com/spf13/cobra"
)

var cfgPath string
var outPath string

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyse les logs en parallèle",
	RunE: func(cmd *cobra.Command, args []string) error {
		entries, err := config.Load(cfgPath)
		if err != nil {
			return err
		}
		results := analyzer.Run(entries)
		for _, r := range results {
			fmt.Printf("%s | %s | %s | %s\n", r.LogID, r.FilePath, r.Status, r.Message)
			if r.ErrorDetails != "" {
				fmt.Printf("  → %s\n", r.ErrorDetails)
			}
		}

		//pour pas que ca crer un fichier meme sans --output
		if outPath != "" {
			date := time.Now().Format("060102")
			dir := filepath.Dir(outPath)
			base := filepath.Base(outPath)
			finalPath := filepath.Join(dir, fmt.Sprintf("%s_%s", date, base))

			if err := reporter.Write(finalPath, results); err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	analyzeCmd.Flags().StringVarP(&cfgPath, "config", "c", "", "Chemin vers le JSON de config (requis)")
	analyzeCmd.Flags().StringVarP(&outPath, "output", "o", "", "Chemin du rapport JSON")
	analyzeCmd.MarkFlagRequired("config")
	rootCmd.AddCommand(analyzeCmd)
}
