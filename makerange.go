package piscine

func MakeRange(min, max int) []int {
	n := max - min
	if n <= 0 {
		return nil
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = min + i
	}
	return res
}
