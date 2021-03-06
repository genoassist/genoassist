// constants holds all the constants shared between packages
package constants

import (
	"fmt"
	"path"

	"github.com/genoassist/config_parser"
)

const (
	// dockerURL denotes the source URL for the Docker image repository
	dockerURL = "docker.io"

	// GenoAssist represents the name of this program
	GenoAssist = "genoassist"

	// BaseOut is the base directory used as the output, mounted by Docker
	BaseOut = "/output"

	// MegaHit specific constants
	MegaHit          = "megahit"
	MegaHitOut       = GenoAssist + "_megahit_out"
	MegaHitDockerURL = dockerURL + "/vout/megahit" // https://github.com/voutcn/megahit

	// Abyss specific constants
	Abyss          = "abyss"
	AbyssOut       = GenoAssist + "_abyss_out"
	AbyssDockerURL = dockerURL + "/bcgsc/abyss" // https://github.com/bcgsc/abyss

	// Flye specific constants
	Flye          = "flye"
	FlyeOut       = GenoAssist + "_fly_out"
	FlyeDockerURL = dockerURL + "nanozoo/flye" // https://github.com/fenderglass/Flye"

	// CreateDir informs the condition function to create a directory for an assembler
	CreateDir = "CreateDir"
)

type (
	// getAssemblerCommand returns the Docker container command associated with an assembler. The commands expects the
	// number of thread to be specified, as assemblers can run on multiple threads
	getAssemblerCommand func(config *config_parser.Config) []string

	// Condition that is run by the condition command
	Condition string

	// AssemblerDetails holds the details of each assembler
	AssemblerDetails struct {
		Name              string              // assembler name
		DHubURL           string              // DockerHub url of the assembler image
		OutputDir         string              // output directory where to read assembled sequences from, no longer bound to Docker
		AssemblyFileName  string              // name of the resulting assembly file
		Comm              getAssemblerCommand // function to return the Docker command of the assembler
		ConditionsPresent bool                // whether there are any special functions to run
		Conditions        []Condition         // list of conditions to run before an assembler runs
	}
)

var (
	// AvailableAssemblers defines the structs of currently integrated assemblers
	AvailableAssemblers = map[string]*AssemblerDetails{
		MegaHit: {
			Name:             MegaHit,
			DHubURL:          MegaHitDockerURL,
			OutputDir:        MegaHitOut,
			AssemblyFileName: "final.contigs.fa",
			Comm: func(cfg *config_parser.Config) []string {
				var finalCmd = []string{
					fmt.Sprintf("--read=%s", InputTarget[cfg.GenoAssist.FileType]),
					fmt.Sprintf("--out-dir=%s", path.Join(BaseOut, MegaHitOut)),
				}

				if cfg.GenoAssist.Threads != 0 {
					finalCmd = append(finalCmd, fmt.Sprintf("--num-cpu-threads=%d", cfg.GenoAssist.Threads))
				}

				return finalCmd
			},
			ConditionsPresent: false,
		},
		Abyss: {
			Name:             Abyss,
			DHubURL:          AbyssDockerURL,
			OutputDir:        AbyssOut,
			AssemblyFileName: "final-contigs.fa",
			Comm: func(cfg *config_parser.Config) []string {

				var finalCmd = []string{
					"name=final",
					fmt.Sprintf("in=%s", InputTarget[cfg.GenoAssist.FileType]),
					fmt.Sprintf("--directory=%s", path.Join(BaseOut, AbyssOut)),
				}

				if cfg.Assemblers.Abyss.KMers != "" {
					finalCmd = append(finalCmd, fmt.Sprintf("k=%s", cfg.Assemblers.Abyss.KMers))
				} else {
					finalCmd = append(finalCmd, fmt.Sprintf("k=25"))
				}

				if cfg.GenoAssist.Threads != 0 {
					finalCmd = append(finalCmd, fmt.Sprintf("j=%d", cfg.GenoAssist.Threads))
				}

				finalCmd = append(finalCmd, "contigs")
				return finalCmd
			},
			ConditionsPresent: true,
			Conditions: []Condition{
				CreateDir,
			},
		},
		Flye: {
			Name:             Flye,
			DHubURL:          FlyeDockerURL,
			OutputDir:        FlyeOut,
			AssemblyFileName: "assembly.fasta",
			Comm: func(cfg *config_parser.Config) []string {

				finalCommand := []string{
					"flye",
					"--genome-size", cfg.Assemblers.Flye.GenomeSize,
					"--out-dir", path.Join(BaseOut, FlyeOut),
				}

				if cfg.GenoAssist.Threads != 0 {
					finalCommand = append(finalCommand, "--threads")
					finalCommand = append(finalCommand, fmt.Sprintf("%d", cfg.GenoAssist.Threads))
				}

				if cfg.Assemblers.Flye.SeqType == "nano" {
					finalCommand = append(finalCommand, "--nano-raw")
				} else {
					finalCommand = append(finalCommand, "--pacbio-raw")
				}
				finalCommand = append(finalCommand, InputTarget[cfg.GenoAssist.FileType])
				return finalCommand
			},
			ConditionsPresent: false,
			Conditions:        nil,
		},
	}
)

// GetDockerURL returns the DockerHub URL of the assembler
func (a *AssemblerDetails) GetDockerURL() string {
	return a.DHubURL
}
