package complexity

import (
	"github.com/serenitysz/serenity/internal/rules"
)

func CheckAlwaysPreferConstNode(runner *rules.Runner) []rules.Issue {
	complexity := runner.Cfg.Linter.Rules.Complexity

	if complexity != nil && complexity.Use != nil && !*complexity.Use {
		return nil
	}

	// vb, ok := runner.Node.(*ast.InterfaceType)

	return nil
}
