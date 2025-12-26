package linter

import (
	"reflect"

	"github.com/serenitysz/serenity/internal/rules"
	"github.com/serenitysz/serenity/internal/rules/bestpractices"
)

func GetActiveRulesMap(cfg *rules.LinterOptions) map[reflect.Type][]rules.Rule {
	activeRules := make(map[reflect.Type][]rules.Rule)
	const initialCap = 8

	register := func(r rules.Rule) {
		for _, target := range r.Targets() {
			t := rules.GetNodeType(target)
			if activeRules[t] == nil {
				activeRules[t] = make([]rules.Rule, 0, initialCap)
			}
			activeRules[t] = append(activeRules[t], r)
		}
	}

	r := cfg.Linter.Rules
	if bp := r.BestPractices; bp != nil && (bp.Use == nil || *bp.Use) {
		if bp.MaxParams != nil {
			register(&bestpractices.MaxParamsRule{})
		}
	}

	return activeRules
}
