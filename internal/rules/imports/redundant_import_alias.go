package imports

import (
	"go/ast"
	"strings"

	"github.com/serenitysz/serenity/internal/rules"
)

type RedundantImportAliasRule struct{}

func (r *RedundantImportAliasRule) Name() string {
	return "redundant-import-alias"
}

func (r *RedundantImportAliasRule) Targets() []ast.Node {
	return []ast.Node{(*ast.ImportSpec)(nil)}
}

func (r *RedundantImportAliasRule) Run(runner *rules.Runner, node ast.Node) {
	if runner.ShouldStop != nil && runner.ShouldStop() {
		return
	}

	if max := runner.Cfg.GetMaxIssues(); max > 0 && *runner.IssuesCount >= max {
		return
	}

	imports := runner.Cfg.Linter.Rules.Imports

	if imports == nil || !imports.Use || imports.RedundantImportAlias == nil {
		return
	}

	spec := node.(*ast.ImportSpec)

	if spec.Name == nil {
		return
	}

	path := strings.Trim(spec.Path.Value, `"`)

	defaultName := defaultImportName(path)

	if spec.Name.Name == defaultName {
		*runner.IssuesCount++

		*runner.Issues = append(*runner.Issues, rules.Issue{
			ArgStr1:  defaultName,
			ID:       rules.RedundantImportAliasID,
			Pos:      runner.Fset.Position(spec.Name.Pos()),
			Severity: rules.ParseSeverity(imports.RedundantImportAlias.Severity),
		})

		if runner.Cfg.ShouldAutofix() {
			spec.Name = nil
			runner.Modified = true
		}
	}
}

func defaultImportName(path string) string {
	if i := strings.LastIndex(path, "/"); i >= 0 {
		return path[i+1:]
	}

	return path
}
