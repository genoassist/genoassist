// main package, and file, is responsible for taking in users arguments, parsing them, and
// calling on the master to perform the work that genomagic does
package main

import (
	"flag"
	"strings"
	"github.com/genomagic/src/master"
)

func main() {
	filesParam := "files"
	filesValues := "file1.txt,file2.txt,..."
	filesUsage := "the file names of the files that store contigs to assemble, parse, and compute statistics for"
	filesSeparator := ","
	files := flag.String(filesParam, filesValues, filesUsage)

	// parsing the flags has to be done after setting up all the flags
	flag.Parse()

	fileNames := strings.Split(*files, filesSeparator)
	mst := master.NewMaster(fileNames)
	mst.Process()
}
