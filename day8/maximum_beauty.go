package day8

import (
	"sort"
)

func maximumBeauty(flowers []int, newFlowers int64, target int, full int, partial int) int64 {
	complete, gardens := 0, len(flowers)
one:
	if len(flowers) == 1 {
		if min := newFlowers + int64(flowers[0]); min >= int64(target) {
			cl := int64((complete + 1) * full)
			if cc := int64(complete*full) + int64((target-1)*partial); cc > cl {
				return cc
			} else {

				return cl
			}
		} else {
			return int64(complete*full) + min*int64(partial)
		}
	}

	sort.Ints(flowers)

	// count any completed
	for last := gardens - 1; last >= 0 && flowers[last] >= target; last-- {
		complete++
		flowers = flowers[:last]
	}
	if complete == gardens {
		return int64(complete * full)
	} else if len(flowers) == 1 {
		goto one
	}

	// record minimum flower needed to plant to equal to next graden
	info := make([]int64, len(flowers))
	for i, ll := 0, 1; ll < len(flowers); i, ll = ll, ll+1 {
		info[i] = int64(flowers[ll] - flowers[i])
		if i > 0 {
			info[i] = int64(ll)*info[i] + info[i-1]
		}
	}
	info[len(flowers)-1] = info[len(flowers)-2]

	minValue := func(last, max int, n int64) int {
		planted := info[last]
		mi := flowers[last]
		if last == 0 {
			planted = 0
		} else if planted < n && last < len(flowers)-1 {
			mi = flowers[last+1]
		} else {
			for planted > n {
				m := last / 2
				if info[m] > n {
					if m == 0 {
						return flowers[0] + int(n)
					}
					last = m
				} else {
					cc := m + (last-m)/2
					for info[cc] < n {
						lo := (last - cc) / 2
						if lo == 0 {
							planted = info[cc]
							last = cc + 1
							mi = flowers[last]
							goto result
						}
						cc += lo
					}
					last = cc
				}
				mi = flowers[last+1]
				planted = info[last]
			}
		}
	result:
		ap := (n - planted) / int64(last+1)
		min := mi + int(ap)
		if min >= max {
			return max - 1
		} else {
			return min
		}
	}

	last := len(flowers) - 1
	max := int64(complete * full)
	max += int64(minValue(last, target, newFlowers) * partial)
	stillRemain := true
	for newFlowers > 0 && last >= 0 && stillRemain {
		vi := flowers[last]
		needed := target - vi
		stillRemain = newFlowers >= int64(needed)
		if stillRemain {
			newFlowers -= int64(needed)
			complete++
			last--
		}
		min := int64(0)
		if last >= 0 {
			min = int64(minValue(last, target, newFlowers))
		}
		mi := int64(complete*full) + min*int64(partial)
		if max < mi {
			max = mi
		}
	}
	if complete == gardens {
		mi := (complete-1)*full + (target-1)*partial
		if max < int64(mi) {
			max = int64(mi)
		}
	}
	return max
}
