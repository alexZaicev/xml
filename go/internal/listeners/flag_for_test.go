package listeners_test

import "flag"

var semanticTestDir = flag.String(
	"semantic-test-dir",         // name
	"./test-resources/semantic", // default path
	"Filepath to directory containing YAML-format test cases.", // usage
)

var _ = flag.String(
	"syntax-test-dir",         // name
	"./test-resources/syntax", // default path
	"Filepath to directory containing YAML-format test cases.", // usage
)
