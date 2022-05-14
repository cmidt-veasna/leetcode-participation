package day11

import (
	"sort"
)

func countDistinct(nums []int, k int, p int) int {
	unique := make([]int, 201)
	one := 0
	divisable := make([]int, 0, 100)
	for i, v := range nums {
		if v%p == 0 {
			divisable = append(divisable, i)
		}
		if unique[v] == 0 {
			unique[v]++
			one++
		}
	}

	count := 0
	log := make([][]int, 0, 10)
	for c := len(nums); c > 1; c-- {
		log = log[:0]
		for i, j := 0, c; j <= len(nums); i, j = i+1, j+1 {
			ii := sort.SearchInts(divisable, i)
			ij := sort.SearchInts(divisable, j-1)
			if ij < len(divisable) && divisable[ij] == j-1 {
				ij++
			}
			if ij-ii > k {
				continue
			}
			diffc := len(log)
			if len(log) > 0 {
				for _, vl := range log {
					for iv, v := range nums[i:j] {
						if vl[iv] != v {
							diffc--
							break
						}
					}
				}
			}
			log = append(log, nums[i:j])

			if diffc == 0 {
				count++
			}
		}
	}
	return one + count
}
