package cli

import (
	"fmt"
	"path"
	"strings"

	"github.com/creack/goray/parser"
)

type ParserCLI struct {
	name   string
	Parser parser.Parser
}

func (pc *ParserCLI) Set(value string) error {
	parse, ok := parser.Parsers[value]
	if !ok {
		possible := make([]string, len(parser.Parsers))
		i := 0
		for k := range parser.Parsers {
			possible[i] = k
			i++
		}
		return fmt.Errorf("Unkown parser: %s. Possible values: %v", value, possible)
	}
	pc.name = value
	pc.Parser = parse
	return nil
}

func (pc *ParserCLI) String() string {
	return pc.name
}

func DetectParser(filename string) string {
	ext := strings.Trim(path.Ext(filename), ".")
	for parserName, parse := range parser.Parsers {
		if exts := parse.Extensions(); inArray(ext, exts) {
			return parserName
		}
	}
	return ""
}
func inArray(s string, array []string) bool {
	for _, v := range array {
		if v == s {
			return true
		}
	}
	return false
}
