package parser_test

import "flag"

var _ = flag.String(
	"semantic-test-dir",         // name
	"./test-resources/semantic", // default path
	"Filepath to directory containing YAML-format test cases.", // usage
)

var syntaxDataDir = flag.String(
	"syntax-test-dir",         // name
	"./test-resources/syntax", // default path
	"Filepath to directory containing YAML-format test cases.", // usage
)
