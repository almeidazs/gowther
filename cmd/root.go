package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "gowther",
	Short: "A modern linter for Go",
	Long:  "Fast, opinionated linter and formatter for Go with auto-fix capabilities", // TODO: adicionar uma desc mais detalhada com a usage
}

func Exec() {
	cobra.CheckErr(rootCmd.Execute())
}
