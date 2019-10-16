package piscine

func SortIntegerTable(table []int) {
	sorted := false
	for !sorted {
		sorted = true
		for i := 0; i+1 < len(table); i += 1 {
			if table[i] > table[i+1] {
				t := table[i]
				table[i] = table[i+1]
				table[i+1] = t
				sorted = false
			}
		}
	}
}
