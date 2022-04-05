package day6

import "sort"

func remove(nums []int, i int) []int {
	if len(nums) > 0 {
		if i == 0 {
			return nums[1:]
		} else if i == len(nums)-1 {
			return nums[:i]
		} else {
			return append(nums[:i], nums[i+1:]...)
		}
	}
	return nums
}

func removeAllUnique(nums []int, i int, include bool) []int {
	for k := i + 1; k < len(nums) && nums[k] == nums[i]; {
		nums = remove(nums, k)
	}
	if include {
		return remove(nums, i)
	} else {
		return nums
	}
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	var ss [2000]int
	for i := 0; i < len(nums1); {
		if l := nums1[i] + 1000; ss[l] == 1 {
			nums1 = removeAllUnique(nums1, i, true)
		} else {
			ss[l] = 1
			i++
		}
	}

	sort.Ints(nums2)

	for i := 0; i < len(nums2); {
		if l := nums2[i] + 1000; ss[l] == 1 {
			nums2 = removeAllUnique(nums2, i, true)
			ss[l] = 0
		} else {
			nums2 = removeAllUnique(nums2, i, false)
			i++
		}
	}

	for i := 0; i < len(nums1); {
		if l := nums1[i] + 1000; ss[l] == 0 {
			nums1 = removeAllUnique(nums1, i, true)
		} else {
			nums1 = removeAllUnique(nums1, i, false)
			i++
		}
	}

	return [][]int{nums1, nums2}
}
