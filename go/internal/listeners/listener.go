package listeners

import genparser "github.com/alexZaicev/xml/go/internal/parser/antlr"

// BaseListenerWithErrors is an abstraction for concrete listener implementation that adds helper
// functionality for recording errors.
type baseListenerWithErrors struct {
	errors []error
}

// Errors gets any errors that occurred during tree building.
func (l *baseListenerWithErrors) Errors() []error {
	return l.errors
}

// recordError internal function recording any error that occurred while building the node tree.
func (l *baseListenerWithErrors) recordError(err error) {
	l.errors = append(l.errors, err)
}

type BaseXMLListenerWithErrors struct {
	baseListenerWithErrors
	genparser.BaseXMLParserListener
}
