package day4

import "sort"

func findKDistantIndices(nums []int, key int, k int) []int {
	indices := make([]int, 0)
	prev := -1
	for i := 0; i < len(nums) && prev < len(nums); i++ {
		if key == nums[i] {
			start := prev
			nexts := i + k
			for m := k; m >= 0; m-- {
				a, b := i-m, i+m
				if a > start {
					indices = append(indices, a)
					if prev < a {
						prev = a
					}
				}
				if a != b && b > start && b < len(nums) {
					indices = append(indices, b)
					if prev < b {
						prev = b
					}
				}
				if b < len(nums) && nums[b] == key && nexts > b {
					nexts = b
				}
			}
			if nexts != i+k {
				i = nexts
			} else {
				i = prev
			}
		}
	}
	sort.Ints(indices)
	return indices
}
