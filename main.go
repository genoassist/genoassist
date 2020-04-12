// main package, and file, is responsible for taking in users arguments, parsing them, and
// calling on the master to perform the work that genomagic does
package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/genomagic/master"
	"github.com/genomagic/prepper"
)

func main() {
	filesParam := "files"
	filesValues := "file1.txt,file2.txt,..."
	filesUsage := "the file names of the files that store contigs to assemble, parse, and compute statistics for"
	filesSeparator := ","
	files := flag.String(filesParam, filesValues, filesUsage)

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
	fileNames := strings.Split(*files, filesSeparator)
	mst := master.NewMaster(fileNames)
	mst.Process()
}
