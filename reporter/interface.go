package reporter

// Reporter defines the interface of the reporter that constructs stats' of the assemblies
type Reporter interface {
	// Process constructs the report for the given assembler results
	Process() error
	// getL50 computes and returns the L50 score of the report contigs
	getL50() (float32, error)
	// getN50 computes and returns the N50 score of the report contigs
	getN50() (float32, error)
}
