package reporter

// Reporter defines the interface of the reporter that constructs stats' of the assemblies
type Reporter interface {
	// Process constructs the Report for the given assembler results
	Process() error
	// GetL50 returns the computed L50 value stored on the Report. An error is returned if the reporter process has not been executed
	GetL50() (int32, error)
	// GetN50 returns the N50 score of the Report contigs. An error is returned if the reporter process has not been executed
	GetN50() (int32, error)
}
