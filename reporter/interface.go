package reporter

// Reporter defines the interface of the reporter that constructs stats' of the assemblies
type Reporter interface {
	Process() error
}
