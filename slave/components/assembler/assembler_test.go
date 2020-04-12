package assembler

import (
	"fmt"
	"testing"
)

// manually testing New and Process for now
func TestNewAssembler(t *testing.T) {
	a, err := NewAssembler("", MegaHit)
	if err != nil {
		fmt.Printf(fmt.Sprintf("err: %v", err))
	}
	err = a.Process()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
