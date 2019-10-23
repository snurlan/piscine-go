package piscine

func SplitWhiteSpaces(str string) []string {
	cnt := 0
	flag := false
	for _, r := range str {
		if r == ' ' || r == '\t' || r == '\n' {
			flag = false
		} else if !flag {
			flag = true
			cnt++
		}
	}
	res := make([]string, cnt)
	cnt = 0
	flag = false
	s := ""
	for _, r := range str {
		if r == ' ' || r == '\t' || r == '\n' {
			flag = false
			if s != "" {
				res[cnt-1] = s
				s = ""
			}
		} else if !flag {
			flag = true
			cnt++
			s += string(r)
		} else {
			s += string(r)
		}
	}
	if s != "" {
		res[cnt-1] = s
	}
	return res
}
