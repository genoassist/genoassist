package visualizer

import "github.com/genomagic/reporter"

type (
	// visualize implements the Visualizer interface for visualizing assembly reports
	visualize struct {
		reports []reporter.Report // the collection of reports to visualize the scores of
	}
)

// NewVisualizer creates and returns a new instance of a struct that implements the Visualizer interface
func NewVisualizer(reports []reporter.Report) Visualizer {
	return &visualize{
		reports: reports,
	}
}

// N50 creates the plots associated with reporting the N50 scores for assembly reports
func (v *visualize) N50() error {
	return nil
}

// L50 creates the plots associated with reporting the N50 scores for assembly reports
func (v *visualize) L50() error {
	return nil
}
