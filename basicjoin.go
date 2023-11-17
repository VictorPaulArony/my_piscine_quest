package piscine

func BasicJoin(elems []string) string {
	runes := ""
	for _, elems := range elems {
		runes += elems
	}
	return runes
}
