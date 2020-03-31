package parser

// structure of the parser
type prser struct {
	fileName string // name of the file the parser will operate on
}

func NewParser(fnm string) interface{} {
	return &prser{
		fileName: fnm,
	}
}

func (p *prser) Process() error {
	return nil
}
