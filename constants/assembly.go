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

	// Flye specific constants
	Flye    = "flye"
	FlyeOut = GenoMagic + "_fly_out"

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
			DHubURL:          "docker.io/vout/megahit", // https://github.com/voutcn/megahit
			OutputDir:        MegaHitOut,
			AssemblyFileName: "final.contigs.fa",
			Comm: func(cfg *config_parser.Config) []string {
				var finalCmd = []string{
					fmt.Sprintf("-r %s", RawSeqIn),
					fmt.Sprintf("-o %s", path.Join(BaseOut, MegaHitOut)),
				}

				if cfg.GenoMagic.Threads != 0 {
					finalCmd = append(finalCmd, fmt.Sprintf("-t %d", cfg.GenoMagic.Threads))
				}

				return finalCmd
			},
			ConditionsPresent: false,
		},
		Abyss: {
			Name:             Abyss,
			DHubURL:          "docker.io/bcgsc/abyss", // https://github.com/bcgsc/abyss
			OutputDir:        AbyssOut,
			AssemblyFileName: "final-contigs.fa",
			Comm: func(cfg *config_parser.Config) []string {

				var finalCmd = []string{
					" name=final",
					fmt.Sprintf("in=%s", RawSeqIn),
					fmt.Sprintf("--directory=%s", path.Join(BaseOut, AbyssOut)),
				}

				if cfg.Assemblers.Abyss.KMers != "" {
					finalCmd = append(finalCmd, fmt.Sprintf("k=%s", cfg.Assemblers.Abyss.KMers))
				}

				if cfg.GenoMagic.Threads != 0 {
					finalCmd = append(finalCmd, fmt.Sprintf("j=%d", cfg.GenoMagic.Threads))
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
			DHubURL:          "dockerhub.io/nanozoo/flye", // https://github.com/fenderglass/Flye
			OutputDir:        FlyeOut,
			AssemblyFileName: "final.contigs.fa",
			Comm: func(cfg *config_parser.Config) []string {

				finalCommand := []string{
					fmt.Sprintf("--genome-size %s", cfg.Assemblers.Flye.GenomeSize),
					fmt.Sprintf("--out-dir %s", path.Join(BaseOut, FlyeOut)),
				}

				if cfg.GenoMagic.Threads != 0 {
					finalCommand = append(finalCommand, fmt.Sprintf("--threads %d", cfg.GenoMagic.Threads))
				}

				if cfg.Assemblers.Flye.SeqType == "nano" {
					finalCommand = append(finalCommand, fmt.Sprintf("--nano-raw %s", RawSeqIn))
				} else {
					finalCommand = append(finalCommand, fmt.Sprintf("--pacbio-raw %s", RawSeqIn))
				}
				return finalCommand
			},
			ConditionsPresent: false,
			Conditions:        nil,
		},
	}
)

// GetDockerHubURL returns the DockerHub URL of the assembler
func (a *AssemblerDetails) GetDockerHubURL() string {
	return a.DHubURL
}
