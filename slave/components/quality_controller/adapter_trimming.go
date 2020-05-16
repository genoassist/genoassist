package quality_controller

type AdapterTrimming struct {
}

func New() Controller {
	return &AdapterTrimming{}
}

func (a *AdapterTrimming) Process() (string, error) {
	return "", nil
}
