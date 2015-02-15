package parser

import "github.com/creack/goray/scene"

// Parsers hold all the available parsers.
// This gets populated from the `init()` of the various
// implementations.
var Parsers = map[string]Parser{}

// RegisterParser registers the given Parser with the given name.
func RegisterParser(name string, parser Parser) {
	Parsers[name] = parser
}

// Parser is the interface that wraps the Parsers.
// TODO: Add Validate() method to allow for simple scene validation.
type Parser interface {
	// Parse parses the given filename into a RT configuration.
	Parse(fileName string) (*scene.Config, error)
	// Extensions declares the supported file types of the parser.
	Extensions() []string
}
