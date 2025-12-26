package bestpractices

import (
	"go/ast"

	"github.com/serenitysz/serenity/internal/rules"
)

type ContextFirstRule struct{}

func (c *ContextFirstRule) Name() string {
	return "context-first-param"
}

func (c *ContextFirstRule) Targets() []ast.Node {
	return []ast.Node{(*ast.FuncDecl, *ast.SelectorExpr, *ast.StarExpr)(nil)}
}

func (c *ContextFirstRule) Run(runner *rules.Runner, node ast.Node) {
}

func isContextType(expr ast.Expr) bool {
	switch t := expr.(type) {
	case *ast.SelectorExpr:
		x, ok := t.X.(*ast.Ident)

		return ok && x.Name == "context" && t.Sel.Name == "Context"
	case *ast.StarExpr:
		if sel, ok := t.X.(*ast.SelectorExpr); ok {
			x, ok := sel.X.(*ast.Ident)

			return ok && x.Name == "context" && sel.Sel.Name == "Context"
		}
	}

	return false
}

func CheckContextFirstParamNode(runner *rules.Runner) {
	bestPractices := runner.Cfg.Linter.Rules.BestPractices

	if bestPractices == nil {
		return
	}

	if bestPractices.Use != nil && !*bestPractices.Use {
		return
	}

	contextFirst := bestPractices.UseContextInFirstParam
	if contextFirst == nil {
		return
	}

	fn, ok := runner.Node.(*ast.FuncDecl)
	if !ok {
		return
	}

	params := fn.Type.Params
	if params == nil || len(params.List) < 2 {
		return
	}

	maxIssues := rules.GetMaxIssues(runner.Cfg)

	for i := 1; i < len(params.List); i++ {
		p := params.List[i]

		if isContextType(p.Type) {
			if maxIssues > 0 && int16(len(*runner.Issues)) >= maxIssues {
				break
			}

			*runner.Issues = append(*runner.Issues, rules.Issue{
				ID:       rules.UseContextInFirstParamID,
				Pos:      runner.Fset.Position(p.Pos()),
				Severity: rules.ParseSeverity(contextFirst.Severity),
			})
		}
	}
}
