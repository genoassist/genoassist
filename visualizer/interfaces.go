package visualizer

type (
	// Visualizer represents the interface that is used to interact with the visualizer package. It allows the
	// construction of specific visualization that can be created after performing assembly and reporting steps
	Visualizer interface {
		// N50 creates the plots associated with reporting the N50 scores for assembly reports
		N50() error
		// L50 creates the plots associated with reporting the L50 scores for assembly reports
		L50() error
	}
)
