package rules

type Config struct {
	Schema        string                  `json:"$schema"`
	Naming        *NamingRuleGroup        `json:"naming"`
	Complexity    *ComplexityRuleGroup    `json:"complexity"`
	BestPractices *BestPracticesRuleGroup `json:"bestPractices"`
	ErrorHandling *ErrorHandlingRuleGroup `json:"errorHandling"`
	Imports       *ImportsRuleGroup       `json:"imports"`
	Exclude       []string                `json:"exclude"`
}

type NamingRuleGroup struct {
	Enabled bool         `json:"enabled"`
	Rules   *NamingRules `json:"rules"`
}

type ComplexityRuleGroup struct {
	Enabled bool             `json:"enabled"`
	Rules   *ComplexityRules `json:"rules"`
}

type BestPracticesRuleGroup struct {
	Enabled bool                `json:"enabled"`
	Rules   *BestPracticesRules `json:"rules"`
}

type ErrorHandlingRuleGroup struct {
	Enabled bool                `json:"enabled"`
	Rules   *ErrorHandlingRules `json:"rules"`
}

type ImportsRuleGroup struct {
	Enabled bool          `json:"enabled"`
	Rules   *ImportsRules `json:"rules"`
}

type NamingRules struct {
	ExportedIdentifiers   *PatternRule   `json:"exported_identifiers,omitempty"`
	UnexportedIdentifiers *PatternRule   `json:"unexported_identifiers,omitempty"`
	ReceiverNames         *MaxLengthRule `json:"receiver_names,omitempty"`
}

type ComplexityRules struct {
	CyclomaticComplexity *ThresholdRule `json:"cyclomatic_complexity,omitempty"`
	MaxFunctionLines     *ThresholdRule `json:"max_function_lines,omitempty"`
	MaxNestingDepth      *ThresholdRule `json:"max_nesting_depth,omitempty"`
}

type BestPracticesRules struct {
	NoBareReturns     *DescriptionRule `json:"no_bare_returns,omitempty"`
	ContextFirstParam *DescriptionRule `json:"context_first_param,omitempty"`
	NoDeferInLoop     *DescriptionRule `json:"no_defer_in_loop,omitempty"`
}

type ErrorHandlingRules struct {
	ErrorWrapping     *ErrorWrappingRule     `json:"error_wrapping,omitempty"`
	ErrorStringFormat *ErrorStringFormatRule `json:"error_string_format,omitempty"`
	NoErrorShadowing  *DescriptionRule       `json:"no_error_shadowing,omitempty"`
}

type ImportsRules struct {
	NoDotImports       *SeverityRule           `json:"no_dot_imports,omitempty"`
	DisallowedPackages *DisallowedPackagesRule `json:"disallowed_packages,omitempty"`
}

type SeverityRule struct {
	Severity string `json:"severity"`
}

type DescriptionRule struct {
	Severity    string `json:"severity"`
	Description string `json:"description"`
}

type PatternRule struct {
	Severity    string `json:"severity"`
	Description string `json:"description"`
	Pattern     string `json:"pattern"`
}

type MaxLengthRule struct {
	Severity    string `json:"severity"`
	Description string `json:"description"`
	MaxLength   int    `json:"maxLength"`
}

type ThresholdRule struct {
	Severity  string `json:"severity"`
	Threshold int    `json:"threshold"`
}

type ErrorWrappingRule struct {
	Severity    string `json:"severity"`
	Description string `json:"description"`
	RequireFmtW bool   `json:"requireFmtW"`
}

type ErrorStringFormatRule struct {
	Severity      string `json:"severity"`
	Case          string `json:"case"`
	NoPunctuation bool   `json:"noPunctuation"`
}

type DisallowedPackagesRule struct {
	Severity string   `json:"severity"`
	Packages []string `json:"packages"`
}
