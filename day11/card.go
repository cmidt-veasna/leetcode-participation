package day11

func minimumCardPickup(cards []int) int {
	var m = make(map[int][]int)
	var min = 100001
	for i, v := range cards {
		vv, ok := m[v]
		if !ok {
			vv = []int{-1, -1}
			m[v] = vv
		}
		switch {
		case vv[0] == -1:
			vv[0] = i
			continue
		case vv[1] == -1:
			vv[1] = i
		default:
			vv[0], vv[1] = vv[1], i
		}
		if npick := vv[1] - vv[0] + 1; min > npick {
			min = npick
		}
	}
	if min == 100001 {
		return -1
	} else {
		return min
	}
}
