package listeners

import (
	"strings"

	xmlerrors "github.com/alexZaicev/xml/go/internal/errors"
	genparser "github.com/alexZaicev/xml/go/internal/parser/antlr"
	"github.com/alexZaicev/xml/go/tree"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

const XMLNS = "xmlns"
const tagWithNamespaceSize = 2

// Prolog attributes
const (
	version    = "version"
	standalone = "standalone"
	encoding   = "encoding"
)

const (
	version10 = "1.0"
	version11 = "1.1"
)

const (
	standaloneYes = "yes"
	standaloneNo  = "no"
)

type XMLListener struct {
	BaseXMLListenerWithErrors

	currentNode     *tree.XMLNode
	prolog          *tree.Prolog
	prologProcessed bool
	namespaces      map[string]string
}

func NewXMLListener() *XMLListener {
	return &XMLListener{}
}

func (l *XMLListener) Document() *tree.XMLDocument {
	if l.currentNode == nil {
		l.recordError(xmlerrors.NewListenerError("current node was nil upon creating XML document"))
		return nil
	}
	if len(l.errors) > 0 {
		return nil
	}
	doc := tree.NewXMLDocument(
		l.prolog,
		l.currentNode,
	)
	if l.namespaces != nil && len(l.namespaces) > 0 {
		return doc.WithNamespaces(l.namespaces)
	}
	return doc
}

func (l *XMLListener) EnterDocument(ctx *genparser.DocumentContext) {
	// reset listener
	l.errors = make([]error, 0)
	l.currentNode = nil
	l.namespaces = make(map[string]string, 0)
	l.prologProcessed = true
}

func (l *XMLListener) EnterProlog(ctx *genparser.PrologContext) {
	// prolog flag indicates that the following XML attributes will be related to XML declaration.
	l.prologProcessed = false
	l.prolog = tree.NewProlog()
}

func (l *XMLListener) ExitProlog(ctx *genparser.PrologContext) {
	// prolog flag set to false will indicate that next read attributes are related to document elements.
	l.prologProcessed = true
}

func (l *XMLListener) EnterElement(ctx *genparser.ElementContext) {
	// initialize new document element
	if l.currentNode == nil {
		l.currentNode = tree.NewXMLRootNode()
	} else {
		newNode := tree.NewXMLChildNode(l.currentNode)
		l.currentNode = newNode
	}
	// set position metadata of the element
	if err := l.setPosition(ctx.BaseParserRuleContext); err != nil {
		l.recordError(err)
		return
	}
	// extract element name and validate if element tag is not self-closing
	var tagName string
	beginningToken, endingToken := ctx.GetBeginning(), ctx.GetEnding()
	if beginningToken == nil {
		// grammatically this case is not possible as this would be picked up by the XML parser
		l.recordError(
			xmlerrors.NewListenerErrorFromPosition(
				"element beginning tag token is nil",
				ctx.GetStart(),
			))
		return
	}
	if endingToken == nil {
		tagName = beginningToken.GetText()
	} else {
		tagStartName, tagEndName := beginningToken.GetText(), endingToken.GetText()
		if tagStartName != tagEndName {
			l.recordError(
				xmlerrors.NewListenerErrorFromPosition(
					"element beginning and ending tags do not match",
					ctx.GetStart(),
				))
			return
		}
		tagName = tagStartName
	}
	// split element name into element namespace and true element name.
	// If element would not have a namespace prefix returned prefix value would be empty.
	// If prefix and name are declared incorrectly an error will be returned
	name, prefix, err := l.splitTagNameAndNS(ctx.BaseParserRuleContext, tagName)
	if err != nil {
		l.recordError(err)
		return
	}
	l.currentNode.Name = name
	l.currentNode = l.currentNode.WithNSPrefix(prefix)
}

func (l *XMLListener) ExitElement(ctx *genparser.ElementContext) {
	if !l.validateCurrentNode(ctx.BaseParserRuleContext) {
		return
	}
	// If namespace is set on the element, validate that the namespace was registered in the document.
	if l.currentNode.Namespace != "" {
		_, ok := l.namespaces[l.currentNode.Namespace]
		if !ok {
			l.recordError(
				xmlerrors.NewListenerErrorFromPosition(
					"unknown element namespace",
					ctx.GetStart(),
				))
		}
	}
	// If current node is not the root element, add it to the parent element children and set current
	// node as it's parent
	if l.currentNode.Parent != nil {
		l.currentNode.Parent.Children = append(l.currentNode.Parent.Children, l.currentNode)
		l.currentNode = l.currentNode.Parent
	}
}

func (l *XMLListener) EnterChardata(ctx *genparser.ChardataContext) {
	if !l.validateCurrentNode(ctx.BaseParserRuleContext) {
		return
	}
	valueToken := ctx.GetValue()
	// character data value token can be nil when element does not contain any text data
	if valueToken != nil {
		l.currentNode.Content = valueToken.GetText()
	}
}

func (l *XMLListener) EnterAttribute(ctx *genparser.AttributeContext) {
	// if prolog processed flag is set to false, process attributes and set them to prolog struct
	if !l.prologProcessed {
		if l.prolog == nil {
			l.recordError(
				xmlerrors.NewListenerErrorFromPosition(
					"prolog was nil",
					ctx.GetStart(),
				))
			return
		}
		l.processPrologAttr(ctx)
		return
	}
	// continue processing attributes for element
	if !l.validateCurrentNode(ctx.BaseParserRuleContext) {
		return
	}
	l.processElementAttr(ctx)
}

func (l *XMLListener) validateCurrentNode(ctx *antlr.BaseParserRuleContext) bool {
	if l.currentNode == nil {
		l.recordError(xmlerrors.NewListenerErrorFromPosition("current node was nil", ctx.GetStart()))
		return false
	}
	return true
}

func (l *XMLListener) setPosition(ctx *antlr.BaseParserRuleContext) error {
	token := ctx.GetStart()
	if token == nil {
		return xmlerrors.NewListenerError("could not get position information from context")
	}
	line, column := ctx.GetStart().GetLine(), ctx.GetStart().GetColumn()
	l.currentNode = l.currentNode.WithPosition(line, column)
	return nil
}

func (l *XMLListener) splitTagNameAndNS(
	ctx *antlr.BaseParserRuleContext,
	name string,
) (tagName, prefix string, err error) {
	nameNSPair := strings.Split(name, ":")
	if len(nameNSPair) > tagWithNamespaceSize {
		return "", "", xmlerrors.NewListenerErrorFromPosition("invalid element name", ctx.GetStart())
	}
	if len(nameNSPair) == 1 {
		return nameNSPair[0], "", nil
	}
	return nameNSPair[1], nameNSPair[0], nil
}

func (l *XMLListener) processPrologAttr(ctx *genparser.AttributeContext) {
	// get attribute name and value pair
	attrName, attrValue, err := l.getAttNameValuePair(ctx)
	if err != nil {
		l.recordError(err)
		return
	}
	// set gathered attribute values to prolog
	switch attrName {
	case version:
		if attrValue == version10 || attrValue == version11 {
			l.prolog = l.prolog.WithVersion(attrValue)
		} else {
			l.recordError(xmlerrors.NewListenerErrorFromPosition(
				"invalid prolog version attribute specified",
				ctx.GetName(),
			))
		}
	case standalone:
		if attrValue == standaloneYes || attrValue == standaloneNo {
			l.prolog = l.prolog.WithStandalone(attrValue)
		} else {
			l.recordError(xmlerrors.NewListenerErrorFromPosition(
				"invalid prolog standalone attribute specified",
				ctx.GetName(),
			))
		}
	case encoding:
		l.prolog = l.prolog.WithEncoding(attrValue)
	default:
		l.recordError(xmlerrors.NewListenerErrorFromPosition(
			"unrecognized prolog attribute",
			ctx.GetName(),
		))
	}
}

func (l *XMLListener) processElementAttr(ctx *genparser.AttributeContext) {
	// get attribute name and value pair
	attrName, attrValue, err := l.getAttNameValuePair(ctx)
	if err != nil {
		l.recordError(err)
		return
	}
	// if attribute starts with XMLNS it's considered as a namespace
	if strings.HasPrefix(attrName, XMLNS) {
		attrNameTokens := strings.Split(attrName, ":")
		if len(attrNameTokens) > tagWithNamespaceSize {
			l.recordError(
				xmlerrors.NewListenerErrorFromPosition(
					"unrecognized namespace definition format",
					ctx.GetStart(),
				))
			return
		}
		if len(attrNameTokens) == 1 {
			l.currentNode = l.currentNode.WithAttribute(attrName, attrValue)
		}
		l.namespaces[attrNameTokens[1]] = attrValue
		return
	}
	// places attribute to node attribute map
	l.currentNode = l.currentNode.WithAttribute(attrName, attrValue)
}

// getAttNameValuePair function extracts name and value of an attribute and removes double quotes from value.
// If either name or value token retrieved from the context are nil, function will return an error.
func (l *XMLListener) getAttNameValuePair(ctx *genparser.AttributeContext) (attrName, attrValue string, err error) {
	nameToken, valueToken := ctx.GetName(), ctx.GetValue()
	if nameToken == nil || valueToken == nil {
		return "", "", xmlerrors.NewListenerErrorFromPosition(
			"could not get name and value of an attribute from nil tokens",
			ctx.GetStart(),
		)
	}
	return nameToken.GetText(), trimQuotes(valueToken.GetText()), nil
}
