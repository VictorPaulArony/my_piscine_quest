package piscine

func Swap(a, b *int) {
	victor := *a
	*a = *b
	*b = victor
}
