package parser_test

import (
	"testing"

	"github.com/alexZaicev/xml/go/internal/parser"
	"github.com/stretchr/testify/assert"
)

func Test_XMLParser_Validate(t *testing.T) {
	for _, testSuite := range readTestSuites(t) {
		ts := testSuite
		for _, testCase := range ts.Valid {
			tc := testCase
			t.Run(tc.Name, func(t *testing.T) {
				p := parser.NewXMLParser(tc.Document)
				err := p.Validate()

				if tc.ExpectToFail {
					assert.Error(t, err, "Test case failed validation: %s", tc.Name)
					return
				}
				assert.NoError(t, err, "Test case failed validation: %s", tc.Name)
			})
		}

		for _, testCase := range ts.Invalid {
			tc := testCase
			t.Run(tc.Name, func(t *testing.T) {
				p := parser.NewXMLParser(tc.Document)
				err := p.Validate()

				if tc.ExpectToFail {
					assert.NoError(t, err, "Test case failed validation: %s", tc.Name)
					return
				}
				assert.Error(t, err, "Test case failed validation: %s", tc.Name)
			})
		}
	}
}
