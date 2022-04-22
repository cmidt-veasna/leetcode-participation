package day9

func countTrailZero(a, b int) (int, int) {
	if a > truncate {
		a %= truncate
	}
	if b > truncate {
		b %= truncate
	}
	c, v := 0, a*b
	for ; v%10 == 0 && v > 0; v /= 10 {
		c++
	}
	if v == 0 {
		v = 1
	}
	return c, v
}

type cat struct {
	v        int
	trailing int
}

func (c *cat) countTrail(v int, pTrail int) (int, int) {
	c.trailing, c.v = countTrailZero(v, 1)
	c.trailing += pTrail
	return c.trailing, c.v
}

type cindex struct {
	r, l, u, d *cat // right, left, up, down
}

func newCIndex() *cindex {
	return &cindex{&cat{v: 1}, &cat{v: 1}, &cat{v: 1}, &cat{v: 1}}
}

const truncate = 100000000

func maxTrailingZeros(grid [][]int) int {
	ig := make([][]*cindex, len(grid))

	// build rows
	for i := 0; i < len(grid); i++ {
		ig[i] = make([]*cindex, len(grid[i]))
		s := len(grid[i])
		tl, tr, last := 1, 1, s-1
		tptl, tptr := 0, 0
		for j, l := 0, last; j < s; j, l = j+1, l-1 {
			if j <= l {
				ig[i][j] = newCIndex()
			}
			if l > j {
				ig[i][l] = newCIndex()
			}
			if j > 0 {
				tptl, tl = ig[i][j].l.countTrail(tl, tptl)
			}
			if l < last {
				tptr, tr = ig[i][l].r.countTrail(tr, tptr)
			}
			if tl > truncate {
				tl %= truncate
			}
			if tr > truncate {
				tr %= truncate
			}
			tl *= grid[i][j]
			tr *= grid[i][l]
		}
	}

	// build columns
	for j := 0; j < len(grid[0]); j++ {
		s := len(grid)
		tu, td, last := 1, 1, s-1
		tptu, tptd := 0, 0
		for i, l := 0, last; i < s; i, l = i+1, l-1 {
			if i > 0 {
				tptu, tu = ig[i][j].u.countTrail(tu, tptu)
			}
			if l < last {
				tptd, td = ig[l][j].d.countTrail(td, tptd)
			}
			if tu > truncate {
				tu %= truncate
			}
			if td > truncate {
				td %= truncate
			}
			tu *= grid[i][j]
			td *= grid[l][j]
		}
	}

	max := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			ci := ig[i][j]
			lv, lvt := ci.l.v, ci.l.trailing
			rv, rvt := ci.r.v, ci.r.trailing
			uv, uvt := ci.u.v, ci.u.trailing
			dv, dvt := ci.d.v, ci.d.trailing
			// left-up
			mi, _ := countTrailZero(grid[i][j], lv*uv)
			mi += lvt + uvt
			if mi > max {
				max = mi
			}
			// left-down
			mi, _ = countTrailZero(grid[i][j], lv*dv)
			mi += lvt + dvt
			if mi > max {
				max = mi
			}
			// right-up
			mi, _ = countTrailZero(grid[i][j], rv*uv)
			mi += rvt + uvt
			if mi > max {
				max = mi
			}
			// right-down
			mi, _ = countTrailZero(grid[i][j], rv*dv)
			mi += rvt + dvt
			if mi > max {
				max = mi
			}
			// left-right
			mi, _ = countTrailZero(grid[i][j], rv*lv)
			mi += rvt + lvt
			if mi > max {
				max = mi
			}
			// up-down
			mi, _ = countTrailZero(grid[i][j], uv*dv)
			mi += uvt + dvt
			if mi > max {
				max = mi
			}
		}
	}

	return max
}
