package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	firstPrint := true
	for d1 := '0'; d1 <= '9'; d1 += 1 {

		if n == 1 {
			if firstPrint {
				firstPrint = false
			} else {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			z01.PrintRune(d1)
			continue
		}

		for d2 := d1 + 1; d2 <= '9'; d2 += 1 {
			if n == 2 {
				if firstPrint {
					firstPrint = false
				} else {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
				z01.PrintRune(d1)
				z01.PrintRune(d2)
				continue
			}

			for d3 := d2 + 1; d3 <= '9'; d3 += 1 {
				if n == 3 {
					if firstPrint {
						firstPrint = false
					} else {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
					z01.PrintRune(d1)
					z01.PrintRune(d2)
					z01.PrintRune(d3)
					continue
				}

				for d4 := d3 + 1; d4 <= '9'; d4 += 1 {
					if n == 4 {
						if firstPrint {
							firstPrint = false
						} else {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
						z01.PrintRune(d1)
						z01.PrintRune(d2)
						z01.PrintRune(d3)
						z01.PrintRune(d4)
						continue
					}

					for d5 := d4 + 1; d5 <= '9'; d5 += 1 {
						if n == 5 {
							if firstPrint {
								firstPrint = false
							} else {
								z01.PrintRune(',')
								z01.PrintRune(' ')
							}
							z01.PrintRune(d1)
							z01.PrintRune(d2)
							z01.PrintRune(d3)
							z01.PrintRune(d4)
							z01.PrintRune(d5)
							continue
						}

						for d6 := d5 + 1; d6 <= '9'; d6 += 1 {
							if n == 6 {
								if firstPrint {
									firstPrint = false
								} else {
									z01.PrintRune(',')
									z01.PrintRune(' ')
								}
								z01.PrintRune(d1)
								z01.PrintRune(d2)
								z01.PrintRune(d3)
								z01.PrintRune(d4)
								z01.PrintRune(d5)
								z01.PrintRune(d6)
								continue
							}

							for d7 := d6 + 1; d7 <= '9'; d7 += 1 {
								if n == 7 {
									if firstPrint {
										firstPrint = false
									} else {
										z01.PrintRune(',')
										z01.PrintRune(' ')
									}
									z01.PrintRune(d1)
									z01.PrintRune(d2)
									z01.PrintRune(d3)
									z01.PrintRune(d4)
									z01.PrintRune(d5)
									z01.PrintRune(d6)
									z01.PrintRune(d7)
									continue
								}
								for d8 := d7 + 1; d8 <= '9'; d8 += 1 {
									if n == 8 {
										if firstPrint {
											firstPrint = false
										} else {
											z01.PrintRune(',')
											z01.PrintRune(' ')
										}
										z01.PrintRune(d1)
										z01.PrintRune(d2)
										z01.PrintRune(d3)
										z01.PrintRune(d4)
										z01.PrintRune(d5)
										z01.PrintRune(d6)
										z01.PrintRune(d7)
										z01.PrintRune(d8)
										continue
									}
									for d9 := d8 + 1; d9 <= '9'; d9 += 1 {
										if n == 9 {
											if firstPrint {
												firstPrint = false
											} else {
												z01.PrintRune(',')
												z01.PrintRune(' ')
											}
											z01.PrintRune(d1)
											z01.PrintRune(d2)
											z01.PrintRune(d3)
											z01.PrintRune(d4)
											z01.PrintRune(d5)
											z01.PrintRune(d6)
											z01.PrintRune(d7)
											z01.PrintRune(d8)
											z01.PrintRune(d9)
											continue
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	z01.PrintRune('\n')
}
