package cmd

import (
	"fmt"

	"github.com/serenitysz/serenity/internal/config"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a new Serenity project by creating a serenity.json file",
	RunE: func(cmd *cobra.Command, args []string) error {
		return createSerenity()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func createSerenity() error {
	path, err := config.GetConfigFilePath()

	if err != nil {
		return err
	}

	exists, err := config.CheckHasConfigFile(path)

	if exists || err != nil {
		return nil
	}

	fmt.Println("Creating default serenity.json config file...")

	cfg := config.GenDefaultConfig()

	if err := config.CreateConfigFile(cfg, path); err != nil {
		return err
	}

	fmt.Println("Config file created successfully.")

	return nil
}
