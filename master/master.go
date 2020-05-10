// master is responsible for interacting with the user to take in the contigs that need to be assembled. Files
// that are specified by the user are passed to a slave that knows how to schedule Docker containers for all the
// assemblers that are integrated with genomagic
package master

import (
	"fmt"

	"github.com/genomagic/reporter"
	"github.com/genomagic/result"
	"github.com/genomagic/slave"
)

// mst defines the master struct, which is used to coordinate slaves and launch assembly, parsing, and
// reporting slaves
type mst struct {
	filePath        string             // a path to a raw sequencing FASTQ file to perform assembly on
	outPath         string             // a path to the location where results will be stored
	numThreads      int                // number of threads to use for slave processes
	assemblyResults chan result.Result // a collection of assembly results used by the assembly slave
	parsingResults  chan result.Result // a collection of parsing results used by the parsing slave
}

// New creates and returns a new master struct for the file located at the given file path
func New(rsf, out string, numThreads int) Master {
	return &mst{
		filePath:        rsf,
		outPath:         out,
		numThreads:      numThreads,
		assemblyResults: make(chan result.Result),
		parsingResults:  make(chan result.Result),
	}
}

// Process launches the assembly of the contigs it was created with
func (m *mst) Process() error {
	assemblySlave := slave.New("assembly process initiated by master", m.filePath, m.outPath, m.numThreads, slave.Assembly)
	if _, err := assemblySlave.Process(); err != nil {
		return fmt.Errorf("slave assembly process failed with err: %v", err)
	}

	parserSlave := slave.New("reporting/parse process initiated by master", m.filePath, m.outPath, m.numThreads, slave.Parse)
	// TODO: Take the result obtained from Parse process and feed it into the reporter
	results, err := parserSlave.Process()
	if err != nil {
		return fmt.Errorf("slave parsing process failed with err: %v", err)
	}

	var reports []reporter.Reporter
	for _, r := range results {
		rep := reporter.New("", r)
		if err := rep.Process(); err != nil {
			return fmt.Errorf("failed to construct report")
		}
		reports = append(reports, rep)
	}
	return nil
}
