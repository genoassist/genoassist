package config_parser

import (
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

type (
	Config struct {
		Assemblers AssemblerConfig  `yaml:"assemblers"`
		GenoAssist GenoAssistConfig `yaml:"genoassist"`
	}

	AssemblerConfig struct {
		Megahit MegahitConfig `yaml:"megahit"`
		Abyss   AbyssConfig   `yaml:"abyss"`
		Flye    FlyeConfig    `yaml:"flye"`
	}

	GenoAssistConfig struct {
		Assemblers     []string `yaml:"assemblers,flow"`
		InputFilePath  string   `yaml:"inputFilePath"`
		OutputPath     string   `yaml:"outputPath"`
		Threads        int      `yaml:"threads"`
		Prep           bool     `yaml:"prep"`
		QualityControl bool     `yaml:"qualityControl"`
		FileType       string   `yaml:"fileType"`
	}

	MegahitConfig struct {
		KMers string `yaml:"kmers"`
	}

	AbyssConfig struct {
		KMers string `yaml:"kmers"`
	}

	FlyeConfig struct {
		SeqType    string `yaml:"seqType"`
		GenomeSize string `yaml:"genomeSize"`
	}
)

const (
	// FASTQ is the FASTQ input type
	FASTQ = "fastq"
	// FASTA is the FASTA input type
	FASTA = "fasta"
)

// ParseConfig takes in YAML file and returns a config struct
func ParseConfig(yamlFilePath string) (*Config, error) {
	config := &Config{}

	yamlFile, err := ioutil.ReadFile(yamlFilePath)
	if err != nil {
		return nil, fmt.Errorf("cannot read file %s, err: %v", yamlFilePath, err)
	}

	if err = yaml.Unmarshal(yamlFile, &config); err != nil {
		return nil, fmt.Errorf("cannot unmarshall the contents of the yamlFile into the struct, err: %v", err)
	}

	if err = validateConfig(config); err != nil {
		return nil, fmt.Errorf("failed config validation, err: %s", err)
	}

	return config, nil
}

// validateConfig validates the given Config struct for invalid inputs
func validateConfig(c *Config) error {
	// TODO: only checking file extension for now, more validation should be added
	if strings.ToLower(c.GenoAssist.FileType) != FASTQ && strings.ToLower(c.GenoAssist.FileType) != FASTA {
		return fmt.Errorf("invalid file type, allowed [%s, %s]", FASTQ, FASTA)
	}
	return nil
}
