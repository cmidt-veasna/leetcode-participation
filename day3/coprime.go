package day3

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func replaceNonCoprimes(nums []int) []int {
	for i, j := 0, 1; j < len(nums); {
		a, b := nums[i], nums[j]
		ii := i
		if cc := gcd(a, b); cc > 1 {
			tmp := append(nums[:i], lcm(a, b))
			nums = append(tmp, nums[j+1:]...)
		} else {
			i, j = i+1, j+1
		}
	last:
		for ia, ja := ii-1, ii; ia >= 0; ia, ja = ia-1, ia {
			a, b = nums[ia], nums[ja]
			if gcd(a, b) > 1 {
				tmp := append(nums[:ia], lcm(a, b))
				nums = append(tmp, nums[ja+1:]...)
				i, j = i-1, j-1
			} else {
				break
			}
		}
		if j == len(nums) {
			ii = i
			j++ // for break out of loop
			goto last
		}
	}
	return nums
}