package constants

const (
	// Canu is the name of the program GenoMagic uses to perform error correction
	Canu = "canu"

	// Trimmomatic is the name of the program GenoMagic uses to perform adapter trimming
	Trimmomatic = "trimmomatic"
)

type (
	// QualityControllerDetails represents the information associated with a quality controller
	QualityControllerDetails struct {
		// dHubURL is the DockerHub URL for the image of the quality controller
		dHubURL string
	}
)

var (
	AvailableQualityControllers = map[string]*QualityControllerDetails{
		Canu: {
			dHubURL: "dockerhub.io/greatfireball/canu",
		},
		Trimmomatic: {
			dHubURL: "dockerhub.io/replikation/porechop",
		},
	}
)

// GetDockerHubURL returns the DockerHub URL of the quality controller
func (q *QualityControllerDetails) GetDockerHubURL() string {
	return q.dHubURL
}
