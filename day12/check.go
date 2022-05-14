package day12

import (
	"sort"
)

func merge(v int, dest, src []int) []int {
	if dest == nil {
		dest = make([]int, 0, len(src))
	}
	for i := 0; i < len(src); i++ {
		cv := v + src[i]
		if cv > 0 {
			continue
		}
		ix := sort.SearchInts(dest, cv)
		if ix == len(dest) {
			dest = append(dest, cv)
		} else if dest[ix] != cv {
			dest = append(dest, 0)
			copy(dest[ix+1:], dest[ix:])
			dest[i] = cv
		}
	}
	return dest
}

func hasValidPath(grid [][]byte) bool {
	var m, n = len(grid), len(grid[0])
	switch {
	case m == n:
		return false
	case m%2 == 0 && n%2 == 0:
		return false
	case m%2 != 0 && n%2 != 0:
		return false
	}

	ll := make([][][]int, m)
	for i := 0; i < m; i++ {
		ll[i] = make([][]int, n)
	}

	var fn func(r, c int) []int
	fn = func(r, c int) []int {
		nr, nc := r+1, c+1
		v := 1
		if grid[r][c] == ')' {
			v = -1
		}

		if nr == m && nc == n {
			if v > 0 {
				return []int{}
			} else {
				return []int{v}
			}
		} else {
			var ps = ll[r][c]
			if ps == nil {
				if nr < m {
					ps = merge(v, ps, fn(nr, c))
				}
				if nc < n {
					ps = merge(v, ps, fn(r, nc))
				}
				ll[r][c] = ps
			}
			return ps
		}
	}

	vv := fn(0, 0)
	return len(vv) > 0 && vv[len(vv)-1] == 0
}
