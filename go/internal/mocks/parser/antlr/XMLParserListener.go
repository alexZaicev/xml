// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	antlr "github.com/antlr/antlr4/runtime/Go/antlr"
	mock "github.com/stretchr/testify/mock"

	parser "github.com/alexZaicev/xml/go/internal/parser/antlr"
)

// XMLParserListener is an autogenerated mock type for the XMLParserListener type
type XMLParserListener struct {
	mock.Mock
}

// EnterAttribute provides a mock function with given fields: c
func (_m *XMLParserListener) EnterAttribute(c *parser.AttributeContext) {
	_m.Called(c)
}

// EnterChardata provides a mock function with given fields: c
func (_m *XMLParserListener) EnterChardata(c *parser.ChardataContext) {
	_m.Called(c)
}

// EnterContent provides a mock function with given fields: c
func (_m *XMLParserListener) EnterContent(c *parser.ContentContext) {
	_m.Called(c)
}

// EnterDocument provides a mock function with given fields: c
func (_m *XMLParserListener) EnterDocument(c *parser.DocumentContext) {
	_m.Called(c)
}

// EnterElement provides a mock function with given fields: c
func (_m *XMLParserListener) EnterElement(c *parser.ElementContext) {
	_m.Called(c)
}

// EnterEveryRule provides a mock function with given fields: ctx
func (_m *XMLParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	_m.Called(ctx)
}

// EnterMisc provides a mock function with given fields: c
func (_m *XMLParserListener) EnterMisc(c *parser.MiscContext) {
	_m.Called(c)
}

// EnterProlog provides a mock function with given fields: c
func (_m *XMLParserListener) EnterProlog(c *parser.PrologContext) {
	_m.Called(c)
}

// EnterReference provides a mock function with given fields: c
func (_m *XMLParserListener) EnterReference(c *parser.ReferenceContext) {
	_m.Called(c)
}

// ExitAttribute provides a mock function with given fields: c
func (_m *XMLParserListener) ExitAttribute(c *parser.AttributeContext) {
	_m.Called(c)
}

// ExitChardata provides a mock function with given fields: c
func (_m *XMLParserListener) ExitChardata(c *parser.ChardataContext) {
	_m.Called(c)
}

// ExitContent provides a mock function with given fields: c
func (_m *XMLParserListener) ExitContent(c *parser.ContentContext) {
	_m.Called(c)
}

// ExitDocument provides a mock function with given fields: c
func (_m *XMLParserListener) ExitDocument(c *parser.DocumentContext) {
	_m.Called(c)
}

// ExitElement provides a mock function with given fields: c
func (_m *XMLParserListener) ExitElement(c *parser.ElementContext) {
	_m.Called(c)
}

// ExitEveryRule provides a mock function with given fields: ctx
func (_m *XMLParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	_m.Called(ctx)
}

// ExitMisc provides a mock function with given fields: c
func (_m *XMLParserListener) ExitMisc(c *parser.MiscContext) {
	_m.Called(c)
}

// ExitProlog provides a mock function with given fields: c
func (_m *XMLParserListener) ExitProlog(c *parser.PrologContext) {
	_m.Called(c)
}

// ExitReference provides a mock function with given fields: c
func (_m *XMLParserListener) ExitReference(c *parser.ReferenceContext) {
	_m.Called(c)
}

// VisitErrorNode provides a mock function with given fields: node
func (_m *XMLParserListener) VisitErrorNode(node antlr.ErrorNode) {
	_m.Called(node)
}

// VisitTerminal provides a mock function with given fields: node
func (_m *XMLParserListener) VisitTerminal(node antlr.TerminalNode) {
	_m.Called(node)
}
