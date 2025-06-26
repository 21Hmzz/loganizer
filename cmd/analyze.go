package cmd

import (
	"fmt"
	"github.com/21Hmzz/loganalyzer/internal/analyzer"
	"github.com/21Hmzz/loganalyzer/internal/config"
	"github.com/spf13/cobra"
)

var cfgPath string

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
		return nil
	},
}

func init() {
	analyzeCmd.Flags().StringVarP(&cfgPath, "config", "c", "", "Chemin vers le JSON de config (requis)")
	analyzeCmd.MarkFlagRequired("config")
	rootCmd.AddCommand(analyzeCmd)
}
