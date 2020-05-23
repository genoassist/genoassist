// constants holds all the constants shared between packages
package constants

import (
	"fmt"
	"path"

	"github.com/genomagic/config_parser"
)

const (
	// GenoMagic represents the name of this program
	GenoMagic = "genomagic"

	// BaseOut is the base directory used as the output, mounted by Docker
	BaseOut = "/output"

	// RawSeqIn is the mapping of the user-specified reads file to the Docker mount file
	RawSeqIn = "/raw_sequence_input.fastq"

	// MegaHit specific constants
	MegaHit    = "megahit"
	MegaHitOut = GenoMagic + "_megahit_out"

	// Abyss specific constants
	Abyss    = "abyss"
	AbyssOut = GenoMagic + "_abyss_out"

	// CreateDir informs the condition function to create a directory for an assembler
	CreateDir = "CreateDir"
)

type (
	// getAssemblerCommand returns the Docker container command associated with an assembler. The commands expects the
	// number of thread to be specified, as assemblers can run on multiple threads
	getAssemblerCommand func(config config_parser.Config) []string

	// Condition that is run by the condition command
	Condition string

	// AssemblerDetails holds the details of each assembler
	AssemblerDetails struct {
		// Name of the assembler
		Name string
		// DHubURL is the url of the assembler image
		DHubURL string
		// OutputDir is directory where to read assembled sequences from, no longer bound to Docker
		OutputDir string
		// AssemblyFileName is the name of the resulting assembly file
		AssemblyFileName string
		// Comm is a function that returns the Docker command of the assembler
		Comm getAssemblerCommand
		// ConditionsPresent informs whether there are any special functions to run
		ConditionsPresent bool
		// Conditions is a list of conditions to run before an assembler runs
		Conditions []Condition
	}
)

var (
	// AvailableAssemblers defines the structs of currently integrated assemblers
	AvailableAssemblers = map[string]*AssemblerDetails{
		MegaHit: {
			Name:             MegaHit,
			DHubURL:          "docker.io/vout/megahit", // https://github.com/voutcn/megahit
			OutputDir:        MegaHitOut,
			AssemblyFileName: "/final.contigs.fa",
			Comm: func(cfg config_parser.Config) []string {
				// NOTE: input filePath and outPath are mapped to Docker mounts during creation (slave/components/assembler/assembler.go:87)
				return []string{
					"-r", RawSeqIn,
					fmt.Sprintf("-t %d", cfg.GenoMagic.Threads),
					"-o", path.Join(BaseOut, MegaHitOut),
				}
			},
			ConditionsPresent: false,
		},
		Abyss: {
			Name:             Abyss,
			DHubURL:          "docker.io/bcgsc/abyss", // https://github.com/bcgsc/abyss
			OutputDir:        AbyssOut,
			AssemblyFileName: "final-contigs.fa",
			Comm: func(cfg config_parser.Config) []string {
				return []string{
					fmt.Sprintf("k=%s", cfg.Assemblers.Abyss.KMers),
					`name=final`,
					fmt.Sprintf("j=%d", cfg.GenoMagic.Threads),
					fmt.Sprintf("in='%s'", RawSeqIn),
					fmt.Sprintf("--directory=%s", path.Join(BaseOut, AbyssOut)),
					"contigs",
				}
			},
			ConditionsPresent: true,
			Conditions: []Condition{
				CreateDir,
			},
		},
	}
)

// GetDockerHubURL returns the DockerHub URL of the assembler
func (a *AssemblerDetails) GetDockerHubURL() string {
	return a.DHubURL
}
