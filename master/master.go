// master is responsible for interacting with the user to take in the contigs that need to be assembled. Files
// that are specified by the user are passed to a slave that knows how to schedule Docker containers for all the
// assemblers that are integrated with genomagic
package master

import (
	"fmt"

	"github.com/genomagic/slave"
)

// mst defines the master struct, which is used to coordinate slaves and launch assembly, parsing, and
// reporting slaves
type mst struct {
	// a path to a raw sequencing FASTQ file to perform assembly on
	filePath string
	// a collection of assembly results used by the assembly slave
	assemblyResults chan slave.Result
	// a collection of parsing results used by the parsing slave
	parsingResults chan slave.Result
}

// NewMaster creates and returns a new master struct for the file located at the given file path
func NewMaster(rSeqFile string) Master {
	return &mst{
		filePath:        rSeqFile,
		assemblyResults: make(chan slave.Result),
		parsingResults:  make(chan slave.Result),
	}
}

// Process launches the assembly of the contigs it was created with
func (m *mst) Process() error {
	// first we perform the assembly
	assemblySlave := slave.NewSlave("assembly process initiated by master", m.filePath, slave.Assembly)
	if err := assemblySlave.Process(); err != nil {
		return fmt.Errorf("slave assembly process failed with err: %v", err)
	}
	reportSlave := slave.NewSlave("reporting/parse process initiated by master", m.filePath, slave.Parse)
	if err := reportSlave.Process(); err != nil {
		return fmt.Errorf("slave parsing process failed with err: %v", err)
	}
	return nil
}
