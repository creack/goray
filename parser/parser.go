package parser

import "github.com/creack/goray/rt"

var Parsers = map[string]Parser{}

func RegisterParser(name string, parser Parser) {
	Parsers[name] = parser
}

// Parser is the interface that wraps the parsers.
// TODO: Add Validate() method to allow for simple scene validation.
type Parser interface {
	Parse(string) (*rt.SceneConfig, error)
	Extensions() []string
}
