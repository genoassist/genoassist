package parser

import (
	"fmt"

	"github.com/genomagic/result"
	"github.com/genomagic/slave/components"
)

// structure of the parser
type prser struct {
	// path of the file the parser will operate on
	filePath string
}

// New creates and returns a new parser struct
// TODO: parser needs to take in an additional param for the assembler results to process, left as _ to satisfy interface
func New(filePath, _, _ string) (components.Component, error) {
	if filePath == "" {
		return nil, fmt.Errorf("cannot initialize parser with an empty file path")
	}

	return &prser{
		filePath: filePath,
	}, nil
}

// Process performs the work of the parser
func (p *prser) Process() (*result.Result, error) {
	return nil, nil
}
