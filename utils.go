package main

// IsAlpha Function to check if a string contains only letters
func IsAlpha(s string) bool {
	for _, char := range s {
		if !('a' <= char && char <= 'z') && !('A' <= char && char <= 'Z') && char != ' ' {
			return false
		}
	}
	return true
}
