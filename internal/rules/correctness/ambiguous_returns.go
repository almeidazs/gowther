package correctness

import (
	"go/ast"

	"github.com/serenitysz/serenity/internal/rules"
)

type AmbiguousReturnRule struct{}

func (r *AmbiguousReturnRule) Name() string {
	return "ambiguous-return"
}

func (r *AmbiguousReturnRule) Targets() []ast.Node {
	return []ast.Node{(*ast.FuncDecl)(nil)}
}

func (r *AmbiguousReturnRule) Run(runner *rules.Runner, node ast.Node) {
	if runner.ShouldStop != nil && runner.ShouldStop() {
		return
	}

	if max := runner.Cfg.GetMaxIssues(); max > 0 && *runner.IssuesCount >= max {
		return
	}

	correctness := runner.Cfg.Linter.Rules.Correctness

	if correctness == nil || !correctness.Use || correctness.AmbiguousReturns == nil {
		return
	}

	fn := node.(*ast.FuncDecl)

	if fn.Type.Results == nil {
		return
	}

	results := fn.Type.Results.List

	if len(results) < 2 {
		return
	}

	seen := make(map[string]int, len(results))
	hasUnnamed := false

	for _, field := range results {
		if len(field.Names) > 0 {
			return
		}

		hasUnnamed = true
		seen[typeKey(field.Type)]++
	}

	if !hasUnnamed {
		return
	}

	maxAllowed := 1

	if correctness.AmbiguousReturns.MaxUnnamedSameType != nil {
		maxAllowed = *correctness.AmbiguousReturns.MaxUnnamedSameType
	}

	for typ, count := range seen {
		if count > maxAllowed && !isAllowedDuplicateReturn(typ) {
			*runner.IssuesCount++
			*runner.Issues = append(*runner.Issues, rules.Issue{
				ArgInt1:  count,
				ArgInt2:  maxAllowed,
				ID:       rules.AmbiguousReturnID,
				Pos:      runner.Fset.Position(fn.Type.Results.Pos()),
				Severity: rules.ParseSeverity(correctness.AmbiguousReturns.Severity),
			})

			return
		}
	}
}

func typeKey(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return t.X.(*ast.Ident).Name + "." + t.Sel.Name
	case *ast.StarExpr:
		return "*" + typeKey(t.X)
	default:
		return ""
	}
}

func isAllowedDuplicateReturn(typ string) bool {
	return typ == "error"
}
