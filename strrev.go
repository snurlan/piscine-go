package piscine

func StrRev(s string) string {
	if s != "" {
		return StrRev(s[1:]) + string(s[0])
	} else {
		return ""
	}
}
