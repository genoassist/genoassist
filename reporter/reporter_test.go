package reporter

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"path"
	"testing"

	"github.com/biogo/biogo/alphabet"
	"github.com/biogo/biogo/io/seqio/fasta"
	"github.com/biogo/biogo/seq"
	"github.com/biogo/biogo/seq/linear"

	"github.com/genomagic/result"
)

const (
	TestAssembly = "test_assembly"
	GoodSeq      = "good_seq.fasta"
	BadSeq       = "bad_seq.fasta"
)

func readSeqs(f string) ([]seq.Sequence, error) {
	wd, _ := os.Getwd()
	ioRdr, _ := os.Open(path.Join(wd, f))
	fRdr := fasta.NewReader(ioRdr, linear.NewSeq("", nil, alphabet.DNA))
	var sequences []seq.Sequence
	for {
		inSeq, err := fRdr.Read()
		if err != nil {
			if err != io.EOF {
				return nil, fmt.Errorf("cannot read sequence: %v", err)
			}
			break
		}
		sequences = append(sequences, inSeq)
	}
	return sequences, nil
}

func TestReport_Process(t *testing.T) {
	t.Parallel()
	var goodSeqs []seq.Sequence
	if gs, err := readSeqs(GoodSeq); err != nil {
		panic("TestReport_Process failed to read good sequences")
	} else {
		goodSeqs = gs
	}

	var badSeqs []seq.Sequence
	if bs, err := readSeqs(BadSeq); err != nil {
		panic("TestReport_Process failed to read bad sequences")
	} else {
		badSeqs = bs
	}

	tests := []struct {
		name         string
		assemblyName string
		result       result.Result
		expectedN50  int32
		expectedL50  int32
		expectedErr  error
	}{
		{
			name:         "test_returns_expected_N50_failure",
			assemblyName: TestAssembly,
			result:       result.New(TestAssembly, badSeqs),
			expectedN50:  0,
			expectedL50:  0,
			expectedErr: fmt.Errorf("failed to compute N50 for assembly %s, err: %v",
				TestAssembly, fmt.Errorf("failed to compute N50 due to potentially missing contigs")),
		},
		{
			name:         "test_returns_expected_stats",
			assemblyName: TestAssembly,
			result:       result.New(TestAssembly, goodSeqs),
			expectedN50:  8,
			expectedL50:  0,
			expectedErr:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rep := New(TestAssembly, tt.result)
			err := rep.Process()
			if err != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
			assert.Equal(t, tt.expectedN50, rep.GetN50())
			assert.Equal(t, tt.expectedL50, rep.GetL50())
		})
	}
}
