package piscine

func BasicJoin(strs []string) string {
	res := ""
	for _, s := range strs {
		res += s
	}
	return res
}
