// a collection of interface specifications for objects that are part of the slave package
// the slave components have to implement the interfaces of this collection in order for the slave and
// its components to do work in a decoupled manner
package result

import "github.com/biogo/biogo/seq"

// Result defines the operations that apply to an assembly result
type Result interface {
	// GetAssemblyName returns the name of the assembly the result was created for
	GetAssemblyName() string
	// GetSequences returns the contigs the result contains
	GetSequences() []seq.Sequence
}