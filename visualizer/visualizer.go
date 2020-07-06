package visualizer

import (
	"os"

	"github.com/go-echarts/go-echarts/charts"

	"github.com/genomagic/reporter"
)

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

// Process creates the plots associated with reporting the N50 and L50 scores for assembly reports
func (v *visualize) Process() error {
	var err error
	names := make([]string, len(v.reports))
	l50s := make([]int32, len(v.reports))
	n50s := make([]int32, len(v.reports))
	for i, n := range v.reports {
		names[i] = n.AssemblyName
		if l50s[i], err = n.GetL50(); err != nil {
			return err
		}
		if n50s[i], err = n.GetN50(); err != nil {
			return err
		}
	}
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{
		Title: "Assembly results",
	})
	bar.AddXAxis(names).
		AddYAxis("N50", l50s).
		AddYAxis("L50", n50s)
	if f, err := os.Create("bar.html"); err != nil {
		return err
	} else {
		return bar.Render(f)
	}
}
