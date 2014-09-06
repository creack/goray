package parser

import "github.com/creack/goray/rt"

type ParseFct func(string) (*rt.SceneConfig, error)

var Parsers = map[string]Parser{}

func RegisterParser(name string, parser Parser) {
	Parsers[name] = parser
}

type Parser interface {
	Parse(string) (*rt.SceneConfig, error)
	Extensions() []string
}
