package day7

func maximumCandies(candies []int, k int64) int {
	if len(candies) == 1 {
		return int(int64(candies[0]) / k)
	}

	sum := int64(0)
	allEqual := true
	for i := 0; i < len(candies); i++ {
		sum += int64(candies[i])
		if allEqual && int64(candies[i]) != k {
			allEqual = false
		}
	}

	switch {
	case sum < k:
		return 0
	case sum == k:
		return 1
	case allEqual && int64(len(candies)) == k:
		return candies[0]
	}

	max := int64(1)
	highEstimated := sum / k
	for max < highEstimated {
		cc := (max + highEstimated + 1) / 2
		total := int64(0)
		for _, count := range candies {
			total += int64(count) / cc
		}
		if total < k {
			highEstimated = cc - 1
		} else {
			max = cc
		}
	}
	return int(max)
}
