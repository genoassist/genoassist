// replica is responsible for launching and coordinating processes such as
// assembly and parsing for the primary
package replica

import (
	"fmt"

	"github.com/genoassist/config_parser"
	"github.com/genoassist/constants"
	"github.com/genoassist/replica/components/assembler"
	"github.com/genoassist/replica/components/parser"
	"github.com/genoassist/replica/components/quality_controller"
	"github.com/genoassist/result"
)

const (
	// Assembly tells the replica that it needs to launch an assembly process
	Assembly = "assembly"
	// Parse tells the replica that it needs to launch a parse process
	Parse = "parse"
	// QualityControl tells the replica that it needs to launch a quality control process
	QualityControl = "quality_control"
)

// the type of component that runs on a specific set of assembly files
type ComponentWorkType string

// replicaProcess defines the structure of a replica
type replicaProcess struct {
	// config is the GenoAssist configuration that is passed through YAML config file
	config *config_parser.Config
	// workType is the type of work that has to be performed by the replica
	workType ComponentWorkType
}

// New creates and returns a new instance of a replica
func New(config *config_parser.Config, workType ComponentWorkType) Replica {
	return &replicaProcess{
		config:   config,
		workType: workType,
	}
}

// Process performs the work that's dictated by the primary
func (s *replicaProcess) Process() ([]*result.Result, error) {
	switch s.workType {
	case Assembly:
		for k := range constants.AvailableAssemblers {
			if contains(s.config.GenoAssist.Assemblers, k) {
				assemblerWorker, err := assembler.New(k, s.config)
				if err != nil {
					return nil, fmt.Errorf("failed to initialize assembler worker, err %v", err)
				}

				_, err = assemblerWorker.Process()
				if err != nil {
					return nil, fmt.Errorf("assembler replica process failed, err: %v", err)
				}
			}
		}
		return nil, nil
	case Parse:
		var results []*result.Result
		for k := range constants.AvailableAssemblers {
			if contains(s.config.GenoAssist.Assemblers, k) {
				parserWorker, err := parser.New(s.config.GenoAssist.InputFilePath, s.config.GenoAssist.OutputPath, k)
				if err != nil {
					return nil, fmt.Errorf("failed to initialize parser worker, err: %v", err)
				}
				res, err := parserWorker.Process()
				if err != nil {
					return nil, fmt.Errorf("parser replica process failed, err: %v", err)
				}
				results = append(results, res)
			}
		}
		return results, nil
	case QualityControl:
		qualityController := quality_controller.NewQualityController(s.config)
		if newSeqFilePath, err := qualityController.Process(); err != nil {
			return nil, fmt.Errorf("quality control replica failed, err: %s", err)
		} else {
			s.config.GenoAssist.InputFilePath = newSeqFilePath
		}
		return nil, nil
	default:
		return nil, fmt.Errorf("replica presented with unknown operation")
	}
}
