package parser

import (
	"github.com/genomagic/slave/components"
)

// structure of the parser
type prser struct {
	// path of the file the parser will operate on
	filePath string
}

// NewParser creates and returns a new parser struct
func NewParser(fnm string) (components.Component, error) {
	return &prser{
		filePath: fnm,
	}, nil
}

// Process performs the work of the parser
func (p *prser) Process() error {
	return nil
}
