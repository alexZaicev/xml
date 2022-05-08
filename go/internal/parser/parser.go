package parser

import "github.com/antlr/antlr4/runtime/Go/antlr"

type Parser interface {
	Validate() error
	Parse(antlr.ParseTreeListener) error
}
