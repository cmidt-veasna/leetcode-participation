package day5

func maximumBobPoints(numArrows int, aliceArrows []int) []int {
	bobArrows := make([]int, 12)
	var bs, p = 0, 0

	for i := 0; i < 4096; i++ {
		bna, score := 0, 0
		for j := 0; j < 12; j++ {
			if i>>j&1 == 0 {
				bna += aliceArrows[j] + 1
				score += j
			}
		}
		if bna <= numArrows && bs <= score {
			bs, p = score, i
		}
	}

	for i := 0; i < 12; i++ {
		if p>>i&1 == 0 {
			bobArrows[i] = aliceArrows[i] + 1
			numArrows -= bobArrows[i]
		}
	}
	bobArrows[11] += numArrows

	return bobArrows
}
