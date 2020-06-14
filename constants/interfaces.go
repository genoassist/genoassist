// constants holds all the constants shared between packages
package constants

// Details is an interface that provides a consistent way of accessing the DockerHub information for Docker images such as
// the assembler or quality control ones
type Details interface {
	// GetDockerURL returns the DockerHub URL of the struct implementing this interface
	GetDockerURL() string
}
