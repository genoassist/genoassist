// primary is responsible for interacting with the user to take in the contigs that need to be assembled. Files
// that are specified by the user are passed to a replica that knows how to schedule Docker containers for all the
// assemblers that are integrated with genomagic
package primary

import (
	"fmt"
	"path"

	"github.com/genomagic/visualizer"

	"github.com/genomagic/config_parser"
	"github.com/genomagic/replica"
	"github.com/genomagic/reporter"
)

// primaryProcess defines the primary struct, which is used to coordinate replicas and launch assembly, parsing, and
// reporting replicas
type primaryProcess struct {
	// config is the configuration of GenoMagic obtained through YAML config file
	config *config_parser.Config
}

// NewReporter creates and returns a new primary struct for the file located at the given file path
func New(config *config_parser.Config) Primary {
	return &primaryProcess{
		config: config,
	}
}

// Process launches the assembly of the contigs it was created with
func (m *primaryProcess) Process() error {
	if m.config.GenoMagic.QualityControl {
		qualityControlReplica := replica.New(m.config, replica.QualityControl)
		if _, err := qualityControlReplica.Process(); err != nil {
			return fmt.Errorf("replica quality control process failed, err: %s", err)
		}
	}

	assemblyReplica := replica.New(m.config, replica.Assembly)
	if _, err := assemblyReplica.Process(); err != nil {
		return fmt.Errorf("replica assembly process failed, err: %s", err)
	}

	parseReplica := replica.New(m.config, replica.Parse)
	results, err := parseReplica.Process()
	if err != nil {
		return fmt.Errorf("replica parsing process failed, err: %s", err)
	}

	var reports []reporter.Reporter
	for _, r := range results {
		rep := reporter.NewReporter(r.AssemblyName, r)
		if err := rep.Process(); err != nil {
			return fmt.Errorf("replica failed to construct report, err: %s", err)
		}
		reports = append(reports, rep)
	}

	vizOutPath := path.Join(m.config.GenoMagic.OutputPath, "report.html")
	viz := visualizer.NewVisualizer(reports, vizOutPath)
	if err := viz.Process(); err != nil {
		return fmt.Errorf("replica failed to visualize the reports, err: %s", err)
	}
	return nil
}
