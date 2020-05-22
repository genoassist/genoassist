package slave

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/genomagic/config_parser"
)

func TestSlaveProcess(t *testing.T) {
	tests := []struct {
		name        string
		config      *config_parser.Config
		workType    ComponentWorkType
		expectedErr error
	}{
		{
			name: "test_process_returns_err_on_unrecognized_work_type",
			config: &config_parser.Config{
				Assemblers: config_parser.AssemblerConfig{},
				GenoMagic: config_parser.GenoMagicConfig{
					InputFilePath: "in",
					OutputPath:    "out",
					Threads:       0,
					Prep:          false,
				},
			},
			workType:    "workType",
			expectedErr: fmt.Errorf("failed to initialize worker"),
		},
		{
			name: "test_process_fails_when_worker_process_fails",
			config: &config_parser.Config{
				Assemblers: config_parser.AssemblerConfig{},
				GenoMagic: config_parser.GenoMagicConfig{
					InputFilePath: "in",
					OutputPath:    "out",
					Threads:       0,
					Prep:          false,
				},
			},
			workType:    Assembly,
			expectedErr: fmt.Errorf("slave process failed"),
		},
		{
			name: "test_process_returns_nil_on_success",
			config: &config_parser.Config{
				Assemblers: config_parser.AssemblerConfig{},
				GenoMagic: config_parser.GenoMagicConfig{
					InputFilePath: "in",
					OutputPath:    "out",
					Threads:       0,
					Prep:          false,
				},
			},
			workType:    Parse,
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := NewMock(tt.config, tt.workType)
			s.Mock.On("Process").Return(tt.expectedErr)
			err := s.Process()
			if err != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}
