package cmd

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"github.com/21Hmzz/loganalyzer/internal/config"
	"github.com/spf13/cobra"
)

var addID, addPath, addType, cfgFile string

var addLogCmd = &cobra.Command{
	Use:   "add-log",
	Short: "Ajoute une entre dans un config existant",
	RunE: func(cmd *cobra.Command, args []string) error {
		list, err := config.Load(cfgFile)
		if err != nil {
			return err
		}
		for _, e := range list {
			if e.ID == addID {
				return errors.New("ID deja prst dans la config")
			}
		}
		list = append(list, config.LogEntry{ID: addID, Path: addPath, Type: addType})
		data, _ := json.MarshalIndent(list, "", "  ")
		if err := os.MkdirAll(filepath.Dir(cfgFile), 0755); err != nil {
			return err
		}
		return os.WriteFile(cfgFile, data, 0644)
	},
}

func init() {
	addLogCmd.Flags().StringVar(&addID, "id", "", "id du log")
	addLogCmd.Flags().StringVar(&addPath, "path", "", "Chemin du fichier de log")
	addLogCmd.Flags().StringVar(&addType, "type", "", "Type de log")
	addLogCmd.Flags().StringVar(&cfgFile, "file", "", "Fichier config")
	addLogCmd.MarkFlagRequired("id")
	addLogCmd.MarkFlagRequired("path")
	addLogCmd.MarkFlagRequired("type")
	addLogCmd.MarkFlagRequired("file")
	rootCmd.AddCommand(addLogCmd)
}
