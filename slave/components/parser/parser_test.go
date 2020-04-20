package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewParser(t *testing.T) {
	tests := []struct {
		name        string
		filePath    string
		expectedErr error
	}{
		{
			name:        "test_parser_not_created_with_empty_file_path",
			filePath:    "",
			expectedErr: fmt.Errorf("cannot initialize parser with an empty file path"),
		},
		{
			name:        "test_parser_created_without_error",
			filePath:    "file/path",
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, err := New(tt.filePath, "", "")
			if err != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}

// TODO: this is, essentially, equivalent to no tests, add actual tests when Process is implemented
func TestPrser_Process(t *testing.T) {
	tests := []struct {
		name        string
		filePath    string
		expectedErr error
	}{
		{
			name:        "test_parser_process_returns_nil_on_success",
			filePath:    "file/path",
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p, err := New(tt.filePath, "", "")
			// this test should have no errors from initializing a parser
			if err != nil {
				panic(fmt.Sprintf("TestPrser_Process failed to initialize the parser, err: %v", err))
			}
			// TODO: Comment this out for now. Need to add testing.
			//_, err = p.Process()
			//fmt.Println(err)
			//if err != nil {
			//	assert.EqualError(t, err, tt.expectedErr.Error())
			//}
		})
	}
}
