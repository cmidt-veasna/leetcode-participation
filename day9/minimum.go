package day9

func minimumRounds(tasks []int) int {
	if len(tasks) == 0 {
		return 0
	} else if len(tasks) == 1 {
		return -1
	}
	var counts = make(map[int]int)
	for _, v := range tasks {
		counts[v]++
	}
	round := 0
	for _, c := range counts {
		switch {
		case c%3 != 1:
			round += c/3 + (c%3)/2
		case c > 3 && (c-3*(c/3-1))%2 == 0:
			ca := (c/3 - 1)
			round += ca + (c-(3*ca))/2
		case c%2 == 0:
			round += c / 2
		default:
			return -1
		}
	}
	return round
}
