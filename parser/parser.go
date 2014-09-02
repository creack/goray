package parser

import "github.com/creack/goray/rt"

type ParseFct func(string) (*rt.SceneConfig, error)

var Parsers = map[string]ParseFct{}

func RegisterParser(name string, parseFct ParseFct) {
	Parsers[name] = parseFct
}
