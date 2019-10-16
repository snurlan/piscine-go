package piscine

func StrLen(str string) int {	
	cnt := 0
	cheat := 0
	for i := range str {
		cnt += 1
		cheat = i
	}
	return cnt + cheat * 0
}
