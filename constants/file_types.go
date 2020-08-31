package constants

const (
	// NOTE: the FASTQ and FASTA constants are duplicated from config_parser. There would be a circular import if these
	// would be defined here only. The config parser can be refactored to isolate the Config struct. Currently,
	// that is the reason why config_parser is imported by this package

	// FASTQ is the FASTQ input type
	FASTQ = "fastq"
	// FASTA is the FASTA input type
	FASTA = "fasta"
)

var (
	// InputTarget defines the mapping of the input file type to a file that is mounted to Docker when running
	// assembly and parsing processes
	InputTarget = map[string]string{
		FASTA: "/raw_sequence_input.fasta",
		FASTQ: "/raw_sequence_input.fastq",
	}
)
