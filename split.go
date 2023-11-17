package piscine

func Split(s, sep string) []string {
	var result []string
	startIndex := 0
	for i := 0; i < len(s); i++ {
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			result = append(result, s[startIndex:i])
			startIndex = i + len(sep)
			i += len(sep) - 1
		}
	}
	result = append(result, s[startIndex:])
	return result
}
