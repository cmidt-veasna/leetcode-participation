package day6

func minDeletion(nums []int) int {
	switch len(nums) {
	case 1:
		return 1
	case 0:
		return 0
	default:
		min := 0
		for i, a := 0, 1; a < len(nums); {
			if nums[i] == nums[a] {
				if b := i + 2; b < len(nums) && nums[a] != nums[b] {
					nums = append(nums[:i], nums[a:]...)
					min++
					i = b
				} else {
					nums = append(nums[:a], nums[b:]...)
					min++
					continue
				}
			} else {
				i += 2
			}
			a = i + 1
		}
		if len(nums) > 0 && len(nums)%2 != 0 {
			min++
		}
		return min
	}
}
