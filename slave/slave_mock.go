package slave

import (
	"github.com/stretchr/testify/mock"

	"github.com/genomagic/config_parser"
)

// SlaveMock is a slave mock
type Mock struct {
	mock.Mock
	config   *config_parser.Config
	workType ComponentWorkType
}

// NewMock creates and returns a new instance of a slave
func NewMock(config *config_parser.Config, workType ComponentWorkType) *Mock {
	return &Mock{
		config:   config,
		workType: workType,
	}
}

// Process mocks the original slave process function
func (s *Mock) Process() error {
	args := s.Mock.Called()
	return args.Error(0)
}
