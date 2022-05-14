package day11

func appealSumA(s string) int64 {
	if len(s) == 1 {
		return 1
	}

	last := make([]int, 26)
	for i := 0; i < 26; i++ {
		last[i] = -1
	}

	result := int64(0)

	for i := 0; i < len(s); i++ {
		last[s[i]-'a'] = i
		result += 26 * int64(i+1)
		for j := 0; j < 26; j++ {
			if last[j] == -1 {
				result -= int64(i + 1)
			} else {
				result -= int64(i - last[j])
			}
		}
	}

	return result
}

func appealSum(s string) int64 {
	if len(s) == 1 {
		return 1
	}
	result := int64(0)
	log := make([]int, 26)

	for i, c := range s {cd
		var l int
		var v = c-'a'
		if log[v] != 0 {
			l = i - log[v] + 1
		} else {
			l = i + 1
		}
		result += int64(l * (len(s) - i))
		log[v] = i + 1
	}
	return result
}
