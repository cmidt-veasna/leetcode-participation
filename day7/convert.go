package day7

import "strconv"

func convertTime(current string, correct string) int {
	cuh, _ := strconv.ParseInt(current[:2], 10, 64)
	cum, _ := strconv.ParseInt(current[3:], 10, 64)
	coh, _ := strconv.ParseInt(correct[:2], 10, 64)
	com, _ := strconv.ParseInt(correct[3:], 10, 64)

	abs := int64(0)
	min := int64(0)
	if com < cum {
		abs = 1
		min += com + (60 - cum)
	} else {
		min = com - cum
	}

	t := coh - cuh - abs
	for min > 0 {
		if min >= 15 {
			min -= 15
			t++
		} else if min >= 5 {
			min -= 5
			t++
		} else {
			min -= 1
			t++
		}
	}
	return int(t)
}
