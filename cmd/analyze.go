package cmd

import (
	"fmt"

	"github.com/21Hmzz/loganalyzer/internal/config"
	"github.com/spf13/cobra"
)

var cfgPath string

var analyzeCmd = &cobra.Command{
    Use:   "analyze",
    Short: "Test : charge et affiche la config JSON",
    RunE: func(cmd *cobra.Command, args []string) error {
        entries, err := config.Load(cfgPath)
        if err != nil {
            return err
        }
        for _, e := range entries {
            fmt.Printf("ID=%s, Path=%s, Type=%s\n", e.ID, e.Path, e.Type)
        }
        return nil
    },
}

func init() {
    analyzeCmd.Flags().StringVarP(&cfgPath, "config", "c", "", "Chemin vesr le JSON de config")
    analyzeCmd.MarkFlagRequired("config")
    rootCmd.AddCommand(analyzeCmd)
}
