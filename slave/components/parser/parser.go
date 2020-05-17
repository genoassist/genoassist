package parser

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/biogo/biogo/alphabet"
	"github.com/biogo/biogo/io/seqio/fasta"
	"github.com/biogo/biogo/seq"
	"github.com/biogo/biogo/seq/linear"

	"github.com/genomagic/constants"
	"github.com/genomagic/result"
	"github.com/genomagic/slave/components"
)

// structure of the parser
type prser struct {
	filePath      string // path of the file the parser will operate on
	outPath       string // output directory where to store results
	assemblerName string // name of the assembler to parse the results of
}

// New creates and returns a new parser struct
func New(fp, op, ap string) (components.Component, error) {
	if fp == "" {
		return nil, fmt.Errorf("cannot initialize parser with an empty file path")
	}
	return &prser{
		filePath:      fp,
		outPath:       op,
		assemblerName: ap,
	}, nil
}

// Process performs the work of the parser
func (p *prser) Process() (*result.Result, error) {
	outDir := constants.AvailableAssemblers[p.assemblerName].OutputDir
	contigsFile := constants.AvailableAssemblers[p.assemblerName].AssemblyFileName
	ioRdr, err := os.Open(path.Join(p.outPath, outDir, contigsFile))
	if err != nil {
		return nil, fmt.Errorf("cannot open file: %v", err)
	}

	fastaRdr := fasta.NewReader(ioRdr, linear.NewSeq("", nil, alphabet.DNA))
	var sequences []seq.Sequence
	for {
		inSeq, err := fastaRdr.Read()
		if err != nil {
			if err != io.EOF {
				return nil, fmt.Errorf("cannot read sequence: %v", err)
			}
			break
		}
		sequences = append(sequences, inSeq)
	}
	res := result.New(p.assemblerName, sequences)
	return res, nil
}
