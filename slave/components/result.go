// contains the definition of an assembly rst
package components

// rst represents the rst of the assembly process
type rst struct {
	// the message associated with the rst
	message string
	// the error of the process
	err error
	// exit status code of the assembler
	exitStatus int
}

// NewResult creates a new result struct and returns it
func NewResult(msg string, err error, est int) *rst {
	return &rst{
		message:    msg,
		err:        err,
		exitStatus: est,
	}
}

// GetMessage returns the message that represents the rst of the assembly, empty if successful
func (r *rst) GetMessage() string {
	return r.message
}

// GetError returns the error that was created during the process, nil if successful
func (r *rst) GetError() error {
	return r.err
}

// GetExitStatusCode returns the status code that was returned upon assembly completion
func (r *rst) GetExitStatusCode() int {
	return r.exitStatus
}
