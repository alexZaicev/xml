package listeners_test

import (
	"fmt"
	"testing"

	"github.com/alexZaicev/xml/go/internal/listeners"
	"github.com/alexZaicev/xml/go/internal/parser"
	"github.com/alexZaicev/xml/go/tree"
	"github.com/stretchr/testify/assert"
)

func Test_XMLListener(t *testing.T) {
	for _, testSuite := range readTestSuites(t) {
		ts := testSuite
		for _, testCase := range ts.Valid {
			tc := testCase
			t.Run(tc.Name, func(t *testing.T) {
				document, errSlice := parseDocument(tc)
				if tc.ExpectToFail {
					validateExpectedErrors(t, errSlice, tc.ExpectedErrors)
					return
				}
				assert.Empty(t, errSlice)
				// Assert: Expected tree generated
				populateParentsRecursively(tc.ExpectedDocument.Element)
				assert.Equal(t, tc.ExpectedDocument, document)
			})
		}

		for _, testCase := range ts.Invalid {
			tc := testCase
			t.Run(tc.Name, func(t *testing.T) {
				_, errSlice := parseDocument(tc)
				if tc.ExpectToFail {
					assert.Empty(t, errSlice)
					return
				}
				validateExpectedErrors(t, errSlice, tc.ExpectedErrors)
			})
		}
	}
}

func validateExpectedErrors(t *testing.T, actualErrors []error, expectedErrors []string) {
	if assert.Len(t, actualErrors, len(expectedErrors)) {
		for i, expectedErr := range expectedErrors {
			assert.EqualError(
				t,
				actualErrors[i],
				expectedErr,
				fmt.Sprintf("expected error message %q not found", expectedErr),
			)
		}
	}
}

func parseDocument(tc *TestCase) (*tree.XMLDocument, []error) {
	p := parser.NewXMLParser(tc.Document)
	listener := listeners.NewXMLListener()
	grammaticalErrors := p.Parse(listener)

	if grammaticalErrors != nil {
		return nil, []error{grammaticalErrors}
	}
	return listener.Document(), listener.Errors()
}

// populateParentsRecursively walks through the tree breath first populating the Parent property of
// all nodes with the preceding node - this makes defining a tree inline much more concise
func populateParentsRecursively(node *tree.XMLNode) {
	for _, child := range node.Children {
		child.Parent = node
		populateParentsRecursively(child)
	}
}
