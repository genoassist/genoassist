package parser

import (
	"github.com/genomagic/src/slave/components"
)

// structure of the parser
type prser struct {
	// name of the file the parser will operate on
	fileName string
}

// NewParser creates and returns a new parser struct
func NewParser(fnm string) (components.Component, error) {
	return &prser{
		fileName: fnm,
	}, nil
}

// Process performs the work of the parser
func (p *prser) Process() error {
	return nil
}
