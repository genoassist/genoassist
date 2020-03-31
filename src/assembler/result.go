// contains the definition of an assembly result
package assembler

// result represents the result of the assembly process
type result struct {
	message    string // the message associated with the result
	err        error  // the error of the process
	exitStatus int    // exit status code of the assembler
}

func NewResult(msg string, err error, est int) *result {
	return &result{
		message:    msg,
		err:        err,
		exitStatus: est,
	}
}

// GetMessage returns the message that represents the result of the assembly, empty if successful
func (r *result) GetMessage() string {
	return r.message
}

// GetError returns the error that was created during the process, nil if successful
func (r *result) GetError() error {
	return r.err
}

// GetExitStatusCode returns the status code that was returned upon assembly completion
func (r *result) GetExitStatusCode() int {
	return r.exitStatus
}
