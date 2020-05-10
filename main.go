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

const (
	// default flag values
	dummyFASTQ      = "dummy_sequence.fastq"
	defaultThreads  = 2
	tempThreadLimit = 10
)

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

	// TODO: add a check to make sure numThreads does not exceed the limit of host computer
	threadsParam := "threads"
	threadsUsage := "the number of threads that is passed to the assembler programs"
	numThreads := flag.Int(threadsParam, defaultThreads, threadsUsage)
	if *numThreads > tempThreadLimit {
		panic(fmt.Sprintf("Cannot run with a thread number greater than %d", tempThreadLimit))
	}

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

	mst := master.New(*rawSequenceFile, *out, *numThreads)
	if err := mst.Process(); err != nil {
		panic(fmt.Sprintf("failed to run master process, err: %v", err))
	}
}
