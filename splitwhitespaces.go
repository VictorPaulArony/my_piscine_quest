package piscine

// Write a function that separates the words of a string and puts them in a string slice.
// The separators are spaces, tabs and newlines.
func SplitWhiteSpaces(s string) []string {
	words := []string{}
	word := ""
	for _, char := range s {
		if char == ' ' || char == '\t' || char == '\n' {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(char)
		}
	}
	if word != "" {
		words = append(words, word)
	}
	return words
}
