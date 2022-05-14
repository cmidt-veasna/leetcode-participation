package day10

import "math"

func countLatticePoints(circles [][]int) int {
	var log [201][201]int
	var count = 0
	for _, ifo := range circles {
		r := float64(ifo[2])
		xlo, xma := -ifo[0], ifo[0]+1
		ylo, yma := -ifo[1], ifo[1]+1
		for i := xlo; i < xma; i++ {
			for j := ylo; j < yma; j++ {
				hp := math.Sqrt(float64(i*i + j*j))
				if hp <= r {
					x, y := i+ifo[0], j+ifo[1]
					if log[x][y] == 0 {
						log[x][y]++
						count++
					}
				}
			}
		}
	}
	return count
}
