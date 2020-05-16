package quality_controller

// Controller represents the
type Controller interface {
	Process() (string, error)
}
