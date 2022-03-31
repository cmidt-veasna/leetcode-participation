package day4

import "sort"

func maximumTop(nums []int, k int) int {
	switch l := len(nums); {
	case k == 0, l == 2 && k == 2, l == 1 && k%2 == 0:
		return nums[0]
	case l == 1:
		return -1
	case k == 1:
		return nums[k]
	case k == l:
		sort.Ints(nums[:k-1])
		return nums[k-2]
	case k < l:
		var n, m int
		if k%2 == 0 {
			sort.Ints(nums[:k-1])
			n, m = nums[k-2], nums[k]
		} else {
			sort.Ints(nums[:k])
			n, m = nums[k-1], nums[k]
		}
		if n < m {
			return m
		} else {
			return n
		}
	default:
		sort.Ints(nums)
		return nums[len(nums)-1]
	}
}
