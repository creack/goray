package cli

import (
	"fmt"
	"path"
	"strings"

	"github.com/creack/goray/parser"
)

// ParserCLI wraps a parser with a name.
// Used to select the parser from the CLI.
type ParserCLI struct {
	name   string
	Parser parser.Parser
}

// Set will set and initialize the ParserCLI.
func (pc *ParserCLI) Set(value string) error {
	parse, ok := parser.Parsers[value]
	if !ok {
		possible := make([]string, 0, len(parser.Parsers))
		for k := range parser.Parsers {
			possible = append(possible, k)
		}
		return fmt.Errorf("Unkown parser: %q. Possible values: %v", value, possible)
	}
	pc.name = value
	pc.Parser = parse
	return nil
}

// String returns the parser's name.
func (pc *ParserCLI) String() string {
	return pc.name
}

// DetectParser lookup the given file extension and determine the
// proper parser.
func DetectParser(filename string) string {
	ext := strings.Trim(path.Ext(filename), ".")
	for parserName, parse := range parser.Parsers {
		if exts := parse.Extensions(); inArray(ext, exts) {
			return parserName
		}
	}
	return ""
}

// inArray check if `s` is the given array.
func inArray(s string, array []string) bool {
	for _, v := range array {
		if v == s {
			return true
		}
	}
	return false
}
