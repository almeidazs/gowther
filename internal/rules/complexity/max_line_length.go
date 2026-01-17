package complexity

import (
	"go/ast"

	"github.com/serenitysz/serenity/internal/rules"
)

type CheckMaxLineLengthRule struct{}

func (c *CheckMaxLineLengthRule) Name() string {
	return "max-line-length"
}

func (c *CheckMaxLineLengthRule) Targets() []ast.Node {
	return []ast.Node{(*ast.File)(nil)}
}

func (c *CheckMaxLineLengthRule) Run(runner *rules.Runner, node ast.Node) {
	if runner.ShouldStop != nil && runner.ShouldStop() {
		return
	}

	if max := runner.Cfg.GetMaxIssues(); max > 0 && *runner.IssuesCount >= max {
		return
	}

	complexity := runner.Cfg.Linter.Rules.Complexity

	if complexity == nil || !complexity.Use {
		return
	}

	ruleConfig := complexity.MaxLineLength

	if ruleConfig == nil {
		return
	}

	tokFile := runner.Fset.File(node.Pos())
	if tokFile == nil {
		return
	}

	var limit int = 80
	if ruleConfig.Max != nil {
		limit = int(*ruleConfig.Max)
	}

	offsets := tokFile.Lines()
	fSize := tokFile.Size()
	lineCount := len(offsets)

	maxIssues := runner.Cfg.GetMaxIssues()
	severity := rules.ParseSeverity(ruleConfig.Severity)

	for i := range lineCount {
		if maxIssues > 0 && *runner.IssuesCount >= maxIssues {
			return
		}

		var lineLen int

		isLastLine := i == lineCount-1
		if isLastLine {
			lineLen = fSize - offsets[i]
			if lineLen > 0 {
				lineLen-- // assume trailing newline
			}
		} else {
			lineLen = offsets[i+1] - offsets[i] - 1
		}

		if lineLen > limit {
			*runner.IssuesCount++
			*runner.Issues = append(*runner.Issues, rules.Issue{
				ID:       rules.MaxLineLengthID,
				Severity: severity,
				Pos:      tokFile.Position(tokFile.LineStart(i + 1)),
				ArgInt1:  limit,
				ArgInt2:  lineLen,
			})
		}
	}
}
