package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/serenitysz/serenity/internal/rules"
)

func ReadConfig(path string) (*rules.LinterOptions, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config rules.LinterOptions

	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config JSON: %w", err)
	}

	return &config, nil
}

func ApplyRecommended(cfg *rules.LinterOptions) {
	if cfg.Linter.Rules == nil || cfg.Linter.Rules.UseRecommended == nil || !*cfg.Linter.Rules.UseRecommended {
		return
	}

	rulesGroup := cfg.Linter.Rules

	if rulesGroup.Imports == nil {
		rulesGroup.Imports = &rules.ImportRulesGroup{}
	}
	if rulesGroup.Imports.NoDotImports == nil {
		rulesGroup.Imports.NoDotImports = &rules.LinterBaseRule{Severity: "error"}
	}

	if rulesGroup.BestPractices == nil {
		rulesGroup.BestPractices = &rules.BestPracticesRulesGroup{}
	}

	if rulesGroup.BestPractices.UseContextInFirstParam == nil {
		rulesGroup.BestPractices.UseContextInFirstParam = &rules.LinterBaseRule{Severity: "warn"}
	}
	if rulesGroup.BestPractices.MaxParams == nil {
		var q int8 = 5
		rulesGroup.BestPractices.MaxParams = &rules.MaxParams{Quantity: &q}
	}
}
