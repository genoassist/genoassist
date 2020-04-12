// master is responsible for interacting with the user to take in the contigs that need to be assembled. Files
// that are specified by the user are passed to a slave that knows how to schedule Docker containers for all the
// assemblers that are integrated with genomagic
package master

import "github.com/genomagic/slave"

// mst defines the master struct, which is used to coordinate slaves and launch assembly, parsing, and
// reporting slaves
type mst struct {
	// a slice of file names that contain contigs to be assembled and analyzed
	fileNames []string
	// a collection of assembly results used by the assembly slave
	assemblyResults chan slave.Result
	// a collection of parsing results used by the parsing slave
	parsingResults chan slave.Result
}

// NewMaster creates and returns a new master struct for the given filenames
func NewMaster(fn []string) Master {
	return &mst{
		fileNames:       fn,
		assemblyResults: make(chan slave.Result),
		parsingResults:  make(chan slave.Result),
	}
}

// Process launches the assembly of the contigs it was created with
func (m *mst) Process() {
}
