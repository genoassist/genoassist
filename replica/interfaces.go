package replica

import "github.com/genomagic/result"

// Replica defines the operations that are accessible on a replica
type Replica interface {
	// Process performs the work that's dictated by the primary
	Process() ([]*result.Result, error)
}
