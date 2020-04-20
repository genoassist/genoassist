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
	OutputDir string
	Comm    getAssemblerCommand // function to return the Docker command of the assembler
}

// AvailableAssemblers defines the structs of currently integrated assemblers
// TODO: Need to find a better way to incorporate the "genomagic-megahit_output" so that both OutputDir and return statement can make use of this
var AvailableAssemblers = map[string]*AssemblerDetails{
	MegaHit: &AssemblerDetails{
		Name:    MegaHit,
		DHubURL: "docker.io/vout/megahit", // https://github.com/voutcn/megahit
		OutputDir: "genomagic-megahit_output",
		Comm: func(i, o string) []string {
			// Runs megahit program with following flags:
			//	-r raw_sequence_input.fastq
			//	-o /output/genomagic-megahit_output
			// NOTE: input filePath and outPath are mapped to docker during creation
			return []string{"-r", "/raw_sequence_input.fastq", "-o", "/output/genomagic-megahit_output" }
		},
	},
}
