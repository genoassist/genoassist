// a collection of interface specifications for objects that are part of the master package
package master

// Master is the interface that defines the actions that are accessible by the user of genomagic
type Master interface {
	// Process launches the assembly of the contings it was created with
	Process()
}
