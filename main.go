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

const dummyFASTQ = "dummy_sequence.fastq"

func main() {
	fileParam := "fastq"
	fileValues := dummyFASTQ
	fileUsage := "the path to the FASTQ file containing raw sequence data for assembly"
	rawSequenceFile := flag.String(fileParam, fileValues, fileUsage)

	prepParam := "prep"
	prepUsage := "whether to install all the necessary Docker containers for assembly as a preparatory step"
	prep := flag.Bool(prepParam, true, prepUsage)

	outParam := "out"
	outValue, _ := os.Getwd()
	outUsage := "the path to the directory where results will be stored, defaults to current working directory"
	out := flag.String(outParam, outValue, outUsage)

	// parsing the flags has to be done after setting up all the flags
	flag.Parse()

	if *rawSequenceFile == dummyFASTQ {
		flag.PrintDefaults()
		panic(fmt.Sprintf("the flag -fastq <path/to/sequence.fastq> is required"))
	}

	if *prep {
		errs := prepper.New()
		for len(errs) > 0 {
			select {
			case err := <-errs:
				fmt.Printf("[WARNING] encountered error pulling Docker images, err: %v\n", err)
			default:
				continue
			}
		}
	}
	mst := master.New(*rawSequenceFile, *out)
	if err := mst.Process(); err != nil {
		panic(fmt.Sprintf("failed to run master process, err: %v", err))
	}
}
