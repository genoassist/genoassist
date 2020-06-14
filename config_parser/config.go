package config_parser

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type (
	Config struct {
		Assemblers AssemblerConfig `yaml:"assemblers"`
		GenoMagic  GenoMagicConfig `yaml:"genomagic"`
	}

	AssemblerConfig struct {
		Megahit MegahitConfig `yaml:"megahit"`
		Abyss   AbyssConfig   `yaml:"abyss"`
		Flye    FlyeConfig    `yaml:"flye"`
	}

	GenoMagicConfig struct {
		Assemblers     []string `yaml:"assemblers,flow"`
		InputFilePath  string   `yaml:"inputFilePath"`
		OutputPath     string   `yaml:"outputPath"`
		Threads        int      `yaml:"threads"`
		Prep           bool     `yaml:"prep"`
		QualityControl bool     `yaml:"qualityControl"`
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

// ParseConfig takes in YAML file and returns a config struct
func ParseConfig(yamlFilePath string) (*Config, error) {
	config := &Config{}

	yamlFile, err := ioutil.ReadFile(yamlFilePath)
	if err != nil {
		return nil, fmt.Errorf("cannot read file %s, err: %v", yamlFilePath, err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshall the contents of the yamlFile into the struct, err: %v", err)
	}
	return config, nil
}
