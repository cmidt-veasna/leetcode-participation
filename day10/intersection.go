package day10

import "sort"

func intersection(nums [][]int) []int {
	result := make([]int, 0, 100)
	log := make([]int, 1001)
	for _, o := range nums {
		for _, v := range o {
			log[v]++
			if log[v] >= len(nums) {
				result = append(result, v)
			}
		}
	}
	sort.Ints(result)
	return result
}
