package day10

import "sort"

type itcc [][]int

func (t itcc) Len() int           { return len(t) }
func (t itcc) Less(i, j int) bool { return t[i][1] <= t[j][1] }
func (t itcc) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }

type rectO struct {
	xaxis []int
	ny    *rectO
}

func countRectangles(rectangles [][]int, points [][]int) []int {
	yaxis := make([]*rectO, 101)
	maxx, maxy := 0, 0
	var ro *rectO
	sort.Sort(itcc(rectangles))
	for _, rect := range rectangles {
		if yaxis[rect[1]] == nil {
			pro := ro
			ro = &rectO{make([]int, 0), nil}
			yaxis[rect[1]] = ro
			if pro != nil {
				pro.ny = ro
			}
		}
		i := sort.SearchInts(yaxis[rect[1]].xaxis, rect[0])
		yaxis[rect[1]].xaxis = append(yaxis[rect[1]].xaxis, 0)
		aa := yaxis[rect[1]].xaxis
		copy(aa[i+1:], aa[i:])
		aa[i] = rect[0]
		if maxx < rect[0] {
			maxx = rect[0]
		}
		if maxy < rect[1] {
			maxy = rect[1]
		}
	}

	result := make([]int, 0, 100)
	for _, p := range points {
		if p[0] > maxx || p[1] > maxy {
			result = append(result, 0)
			continue
		}

		mrecty := yaxis[p[1]]
		if mrecty == nil {
			// non edge case
			for y := p[1] + 1; y < 101; y++ {
				if mrecty = yaxis[y]; mrecty != nil {
					break
				}
			}
		}

		pc := 0
		for mrecty != nil {
			last := len(mrecty.xaxis) - 1
			blast := -1
			for i := 0; i <= last; i, last = i+1, last-1 {
				if p[0] <= mrecty.xaxis[i] {
					pc += len(mrecty.xaxis) - i
					blast = -1
					break
				}
				if p[0] <= mrecty.xaxis[last] {
					blast = last
				} else {
					// not in any rectangle
					break
				}
			}
			if blast != -1 {
				pc += len(mrecty.xaxis) - blast
			}
			mrecty = mrecty.ny
		}
		result = append(result, pc)
	}
	return result
}
