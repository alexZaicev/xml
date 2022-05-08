package errors

import "fmt"

type SyntaxError struct {
	baseError
	Line, Column int
}

func NewSyntaxError(msg string, line, column int) *SyntaxError {
	return &SyntaxError{
		baseError: newBaseError(
			fmt.Sprintf("syntax error: %s", msg),
			nil,
		),
		Line:   line,
		Column: column,
	}
}

type ParserGrammarError struct {
	baseError
	LexerErrors  []error
	ParserErrors []error
}

func NewParserGrammarError(msg string, lexerErrors, parserErrors []error) *ParserGrammarError {
	return &ParserGrammarError{
		baseError: newBaseError(
			fmt.Sprintf("parser grammar error: %s", msg),
			nil,
		),
		LexerErrors:  lexerErrors,
		ParserErrors: parserErrors,
	}
}
