package parser

import (
	"fmt"
	"io"
	"os"

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
	// path of the file the parser will operate on
	filePath         string
	outPath          string
	assemblerProcess string
}

// New creates and returns a new parser struct
func New(filePath, outPath, asmblrProcess string) (components.Component, error) {
	if filePath == "" {
		return nil, fmt.Errorf("cannot initialize parser with an empty file path")
	}

	return &prser{
		filePath:         filePath,
		outPath:          outPath,
		assemblerProcess: asmblrProcess,
	}, nil
}

// Process performs the work of the parser
func (p *prser) Process() (*result.Result, error) {

	ioReader, err := os.Open(p.outPath +  constants.AvailableAssemblers[p.assemblerProcess].OutputDir +  "/final.contigs.fa");
	if err != nil {
		return nil, fmt.Errorf("cannot open file: %v", err)
	}

	reader := fasta.NewReader(ioReader, linear.NewSeq("",nil,alphabet.DNA))
	var sequences []seq.Sequence
	for  {
		seq, err := reader.Read()
		if err != nil {
			if err != io.EOF {
				return nil, fmt.Errorf("cannot read sequence: %v", err)
			}
			break
		}
		sequences = append(sequences,seq)
	}

	res := result.New(sequences)
	return &res, nil
}
