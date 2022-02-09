package day1

import (
	"sort"
)

type ArrayMetaSlice struct {
	even  bool
	array []int
}

func (am *ArrayMetaSlice) Len() int {
	osize := len(am.array) / 2
	if am.even && len(am.array)%2 == 1 {
		return osize + 1
	} else {
		return osize
	}
}

func (am *ArrayMetaSlice) Less(i, j int) bool {
	if am.even {
		return am.array[i*2] < am.array[j*2]
	} else {
		return am.array[(i*2)+1] > am.array[(j*2)+1]
	}
}

func (am *ArrayMetaSlice) Swap(i, j int) {
	if am.even {
		am.array[i*2], am.array[j*2] = am.array[j*2], am.array[i*2]
	} else {
		am.array[(i*2)+1], am.array[(j*2)+1] = am.array[(j*2)+1], am.array[(i*2)+1]
	}
}

func sortEvenOdd(nums []int) []int {
	if len(nums) == 3 {
		if nums[0] > nums[2] {
			nums[0], nums[2] = nums[2], nums[0]
		}
	} else if len(nums) > 3 {
		ams := &ArrayMetaSlice{array: nums, even: true}
		sort.Sort(ams)
		ams.even = false
		sort.Sort(ams)
	}
	return nums
}
