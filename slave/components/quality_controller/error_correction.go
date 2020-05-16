package quality_controller

type ErrorCorrection struct{}

func NewErrorCorrection() Controller {
	return &ErrorCorrection{}
}

func (e *ErrorCorrection) Process() (string, error) {
	return "", nil
}
