// constants holds all the constants shared between packages
package constants

const (
	MegaHit = "megahit"
)

// getAssemblerCommand returns the Docker container command associated with an assembler
type getAssemblerCommand func(i, o string) []string

// assemblerDetails holds the details of each assembler
type AssemblerDetails struct {
	Name    string              // assembler name
	DHubURL string              // DockerHub url of the assembler image
	Comm    getAssemblerCommand // function to return the Docker command of the assembler
}

// AvailableAssemblers defines the structs of currently integrated assemblers
var AvailableAssemblers = map[string]*AssemblerDetails{
	MegaHit: &AssemblerDetails{
		Name:    MegaHit,
		DHubURL: "docker.io/vout/megahit", // https://github.com/voutcn/megahit
		Comm: func(i, o string) []string {
			return []string{"-r","/raw_sequence_input.fastq", "-o", "/output/genomagic-megahit_output"}
		},
	},
}
