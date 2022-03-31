package day5

import "sort"

type indexed struct {
	i     int
	ch    byte
	count int
}

type indexes struct {
	ixs  []*indexed
	size int
}

func newIndexes(s string) (is *indexes, max *indexed) {
	is = &indexes{make([]*indexed, 0, len(s)), len(s)}
	var index *indexed
	var tmp = byte(0)
	for i := 0; i < len(s); i++ {
		if tmp == s[i] {
			index.count++
		} else {
			if index != nil && (max == nil || max.count < index.count) {
				max = index
			}
			tmp = s[i]
			index = &indexed{i, tmp, 1}
			is.ixs = append(is.ixs, index)
		}
	}
	if max == nil || max.count < index.count {
		max = index
	}
	return is, max
}

func (is *indexes) find(fi int) (int, *indexed) {
	pi := sort.Search(len(is.ixs), func(i int) bool {
		return is.ixs[i].i >= fi || (is.ixs[i].i+is.ixs[i].count-1) >= fi
	})
	return pi, is.ixs[pi]
}

func (is *indexes) insert(i, v int, s byte, count int) int {
	ii := sort.Search(len(is.ixs), func(i int) bool {
		return is.ixs[i].i >= v || (is.ixs[i].i+is.ixs[i].count-1) >= v
	})
	if ii == len(is.ixs) {
		is.ixs = append(is.ixs, &indexed{v, s, count})
	} else {
		is.ixs = append(is.ixs, nil)
		copy(is.ixs[ii+1:], is.ixs[ii:])
		is.ixs[ii] = &indexed{v, s, count}
	}
	return ii
}

func (is *indexes) max() (m *indexed) {
	if len(is.ixs) > 0 {
		m = is.ixs[0]
		for i := 1; i < len(is.ixs); i++ {
			if m.count < is.ixs[i].count {
				m = is.ixs[i]
			}
		}
	}
	return
}

func (is *indexes) update(v int, s byte) *indexed {
	if 0 <= v && v < is.size {
		i, ix := is.find(v)
		if ix.count == 1 {
			if l, ok := is.mergeLeft(i, s); ok {
				if r, ok := is.mergeRight(i-1, s); ok {
					return r
				}
				return l
			} else if r, ok := is.mergeRight(i, s); ok {
				return r
			} else {
				ix.ch = s
			}
		} else if ix.i == v {
			ix.i++
			ix.count--
			is.insert(i, v, s, 1)
			if l, ok := is.mergeLeft(i, s); ok {
				return l
			}
		} else if ix.i+ix.count-1 == v {
			ix.count--
			is.insert(i, v, s, 1)
			if r, ok := is.mergeRight(i+1, s); ok {
				return r
			}
		} else {
			oc := is.ixs[i].count
			c1 := v - is.ixs[i].i
			c2 := oc - c1 - 1
			is.ixs[i].count = c1
			i2 := is.insert(0, v+1, is.ixs[i].ch, c2)
			is.insert(0, v, s, 1)
			if c1 > c2 {
				return is.ixs[i]
			} else {
				return is.ixs[i2]
			}
		}
	}
	return nil
}

func (is *indexes) mergeLeft(v int, s byte) (*indexed, bool) {
	if a := v - 1; a >= 0 && is.ixs[a].ch == s {
		is.ixs[a].count++
		if v == len(is.ixs)-1 {
			is.ixs = is.ixs[:v]
		} else {
			is.ixs = append(is.ixs[:v], is.ixs[v+1:]...)
		}
		return is.ixs[a], true
	}
	return nil, false
}

func (is *indexes) mergeRight(v int, s byte) (*indexed, bool) {
	if b := v + 1; b < len(is.ixs) && is.ixs[b].ch == s {
		if is.ixs[v].ch != s {
			is.ixs[b].i--
			is.ixs[b].count++
			is.ixs = append(is.ixs[:v], is.ixs[b:]...)
			return is.ixs[v], true
		} else {
			is.ixs[v].count += is.ixs[b].count
			if b == len(is.ixs)-1 {
				is.ixs = is.ixs[:b]
			} else {
				is.ixs = append(is.ixs[:b], is.ixs[b+1:]...)
			}
			return is.ixs[v], true
		}
	}
	return nil, false
}

func longestRepeating(s string, queryCharacters string, queryIndices []int) []int {
	res := make([]int, 0)
	is, max := newIndexes(s)

	sii := []byte(s)
	for i, v := range queryIndices {
		if sii[v] != queryCharacters[i] {
			oc := max.count
			ix := is.update(v, queryCharacters[i])
			if ix != nil && max.count < ix.count {
				max = ix
			} else if oc > max.count {
				max = is.max()
			}
			sii[v] = queryCharacters[i]
		}
		res = append(res, max.count)
	}

	return res
}
