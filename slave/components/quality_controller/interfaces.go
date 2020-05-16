package quality_controller

// Controller represents the interface of the quality control process
type Controller interface {
	// Process performs the work defined by a specific controller implementation. Since processes occurs on
	// sequence files, the paths to the newly generated quality controlled files is returned, along with
	// and error, in case there is one
	Process() (string, error)
}
