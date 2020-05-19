// slave is responsible for launching and coordinating processes such as
// assembly and parsing for the master
package slave

import (
	"fmt"

	"github.com/genomagic/config_parser"
	"github.com/genomagic/constants"
	"github.com/genomagic/result"
	"github.com/genomagic/slave/components/assembler"
	"github.com/genomagic/slave/components/parser"
	"github.com/genomagic/slave/components/quality_controller"
)

const (
	// Assembly tells the slave that it needs to launch an assembly process
	Assembly = "assembly"
	// Parse tells the slave that it needs to launch a parse process
	Parse = "parse"
	// QualityControl tells the slave that it needs to launch a quality control process
	QualityControl = "quality_control"
)

// the type of component that runs on a specific set of assembly files
type ComponentWorkType string

// slaveProcess defines the structure of a slave
type slaveProcess struct {
	// config is the GenoMagic configuration that is passed through YAML config file
	config *config_parser.Config
	// workType is the type of work that has to be performed by the slave
	workType ComponentWorkType
}

// New creates and returns a new instance of a slave
func New(config *config_parser.Config, workType ComponentWorkType) Slave {
	return &slaveProcess{
		config:   config,
		workType: workType,
	}
}

// Process performs the work that's dictated by the master
func (s *slaveProcess) Process() ([]*result.Result, error) {
	switch s.workType {
	case Assembly:
		for k := range constants.AvailableAssemblers {
			assemblerWorker, err := assembler.New(s.config.GenoMagic.InputFilePath, s.config.GenoMagic.OutputPath, k, s.config)
			if err != nil {
				return nil, fmt.Errorf("failed to initialize assembler worker, err %v", err)
			}

			_, err = assemblerWorker.Process()
			if err != nil {
				return nil, fmt.Errorf("assembler slave process failed, err: %v", err)
			}
		}
		return nil, nil
	case Parse:
		var results []*result.Result
		for k := range constants.AvailableAssemblers {
			parserWorker, err := parser.New(s.config.GenoMagic.InputFilePath, s.config.GenoMagic.OutputPath, k)
			if err != nil {
				return nil, fmt.Errorf("failed to initialize parser worker, err: %v", err)
			}
			res, err := parserWorker.Process()
			if err != nil {
				return nil, fmt.Errorf("parser slave process failed, err: %v", err)
			}
			results = append(results, res)
		}
		return results, nil
	case QualityControl:
		qualityController := quality_controller.NewQualityController(s.config)
		if newSeqFilePath, err := qualityController.Process(); err != nil {
			return nil, fmt.Errorf("quality control slave failed, err: %s", err)
		} else {
			s.config.GenoMagic.InputFilePath = newSeqFilePath
		}
		return nil, nil
	default:
		return nil, fmt.Errorf("slave presented with unknown operation")
	}
}
