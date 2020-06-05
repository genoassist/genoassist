package slave

import (
	"fmt"
)

// StrSlice a slide of string
type StrSlice []string

// StringInSlice checks if a string exists in a StrSlice.
func StringInSlice(list StrSlice, a string) bool {
	for _, b := range list {
		if b == a {
			fmt.Println("DEBUG: ", a)
			return true
		}
	}
	return false
}
