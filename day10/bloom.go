package day10

import (
	"sort"
)

type so [][]int

func (s so) Len() int { return len(s) }
func (s so) Less(i, j int) bool {
	return s[i][0] < s[j][0] || (s[i][0] == s[j][0] && s[i][1] < s[j][1])
}
func (s so) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func fullBloomFlowersA(flowers [][]int, persons []int) []int {
	sort.Sort(so(flowers))
	s, sm, e, em := 1000000000, 0, 0, 1000000000
	for i := 0; i < len(flowers); i++ {
		s = min(s, flowers[i][0])
		sm = max(sm, flowers[i][0])
		e = max(e, flowers[i][1])
		em = min(em, flowers[i][1])
	}

	result := make([]int, len(persons))
	if s == sm && e == em && s < e {
		for i := 0; i < len(persons); i++ {
			result[i] = len(flowers)
		}
	} else {
		log := make(map[int]int)
		for pi, pt := range persons {
			if pt < s || pt > e {
				result[pi] = 0
				continue
			}
			if c, ok := log[pt]; ok {
				result[pi] = c
				continue
			}

			// looking for early start before person visit
			i, j := 0, len(flowers)
			for i < j {
				h := int(uint(i+j) >> 1)
				if flowers[h][0] <= pt {
					i = h + 1
				} else {
					j = h
				}
			}
			i--

			c := 0
			for ; i >= 0; i-- {
				if pt <= flowers[i][1] {
					c++
				}
			}

			log[pt] = c
			result[pi] = c
		}
	}

	return result
}

func fullBloomFlowers(flowers [][]int, persons []int) []int {
	var ss = make([]int, 0, 100)
	var inf = 1000000001

	for i, p := range persons {
		flowers = append(flowers, []int{p, inf + i})
	}

	sort.Sort(so(flowers))

	result := make([]int, len(persons))

	for _, flower := range flowers {
		for len(ss) > 0 && ss[0] < flower[0] {
			ss = ss[1:]
		}
		if flower[1] >= inf {
			result[flower[1]-inf] = len(ss)
		} else {
			ss = append(ss, flower[1])
		}
	}

	return result
}
