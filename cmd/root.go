package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "loganalyzer",
    Short: "Analyse distribuée de logs en Go",
    Long:  "loganalyzer est un outil CLI pour analyser des fichiers de logs en parallèle.",
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func init() {
   //afaire
}
