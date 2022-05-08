package parser

import (
	xmlerrors "github.com/alexZaicev/xml/go/internal/errors"
	genparser "github.com/alexZaicev/xml/go/internal/parser/antlr"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type XMLParser struct {
	parser *genparser.XMLParser

	lexerErrorListener  *errorListener
	parserErrorListener *errorListener
}

func NewXMLParser(input string) *XMLParser {
	lexerErrorListener := newErrorListener()
	lexer := genparser.NewXMLLexer(antlr.NewInputStream(input))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexerErrorListener)

	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parserErrorListener := newErrorListener()
	parser := genparser.NewXMLParser(tokenStream)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(parserErrorListener)

	return &XMLParser{
		parser:              parser,
		lexerErrorListener:  lexerErrorListener,
		parserErrorListener: parserErrorListener,
	}
}

func (p *XMLParser) Validate() error {
	p.parser.Document()
	lexerErrors := p.lexerErrorListener.Errors()
	parserErrors := p.parserErrorListener.Errors()
	if len(lexerErrors) > 0 || len(parserErrors) > 0 {
		return xmlerrors.NewParserGrammarError("XML document contains syntax errors", lexerErrors, parserErrors)
	}
	return nil
}

func (p *XMLParser) Parse(listener antlr.ParseTreeListener) error {
	if listener == nil {
		return xmlerrors.NewIllegalArgumentError("listener", "cannot be nil")
	}
	antlr.ParseTreeWalkerDefault.Walk(listener, p.parser.Document())
	lexerErrors := p.lexerErrorListener.Errors()
	parserErrors := p.parserErrorListener.Errors()
	if len(lexerErrors) > 0 || len(parserErrors) > 0 {
		return xmlerrors.NewParserGrammarError("XML document contains syntax errors", lexerErrors, parserErrors)
	}
	return nil
}
