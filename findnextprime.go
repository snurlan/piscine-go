package piscine

func FindNextPrime(nb int) int {
	for !IsPrime(nb) {
		nb++
	}
	return nb
}
