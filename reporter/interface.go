package reporter

// Reporter defines the interface of the reporter that constructs stats' of the assemblies
type Reporter interface {
	// Process constructs the report for the given assembler results
	Process() error
	// GetL50 returns the computed L50 value stored on the report
	GetL50() int32
	// GetN50 returns the N50 score of the report contigs
	GetN50() int32
}
