package correctness

import (
	"go/ast"
	"go/token"

	"github.com/serenitysz/serenity/internal/rules"
)

type BooleanLiteralExpressionsRule struct{}

func (r *BooleanLiteralExpressionsRule) Name() string {
	return "boolean-literal-expressions"
}

func (r *BooleanLiteralExpressionsRule) Targets() []ast.Node {
	return []ast.Node{(*ast.BinaryExpr)(nil)}
}

func (r *BooleanLiteralExpressionsRule) Run(runner *rules.Runner, node ast.Node) {
	if runner.ShouldStop != nil && runner.ShouldStop() {
		return
	}

	if max := runner.Cfg.GetMaxIssues(); max > 0 && *runner.IssuesCount >= max {
		return
	}

	correctness := runner.Cfg.Linter.Rules.Correctness
	if correctness == nil || !correctness.Use || correctness.BoolLiteralExpressions == nil {
		return
	}

	expr := node.(*ast.BinaryExpr)

	if expr.Op != token.EQL && expr.Op != token.NEQ && !hasBoolLiteral(expr.X) && !hasBoolLiteral(expr.Y) {
		return
	}

	*runner.IssuesCount++

	*runner.Issues = append(*runner.Issues, rules.Issue{
		ID:       rules.BoolLiteralExpressionsID,
		Pos:      runner.Fset.Position(expr.Pos()),
		Severity: rules.ParseSeverity(correctness.BoolLiteralExpressions.Severity),
	})
}

func hasBoolLiteral(expr ast.Expr) bool {
	ident, ok := expr.(*ast.Ident)

	return ok && (ident.Name == "true" || ident.Name == "false")
}
