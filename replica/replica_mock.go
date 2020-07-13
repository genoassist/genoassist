package replica

import (
	"github.com/stretchr/testify/mock"

	"github.com/genoassist/config_parser"
)

// Mock is a replica mock
type Mock struct {
	mock.Mock
	config   *config_parser.Config
	workType ComponentWorkType
}

// NewMock creates and returns a new instance of a replica
func NewMock(config *config_parser.Config, workType ComponentWorkType) *Mock {
	return &Mock{
		config:   config,
		workType: workType,
	}
}

// Process mocks the original replica process function
func (s *Mock) Process() error {
	args := s.Mock.Called()
	return args.Error(0)
}
