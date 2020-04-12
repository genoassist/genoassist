// defines interfaces that have to be satisfied by components
package components

// Component defines the operations that apply to slave components such as assemblers and parsers
type Component interface {
	// Process performs the work associated with this worker
	Process() error
}
