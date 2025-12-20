package complexity

import (
	"fmt"
	"go/ast"

	"github.com/serenitysz/serenity/internal/rules"
)

func CheckMaxFuncLinesNode(runner *rules.Runner) []rules.Issue {
	if runner.Cfg.Linter.Rules == nil {
		return nil
	}

	complexity := runner.Cfg.Linter.Rules.Complexity
	if complexity == nil {
		return nil
	}

	if complexity.Use != nil && !*complexity.Use {
		return nil
	}

	var limit int16 = 20
	if complexity.MaxFuncLines != nil && complexity.MaxFuncLines.Max != nil {
		limit = int16(*complexity.MaxFuncLines.Max)
	}

	fn, ok := runner.Node.(*ast.FuncDecl)
	if !ok || fn.Body == nil {
		return nil
	}

	start := runner.Fset.Position(fn.Pos()).Line
	end := runner.Fset.Position(fn.End()).Line
	lines := end - start + 1

	if int16(lines) <= limit {
		return nil
	}

	return []rules.Issue{{
		Pos:     runner.Fset.Position(fn.Pos()),
		Message: fmt.Sprintf("functions exceed the maximum line limit of %d (actual: %d)", limit, lines),
	}}
}
