package cronparser

import (
	"cron-parser/internal/parser"
)

type Parser interface {
	Parse(cronExpr string) (string, error)
}

func NewParser(cronExpr string) Parser {
	// Determine the type of the cron expression and return the appropriate parser
	return &StandardParser{}

}

type StandardParser struct{}

func (p *StandardParser) Parse(cronExpr string) (string, error) {
	return parser.ParseStdCronExpression(cronExpr)
}
