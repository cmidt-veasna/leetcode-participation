package day7

func findWinners(matches [][]int) [][]int {
	winners := make([]int, 100001)
	losers := make([]int, 100001)

	for _, match := range matches {
		winners[match[0]]++
		losers[match[1]]++
	}

	absWinners := make([]int, 0, 50)
	oneLosers := make([]int, 0, 50)
	for i := 1; i < 100001; i++ {
		if losers[i] == 0 && winners[i] > 0 {
			absWinners = append(absWinners, i)
		} else if losers[i] == 1 {
			oneLosers = append(oneLosers, i)
		}
	}
	return [][]int{absWinners, oneLosers}
}
