// constants holds all the constants shared between packages
package constants

const (
	MegaHit = "megahit"
)

// assemblerDetails holds the details of each assembler
type AssemblerDetails struct {
	Name    string // assembler name
	DHubURL string // DockerHub url of the assembler image
}

// AvailableAssemblers defines the structs of currently integrated assemblers
var AvailableAssemblers = map[string]*AssemblerDetails{
	MegaHit: &AssemblerDetails{
		Name:    MegaHit,
		DHubURL: "docker.io/vout/megahit", // https://github.com/voutcn/megahit
	},
}
