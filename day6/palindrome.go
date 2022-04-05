package day6

func palin(mm, q, bitl int) (res int64) {
	step := bitl / 2
	var i, s, cc = 0, 1, 10
	if bitl%2 != 0 {
		res += int64((q % 10) * mm)
		q = q / 10
		s = 10
		cc = 1
	}
	for ; i < step; i, s = i+1, s*10 {
		a, b := mm*s, mm/(s*cc)
		if b == 0 {
			b = 1
		}
		c := q % 10
		if i+1 == step {
			c += 1
		}
		res += int64(a*c + b*c)
		q = q / 10
	}
	return
}

func kthPalindrome(queries []int, intLength int) []int64 {
	result := make([]int64, len(queries))
	for i := 0; i < len(queries); i++ {
		result[i] = -1
	}

	if intLength == 1 {
		for i, q := range queries {
			if q < 10 {
				result[i] = int64(q)
			}
		}
		return result
	}

	var max, mm = 1, 1
	if intLength%2 == 0 {
		half := intLength / 2
		for i := 0; i < half; i++ {
			max *= 10
		}
		mm = max
		if max > 10 {
			max = (max / 10) * 9
		} else {
			max--
		}
	} else {
		half := intLength / 2
		for i := 0; i < half; i++ {
			max *= 10
		}
		mm = max
		max *= 9
	}
	for i, q := range queries {
		if q <= max {
			result[i] = palin(mm, q-1, intLength)
		}
	}

	return result
}
