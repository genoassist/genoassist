package config_parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParseConfig(t *testing.T) {
	testCases := []struct {
		name           string
		filePath       string
		expectedConfig *Config
		expectedError  error
	}{
		{
			name:           "test_process_returns_nil_on_empty_yaml_file",
			filePath:       "test1.yaml",
			expectedConfig: &Config{},
			expectedError:  nil,
		},
		{
			name:     "test_process_returns_expected_config_on_standard_yaml_file",
			filePath: "test2.yaml",
			expectedConfig: &Config{
				Assemblers: AssemblerConfig{
					Megahit: MegahitConfig{KMers: "27"},
					Abyss:   AbyssConfig{KMers: "27"},
					Flye: FlyeConfig{
						SeqType:    "nano",
						GenomeSize: "5m",
					},
				},
				GenoMagic: GenoMagicConfig{
					InputFilePath: "/test/input1.fastq",
					OutputPath:    "/test/output",
					Threads:       2,
				},
			},
			expectedError: nil,
		},
		{
			name:     "test_process_returns_expected_config_on_missing_flye_in_yaml_file",
			filePath: "test3.yaml",
			expectedConfig: &Config{
				Assemblers: AssemblerConfig{
					Megahit: MegahitConfig{KMers: "27"},
					Abyss:   AbyssConfig{KMers: "27"},
					Flye:    FlyeConfig{},
				},
				GenoMagic: GenoMagicConfig{
					InputFilePath: "/test/input1.fastq",
					OutputPath:    "/test/output",
					Threads:       2,
				},
			},
			expectedError: nil,
		},
		{
			name:     "test_process_returns_expected_config_on_additional_data_in_yaml_file",
			filePath: "test4.yaml",
			expectedConfig: &Config{
				Assemblers: AssemblerConfig{
					Megahit: MegahitConfig{KMers: "27"},
					Abyss:   AbyssConfig{KMers: "27"},
					Flye:    FlyeConfig{},
				},
				GenoMagic: GenoMagicConfig{
					InputFilePath: "/test/input1.fastq",
					OutputPath:    "/test/output",
					Threads:       2,
				},
			},
			expectedError: nil,
		},
		{
			name:           "test_process_returns_error_when_invalid_file_path_to_YAML_is_provided",
			filePath:       "",
			expectedConfig: nil,
			expectedError:  fmt.Errorf("cannot read file %s, err: open : no such file or directory", ""),
		},
		{
			name:           "test_process_returns_error_when_wrong_YAML_config_filepath_is_provided",
			filePath:       "wrong_file.fasta",
			expectedConfig: nil,
			expectedError:  fmt.Errorf("cannot unmarshall the contents of the yamlFile into the struct, err: yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `` into config_parser.Config"),
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
