package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	} else {
		res := IterativeFactorial(nb-1) * nb
		if res < 0 {
			return 0
		}
		return res
	}
}
