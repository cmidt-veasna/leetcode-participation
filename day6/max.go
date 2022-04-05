package day6

func maxValueOfCoins(piles [][]int, k int) int {
	if k == 1 {
		max := piles[0][0]
		for _, pile := range piles {
			if max < pile[0] {
				max = pile[0]
			}
		}
		return max
	} else if k == 2000 {
		max := 0
		for _, pile := range piles {
			for _, coin := range pile {
				max += coin
			}
		}
		return max
	}

	sums := make([][]int, len(piles))
	for i, pile := range piles {
		sums[i] = make([]int, len(pile))
		s := 0
		for j, coin := range pile {
			s += coin
			sums[i][j] = s
		}
	}

	var f func(spiles [][]int, pile, remain int, record [][]int) int
	f = func(spiles [][]int, pile, remain int, record [][]int) int {
		if pile >= len(spiles) || remain <= 0 {
			return 0
		}
		if record[pile][remain] != -1 {
			return record[pile][remain]
		}

		mcoin := 0
		maxi := remain
		if remain > len(spiles[pile]) {
			maxi = len(spiles[pile])
		}
		for i := 0; i <= maxi; i++ {
			cc := f(spiles, pile+1, remain-i, record)
			if i > 0 {
				cc += spiles[pile][i-1]
			}
			if mcoin < cc {
				mcoin = cc
			}
		}
		record[pile][remain] = mcoin
		return mcoin
	}

	records := make([][]int, len(piles))
	for i := 0; i < len(piles); i++ {
		records[i] = make([]int, k+1)
		for j := 0; j <= k; j++ {
			records[i][j] = -1
		}
	}

	return f(sums, 0, k, records)
}
