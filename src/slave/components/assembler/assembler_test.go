package assembler

import (
	"fmt"
	"testing"
)

func TestNewAssembler(t *testing.T) {
	asm, err := NewAssembler("hello")
	if err != nil {
		fmt.Printf("err in new")
		return
	}
	if err := asm.Process(); err != nil {
		fmt.Println("err")
	}
}