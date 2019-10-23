package piscine

func ConcatParams(args []string) string {
	res := ""
	isFirst := true
	for _, arg := range args {
		if isFirst {
			isFirst = false
		} else {
			res += "\n"
		}
		res += arg
	}
	return res
}
