// contains the definition of an rst
package result

// TODO: Implement this so that it is consistent with what processor returns and reporter accepts

// rst represents the result that the parser function returns and is used by reporting component of genomagic
type rst struct {
}

// New creates a new result struct and returns it
func New(msg string, err error, est int) Result {
	return &rst{}
}

func (r rst) GetMessage() string {
	panic("implement me")
}

func (r rst) GetError() error {
	panic("implement me")
}

func (r rst) GetExitStatusCode() {
	panic("implement me")
}
