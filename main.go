// main package, and file, is responsible for taking in users arguments, parsing them, and
// calling on the primary to perform the work that genomagic does
package main

import (
	"flag"
	"fmt"

	"github.com/genomagic/config_parser"
	"github.com/genomagic/prepper"
	"github.com/genomagic/primary"
)

const (
	// default flag values
	dummyYAML = "noyaml.yaml"
)

func main() {
	yamlParam := "yaml"
	yamlValue := dummyYAML
	yamlUsage := "the path to the YAML configuration file"
	yaml := flag.String(yamlParam, yamlValue, yamlUsage)

	// parsing the flags has to be done after setting up all the flags
	flag.Parse()

	if *yaml == dummyYAML {
		panic("ERROR: The flag -yaml is required")
	}

	cfg, err := config_parser.ParseConfig(*yaml)
	if err != nil {
		panic(fmt.Sprintf("the YAML config file was incorrect, err: %s", err))
	}

	errs := prepper.NewPrep(cfg)
	for len(errs) > 0 {
		select {
		case err := <-errs:
			if err != nil {
				fmt.Printf("[WARNING] encountered error pulling Docker images, err: %s", err)
			}
		default:
			continue
		}
	}

	prim := primary.New(cfg)
	if err := prim.Process(); err != nil {
		panic(fmt.Sprintf("failed to run primary process, err: %v", err))
	}
}
