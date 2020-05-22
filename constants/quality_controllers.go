// constants holds all the constants shared between packages
package constants

const (
	// Canu is the name of the program GenoMagic uses to perform error correction
	Canu = "canu"

	// Porechop is the name of the program GenoMagic uses to perform adapter trimming
	Porechop = "porechop"
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
			dHubURL: "dockerhub.io/greatfireball/canu",
		},
		Porechop: {
			dHubURL: "dockerhub.io/replikation/porechop",
		},
	}
)

// GetDockerHubURL returns the DockerHub URL of the quality controller
func (q *QualityControllerDetails) GetDockerHubURL() string {
	return q.dHubURL
}
