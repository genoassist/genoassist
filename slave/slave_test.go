package slave

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlaveProcess(t *testing.T) {
	tests := []struct {
		name        string
		filePath    string
		workType    ComponentWorkType
		shouldFail  bool
		expectedErr error
	}{
		{
			name:        "test_process_returns_err_on_unrecognized_work_type",
			filePath:    "file/path",
			workType:    "workType",
			shouldFail:  false,
			expectedErr: fmt.Errorf("failed to initialize worker"),
		},
		{
			name:        "test_process_fails_when_worker_process_fails",
			filePath:    "file/path",
			workType:    Assembly,
			shouldFail:  true,
			expectedErr: fmt.Errorf("slave process failed"),
		},
		{
			name:        "test_process_returns_nil_on_success",
			filePath:    "file/path",
			workType:    Parse,
			shouldFail:  false,
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := NewMock(tt.name, tt.filePath, tt.workType, tt.shouldFail)
			s.Mock.On("Process", tt.workType, tt.filePath, tt.workType, tt.shouldFail).Return(tt.expectedErr)
			err := s.Process()
			if err != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}
