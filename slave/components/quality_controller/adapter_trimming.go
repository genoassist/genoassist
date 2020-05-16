package quality_controller

type AdapterTrimming struct {
}

func NewAdapterTrimming() Controller {
	return &AdapterTrimming{}
}

func (a *AdapterTrimming) Process() (string, error) {
	return "", nil
}
