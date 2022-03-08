package day3

import (
	"sort"
)

func minimalKSum(nums []int, k int) int64 {
	sort.Ints(nums)
	sum := int64(0)
	start := 1
	index := 0
	prev := 0
	for k > 0 && start < 1000000001 {
		switch {
		case index >= len(nums) || (start < nums[index] && start > prev):
			// not found
			sum += int64(start)
			k--
			start++
		case start == nums[index]:
			start = nums[index] + 1
			fallthrough
		case start > nums[index]:
			prev = nums[index]
			index++
		default:
			start++
		}
	}
	return sum
}
