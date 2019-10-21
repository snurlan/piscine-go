package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	} else {
		res := RecursiveFactorial(nb-1) * nb
		if res < 0 {
			return 0
		}
		return res
	}
}
