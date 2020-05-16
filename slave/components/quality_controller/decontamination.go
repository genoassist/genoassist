package quality_controller

type Decontamination struct{}

func NewDecontamination() Controller {
	return &Decontamination{}
}

func (d *Decontamination) Process() (string, error) {
	return "", nil
}
