// constants holds all the constants shared between packages
package constants

import "path"

const (
	GenoMagic = "genomagic"

	BaseOut  = "/output"
	RawSeqIn = "/raw_sequence_input.fastq"

	MegaHit    = "megahit"
	MegaHitOut = GenoMagic + "_megahit_out"

	Abyss    = "abyss"
	AbyssOut = GenoMagic + "_abyss_out"
)

// getAssemblerCommand returns the Docker container command associated with an assembler
type getAssemblerCommand func() []string

// assemblerDetails holds the details of each assembler
type AssemblerDetails struct {
	Name             string              // assembler name
	DHubURL          string              // DockerHub url of the assembler image
	OutputDir        string              // output directory where to read assembled sequences from, no longer bound to Docker
	AssemblyFileName string              // name of the resulting assembly file
	Comm             getAssemblerCommand // function to return the Docker command of the assembler
}

// AvailableAssemblers defines the structs of currently integrated assemblers
var AvailableAssemblers = map[string]*AssemblerDetails{
	MegaHit: {
		Name:             MegaHit,
		DHubURL:          "docker.io/vout/megahit", // https://github.com/voutcn/megahit
		OutputDir:        MegaHitOut,
		AssemblyFileName: "/final.contigs.fa",
		Comm: func() []string {
			// NOTE: input filePath and outPath are mapped to Docker mounts during creation (slave/components/assembler/assembler.go:87)
			return []string{"-r", RawSeqIn, "-o", path.Join(BaseOut, MegaHitOut)}
		},
	},
	Abyss: {
		Name:             Abyss,
		DHubURL:          "docker.io/bcgsc/abyss", // https://github.com/bcgsc/abyss
		OutputDir:        AbyssOut,
		AssemblyFileName: "/final.contigs.fa",
		Comm: func() []string {
			return []string{""}
		},
	},
}
