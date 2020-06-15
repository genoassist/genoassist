package reporter

import (
	"fmt"
	"io"
	"os"
	"path"
	"testing"

	"github.com/biogo/biogo/alphabet"
	"github.com/biogo/biogo/io/seqio/fasta"
	"github.com/biogo/biogo/seq"
	"github.com/biogo/biogo/seq/linear"
	"github.com/stretchr/testify/assert"

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
		name           string
		assemblyName   string
		result         *result.Result
		expectedN50    int32
		expectedN50Err error
		expectedL50    int32
		expectedL50Err error
		expectedErr    error
	}{
		{
			name:           "test_returns_expected_N50_failure",
			assemblyName:   TestAssembly,
			result:         result.New(TestAssembly, badSeqs),
			expectedN50:    0,
			expectedN50Err: fmt.Errorf("the reporter process has not been executed"),
			expectedL50:    0,
			expectedL50Err: fmt.Errorf("the reporter process has not been executed"),
			expectedErr: fmt.Errorf("failed to compute N50 for assembly %s, err: %v",
				TestAssembly, fmt.Errorf("failed to compute N50 due to potentially missing contigs")),
		},
		{
			name:           "test_returns_expected_stats",
			assemblyName:   TestAssembly,
			result:         result.New(TestAssembly, goodSeqs),
			expectedN50:    8,
			expectedN50Err: nil,
			expectedL50:    3,
			expectedL50Err: nil,
			expectedErr:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rep := NewReporter(TestAssembly, tt.result)
			if err := rep.Process(); err != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
			if n50, n50err := rep.GetN50(); n50err != nil {
				assert.EqualError(t, tt.expectedN50Err, n50err.Error())
			} else {
				assert.Equal(t, tt.expectedN50, n50)
			}
			if l50, l50err := rep.GetL50(); l50err != nil {
				assert.EqualError(t, tt.expectedL50Err, l50err.Error())
			} else {
				assert.Equal(t, tt.expectedL50, l50)
			}
		})
	}
}
