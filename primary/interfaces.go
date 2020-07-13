// a collection of interface specifications for objects that are part of the primary package
package primary

// Primary is the interface that defines the actions that are accessible by the user of genoassist
type Primary interface {
	// Process launches the assembly of the contigs the primary received for assembly
	Process() error
}
