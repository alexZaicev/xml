package listeners_test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/alexZaicev/xml/go/tree"
	"github.com/ghodss/yaml"
	"github.com/stretchr/testify/require"
)

type TestSuite struct {
	Name    string      `json:"name"`
	Valid   []*TestCase `json:"valid,omitempty"`
	Invalid []*TestCase `json:"invalid,omitempty"`
}

type TestCase struct {
	Name             string            `json:"name"`
	Document         string            `json:"document"`
	ExpectedDocument *tree.XMLDocument `json:"expectedDocument,omitempty"`
	ExpectToFail     bool              `json:"expectToFail,omitempty"`
	ExpectedErrors   []string          `json:"expectedErrors,omitempty"`
}

// readTreeTestData function reads YAML test cases targeting OData tree structure
func readTestSuites(t *testing.T) []*TestSuite {
	// filterTestDataDir is a *string, the `flag` package requires a default value, so
	// this pointer is always non-nil
	files, err := os.ReadDir(*semanticTestDir)
	require.NoError(t, err, "Could not read test case directory.")

	var testCases []*TestSuite
	for _, file := range files {
		require.True(t, isTestFile(file), fmt.Sprintf("Test case file %q is not a valid YAML file", file))
		// get full filepath
		bytes, err := os.ReadFile(filepath.Join(*semanticTestDir, file.Name()))
		require.NoError(t, err)

		// read test cases into buffer and then append to test data slice
		var suite *TestSuite
		require.NoError(t, yaml.Unmarshal(bytes, &suite), "Failed to unmarshal test cases")
		testCases = append(testCases, suite)
	}

	return testCases
}

// isTestFile returns true if the DirEntry ends with a .yaml or .yml suffix
func isTestFile(f os.DirEntry) bool {
	return strings.HasSuffix(f.Name(), ".yaml") || strings.HasSuffix(f.Name(), ".yml")
}
