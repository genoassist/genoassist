// a collection of interface specifications for objects that are part of the slave package
// the slave components have to implement the interfaces of this collection in order for the slave and
// its components to do work in a decoupled manner
package result

// Result defines the operations that apply to an assembly result
type Result interface {
	// GetMessage returns the message that represents the result of the assembly, empty if successful
	GetMessage() string
	// GetError returns the error that was created during the process, nil if successful
	GetError() error
	// GetExitStatusCode returns the status code that was returned upon assembly completion
	GetExitStatusCode()
}