// constants holds all the constants shared between packages
package constants

const (
	// Canu is the name of the program GenoAssist uses to perform error correction
	Canu          = "canu"
	CanuDockerURL = dockerURL + "/greatfireball/canu" // https://github.com/greatfireball/ime_canu"

	// Porechop is the name of the program GenoAssist uses to perform adapter trimming
	Porechop          = "porechop"
	PorechopDockerURL = dockerURL + "/replikation/porechop" // https://github.com/rrwick/porechop"
)

type (
	// QualityControllerDetails represents the information associated with a quality controller
	QualityControllerDetails struct {
		// dHubURL is the DockerHub URL for the image of the quality controller
		dHubURL string
	}
)

var (
	// AvailableQualityControllers defines the mapping of quality controller names to their associated details
	AvailableQualityControllers = map[string]*QualityControllerDetails{
		Canu: {
			dHubURL: CanuDockerURL,
		},
		Porechop: {
			dHubURL: PorechopDockerURL,
		},
	}
)

// GetDockerURL returns the DockerHub URL of the quality controller
func (q *QualityControllerDetails) GetDockerURL() string {
	return q.dHubURL
}
