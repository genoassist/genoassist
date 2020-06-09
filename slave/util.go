package slave

import (
	"fmt"
)

// StrSlice a slide of string
type StrSlice []string

// IsUserRequestedAssembler checks if list
func IsUserRequestedAssembler(list StrSlice, a string) bool {
	for _, b := range list {
		if b == a {
			fmt.Println("DEBUG: ", a)
			return true
		}
	}
	return false
}
