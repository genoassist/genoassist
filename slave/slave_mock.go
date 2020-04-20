package slave

import (
	"github.com/stretchr/testify/mock"
)

// mockSlv is a slave mock
type mockSlv struct {
	mock.Mock
	description string            // name/description of the work performed by the slave
	filePath    string            // the file the slave is supposed to perform work on
	workType    ComponentWorkType // the type of work that has to be performed by the slave
	fail        bool              // whether to fail the worker process
}

// NewMock creates and returns a new instance of a slave
func NewMock(dsc, fnm string, wtp ComponentWorkType, fail bool) *mockSlv {
	return &mockSlv{
		description: dsc,
		filePath:    fnm,
		workType:    wtp,
	}
}

// Process mocks the original slave process function
func (s *mockSlv) Process() error {
	_ = s.Called(s.workType, s.filePath, s.workType, s.fail)
	// TODO: Need to find another way to test this, left commented for now.
	//wrkr := WorkType[s.workType]
	//if wrkr == nil {
	//	return fmt.Errorf("failed to initialize worker")
	//}
	//if s.fail {
	//	return fmt.Errorf("slave process failed")
	//}
	return nil
}
