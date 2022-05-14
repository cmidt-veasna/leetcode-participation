package day11

func removeDigit(number string, digit byte) string {
	sb := []byte(number)
	lasti := -1 // smalless index to the right side
	for i1, i2 := 0, 1; i1 < len(sb); i1, i2 = i2, i2+1 {
		if sb[i1] == digit {
			if i2 == len(sb) || sb[i1] < sb[i2] {
				sb = append(sb[:i1], sb[i2:]...)
				return string(sb)
			}
			lasti = i1
		}
	}
	if lasti == -1 {
		return number
	} else {
		sb = append(sb[:lasti], sb[lasti+1:]...)
		return string(sb)
	}
}
