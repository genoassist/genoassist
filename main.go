// main package, and file, is responsible for taking in users arguments, parsing them, and
// calling on the master to perform the work that genomagic does
package main

import (
	"flag"
	"fmt"
	"github.com/genomagic/master"
	"github.com/genomagic/prepper"
)

func main() {
	filesParam := "fastq"
	filesValues := "raw_sequence.fastq"
	filesUsage := "the FASTQ file containing raw sequence data for assembly"
	raw_sequence_file := flag.String(filesParam, filesValues, filesUsage)

	prepParam := "prepper"
	prepUsage := "whether to install all the necessary Docker containers for assembly as a preparatory step"
	prep := flag.Bool(prepParam, true, prepUsage)

	// parsing the flags has to be done after setting up all the flags
	flag.Parse()

	if *prep {
		if err := prepper.NewPrep(); err != nil {
			panic(fmt.Sprintf("failed to prep GenoMagic, err: %v", err))
		}
	}
	mst := master.NewMaster(*raw_sequence_file)
	mst.Process()
}
