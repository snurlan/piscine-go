package piscine

func Join(strs []string, sep string) string {
	res := ""
	flag := true
	for _, s := range strs {
		if flag {
			flag = false
		} else {
			res += sep
		}
		res += s
	}
	return res
}
