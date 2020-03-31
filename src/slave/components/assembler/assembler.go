// contains the definition and work associated with assemblers
package assembler

// structure of the assembler
type asmbler struct {
	fileName string // name of the file the assembler will operate on
}

func NewAssembler(fnm string) interface{} {
	return &asmbler{
		fileName: fnm,
	}
}

func (a *asmbler) Process() error {
	return nil
}
