package slave

// contains checks if the source contains the given element
func contains(src []string, el string) bool {
	for _, b := range src {
		if b == el {
			return true
		}
	}
	return false
}
