package day2

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func countPairs(nums []int, k int) int64 {
	possibleDivFrag := make(map[int]int)
	divisableIndividual := int64(0)
	pairDivisable := int64(0)
	for i := 0; i < len(nums); i++ {
		if nums[i]%k == 0 {
			// nums[i] is divisable with K, any multiple will produce the same result.
			divisableIndividual++
			continue
		} else if h := gcd(k, nums[i]); h == 1 {
			// it only pair with first case nums[i]%k == 0
			continue
		} else {
			// only those that produce gcd of k/h will consider as divisable
			// when multipling the pair
			h = k / h
			pairDivisable += int64(possibleDivFrag[h])
			for j := 1; true; j++ {
				if j*j == nums[i] {
					possibleDivFrag[j]++
					break
				} else if j*j < nums[i] {
					if nums[i]%j == 0 {
						possibleDivFrag[nums[i]/j]++
						possibleDivFrag[j]++
					}
				} else {
					break
				}
			}
		}
	}
	// a pair between (x,y) where is x is divisable with k and y does not
	a := ((int64(len(nums)) - divisableIndividual) * divisableIndividual)
	// b pair between (x,y) where x and y is both divisable with k
	b := (divisableIndividual-1)*divisableIndividual/2
	return a + b + pairDivisable
}
