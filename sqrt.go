package piscine

func Sqrt(nb int) int {
	for i := 2; i*i <= nb; i++ {
		if i*i == nb {
			return i
		}
	}
	return 0
}
