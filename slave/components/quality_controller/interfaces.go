package quality_controller

type Controller interface {
	Process() (string, error)
}
