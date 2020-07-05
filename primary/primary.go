// primary is responsible for interacting with the user to take in the contigs that need to be assembled. Files
// that are specified by the user are passed to a slave that knows how to schedule Docker containers for all the
// assemblers that are integrated with genomagic
package primary

import (
	"fmt"

	"github.com/genomagic/config_parser"
	"github.com/genomagic/reporter"
	"github.com/genomagic/slave"
)

// primaryProcess defines the primary struct, which is used to coordinate slaves and launch assembly, parsing, and
// reporting slaves
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
		qualityControlSlave := slave.New(m.config, slave.QualityControl)
		if _, err := qualityControlSlave.Process(); err != nil {
			return fmt.Errorf("slave quality control process failed, err: %s", err)
		}
	}

	assemblySlave := slave.New(m.config, slave.Assembly)
	if _, err := assemblySlave.Process(); err != nil {
		return fmt.Errorf("slave assembly process failed with err: %s", err)
	}

	parserSlave := slave.New(m.config, slave.Parse)
	results, err := parserSlave.Process()
	if err != nil {
		return fmt.Errorf("slave parsing process failed with err: %s", err)
	}

	var reports []reporter.Reporter
	for _, r := range results {
		rep := reporter.NewReporter("", r)
		if err := rep.Process(); err != nil {
			return fmt.Errorf("failed to construct report")
		}
		reports = append(reports, rep)
	}
	return nil
}