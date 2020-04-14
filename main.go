// main package, and file, is responsible for taking in users arguments, parsing them, and
// calling on the master to perform the work that genomagic does
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/genomagic/master"
	"github.com/genomagic/prepper"
)

func main() {
	fileParam := "fastq"
	fileValues := "raw_sequence.fastq"
	fileUsage := "*REQUIRED* the path to the FASTQ file containing raw sequence data for assembly"
	rawSequenceFile := flag.String(fileParam, fileValues, fileUsage)

	prepParam := "prepper"
	prepUsage := "whether to install all the necessary Docker containers for assembly as a preparatory step"
	prep := flag.Bool(prepParam, true, prepUsage)

	// parsing the flags has to be done after setting up all the flags
	flag.Parse()

	// If no flags are given, print the default usage message
	if numFlags := flag.NFlag(); numFlags == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *prep {
		if err := prepper.NewPrep(); err != nil {
			panic(fmt.Sprintf("failed to prep GenoMagic, err: %v", err))
		}
	}
	mst := master.NewMaster(*rawSequenceFile)
	if err := mst.Process(); err != nil {
		panic(fmt.Sprintf("failed to run master process, err: %v", err))
	}
}
