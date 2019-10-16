package piscine

func StrLen(str string) int {
	cnt := 0
	for index := range str {
		cnt = index
	}
	return cnt + 1
}
