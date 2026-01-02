package cmd

import (
	"cmp"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/serenitysz/serenity/internal/config"
	"github.com/serenitysz/serenity/internal/prompts"
	"github.com/serenitysz/serenity/internal/rules"
	"github.com/spf13/cobra"
)

var (
	interactive bool
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a new Serenity project by creating a serenity.json file",
	RunE: func(cmd *cobra.Command, args []string) error {
		if interactive {
			return createSerenityInteractive()
		}

		return createSerenity()
	},
}

func init() {
	initCmd.Flags().BoolVarP(&interactive, "interactive", "i", false, "Use an interactive mode to init Serenity project")

	rootCmd.AddCommand(initCmd)
}

func createSerenity() error {
	path, err := config.GetConfigFilePath()

	if err != nil {
		return err
	}

	exists, err := config.CheckHasConfigFile(path)

	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("config file already exists at %s", path)
	}

	autofix := false
	cfg := config.GenDefaultConfig(&autofix)

	return createConfig(path, cfg)
}

func createSerenityInteractive() error {
	defaultPath := cmp.Or(os.Getenv("SERENITY_CONFIG_PATH"), "serenity.json")

	path, err := prompts.Input(
		"Which configuration path location do you want to use",
		defaultPath,
	)

	if err != nil {
		return err
	}

	if path == "" {
		return errors.New("config path cannot be empty")
	}

	if !strings.HasSuffix(path, ".json") {
		return errors.New("config file must have .json extension")
	}

	exists, err := config.CheckHasConfigFile(path)

	if err != nil {
		return err
	}

	if exists {
		overwrite, err := prompts.Confirm("Config file already exists. Overwrite?")

		if err != nil {
			return err
		}

		if !overwrite {
			return nil
		}
	}

	strict, err := prompts.Confirm("Enable strict mode? (recommended)")

	if err != nil {
		return err
	}

	autofix, err := prompts.Confirm("Enable autofix when possible?")

	if err != nil {
		return err
	}

	cfg := config.GenDefaultConfig(&autofix)

	if strict {
		cfg = config.GenStrictDefaultConfig(&autofix)
	}

	return createConfig(path, cfg)
}

func createConfig(path string, cfg *rules.LinterOptions) error {
	fmt.Print("Creating serenity.json config file...")

	if err := config.CreateConfigFile(cfg, path); err != nil {
		return err
	}

	defaultSerenityPath := "serenity.json"

	if path != defaultSerenityPath {
		if err := os.Setenv("SERENITY_CONFIG_PATH", path); err != nil {
			return err
		}

		fmt.Printf("\nSetting SERENITY_CONFIG_PATH as \"%s\" in env file", path)
	}

	fmt.Println("\nConfig file created successfully.")

	return nil
}
