package visualizer

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-echarts/go-echarts/charts"

	"github.com/genoassist/reporter"
)

type (
	// visualize implements the Visualizer interface for visualizing assembly reports
	visualize struct {
		reports []reporter.Reporter // the collection of reports to visualize the scores of
		name    string              // the testName of the file to save the report to
	}
)

// NewVisualizer creates and returns a new instance of a struct that implements the Visualizer interface
func NewVisualizer(reports []reporter.Reporter, name string) Visualizer {
	return &visualize{
		reports: reports,
		name:    name,
	}
}

// Process creates the plots associated with reporting the N50 and L50 scores for assembly reports
func (v *visualize) Process() error {
	if v.reports == nil || len(v.reports) == 0 {
		return fmt.Errorf("visualizer process cannot create a visualizer with no reports")
	}
	if !strings.Contains(v.name, ".html") {
		return fmt.Errorf("visualizer process cannot take in a name without an .html extension")
	}
	var err error
	names := make([]string, len(v.reports))
	l50s := make([]int32, len(v.reports))
	n50s := make([]int32, len(v.reports))
	for i, n := range v.reports {
		names[i] = n.GetAssemblyName()
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
		AddYAxis("N50", n50s).
		AddYAxis("L50", l50s)
	if f, err := os.Create(v.name); err != nil {
		return err
	} else {
		return bar.Render(f)
	}
}
