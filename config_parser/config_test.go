package config_parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TODO: Add more thorough tests
func Test_ParseConfig(t *testing.T) {
	testCases := []struct {
		name           string
		filePath       string
		expectedConfig *Config
		expectedError  error
	}{
		{
			name:           "test_process_returns_expected_config",
			filePath:       "test1.yaml",
			expectedConfig: &Config{},
			expectedError:  nil,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(u *testing.T) {
			retConfig, err := ParseConfig(tt.filePath)
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedConfig, retConfig)
		})
	}
}
