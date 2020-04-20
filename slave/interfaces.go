package slave

import "github.com/genomagic/result"

// Slave defines the operations that are accessible on a slave
type Slave interface {
	// Process performs the work that's dictated by the master
	Process() (*result.Result, error)
}
