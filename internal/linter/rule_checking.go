package linter

import (
	"github.com/serenitysz/serenity/internal/rules"
	"github.com/serenitysz/serenity/internal/rules/bestpractices"
	"github.com/serenitysz/serenity/internal/rules/complexity"
)

func CheckAvailableRules(runner *rules.Runner) {
	rules := runner.Cfg.Linter.Rules
	activeRules := make([]func(*rules.Runner), 0, 20)

	if rules.BestPractices.MaxParams.Use != nil && *rules.BestPractices.MaxParams.Use {
		activeRules = append(activeRules, bestpractices.CheckMaxParamsNode)
	}

	if rules.Complexity.MaxFuncLines.Use != nil && *rules.Complexity.MaxFuncLines.Use {
		activeRules = append(activeRules, complexity.CheckMaxFuncLinesNode)
	}
}
