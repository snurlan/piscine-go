package piscine

func Swap(a *int, b *int) {
	t := *a
	*a = *b
	*b = t
}
