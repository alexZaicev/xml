package parser

import (
	xmlerrors "github.com/alexZaicev/xml/go/internal/errors"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type errorListener struct {
	antlr.DefaultErrorListener
	errorSlice []error
}

func newErrorListener() *errorListener {
	return &errorListener{}
}

func (l *errorListener) Errors() []error {
	return l.errorSlice
}

func (l *errorListener) SyntaxError(
	recognizer antlr.Recognizer,
	offendingSymbol interface{},
	line, column int,
	msg string,
	e antlr.RecognitionException,
) {
	l.errorSlice = append(l.errorSlice, xmlerrors.NewSyntaxError(msg, line, column))
}
