package day5

func countHillValley(nums []int) int {
	last := len(nums) - 1
	hill, valley := 0, 0
next:
	for i := 1; i < last; i++ {
		left, right := i-1, i+1
		for {
			if nums[i] == nums[left] {
				left--
			} else if nums[i] == nums[right] {
				i++
				right++
			} else {
				break
			}
			if left < 0 || right > last {
				continue next
			}
		}
		if nums[left] > nums[i] && nums[right] > nums[i] {
			hill++
		} else if nums[left] < nums[i] && nums[right] < nums[i] {
			valley++
		}
	}
	return hill + valley
}
