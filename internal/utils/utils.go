package utils

import (
	"fmt"

	"github.com/serenitysz/serenity/internal/rules"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
)

func FormatLog(issue rules.Issue, msg string) {
	var label, color string

	switch issue.Severity {
	case rules.SeverityError:
		label = "ERROR"
		color = colorRed
	case rules.SeverityWarn:
		label = "WARN"
		color = colorYellow
	case rules.SeverityInfo:
		label = "INFO"
		color = colorBlue
	default:
		label = "ISSUE"
		color = colorReset
	}

	fmt.Printf("%s%s:%d:%d: [%s] %s%s\n",
		color,
		issue.Pos.Filename,
		issue.Pos.Line,
		issue.Pos.Column,
		msg,
		label,
		colorReset,
	)
}
