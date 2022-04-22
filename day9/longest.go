package day9

import (
	"sort"
)

func longestPath(parent []int, s string) int {
	if len(parent) == 1 {
		return 1
	}

	// tree from parent to child
	reverse := make([][]int, len(parent))
	for i := len(parent) - 1; i > 0; i-- {
		p := parent[i]
		reverse[p] = append(reverse[p], i)
	}

	var fn func(node int) (int, int)
	fn = func(parent int) (int, int) {
		if len(reverse[parent]) == 0 {
			// reaching the leaf
			return 1, 1
		}
		max, m1, m2 := 1, 0, 0
		for _, child := range reverse[parent] {
			a, b := fn(child)
			if max < a {
				max = a
			}
			if s[child] != s[parent] {
				if b > m1 {
					m2, m1 = m1, b
				} else if b > m2 {
					m2 = b
				}
			}
		}
		if c := 1 + m1 + m2; c > max {
			return c, 1 + m1
		} else {
			return max, 1 + m1
		}
	}

	max, _ := fn(0)
	return max
}

type childPath struct {
	node int
	max  int
}

type childPaths []*childPath

func (cps childPaths) Len() int           { return len(cps) }
func (cps childPaths) Less(i, j int) bool { return cps[i].max > cps[j].max }
func (cps childPaths) Swap(i, j int)      { cps[i], cps[j] = cps[j], cps[i] }

type cachePath struct {
	child  childPaths // sorted
	parent int        // max path of througth parent node
}

func (cp cachePath) maxExcept(child int) int {
	for i := len(cp.child) - 1; i >= 0; i-- {
		if cp.child[i].node != child {
			return cp.child[i].max
		}
	}
	return 0
}

func (cp *cachePath) addChildPath(child, maxPath int) {
	index := sort.Search(cp.child.Len(), func(i int) bool { return cp.child[i].max >= maxPath })
	cp.child = append(cp.child, nil)
	copy(cp.child[index+1:], cp.child[index:])
	cp.child[index] = &childPath{child, maxPath}
}

func longestPathB(parent []int, s string) int {
	if len(parent) == 1 {
		return 1
	}

	reverse := make([][]int, len(parent))
	for i := len(parent) - 1; i > 0; i-- {
		p := parent[i]
		reverse[p] = append(reverse[p], i)
	}

	caches := make([]cachePath, len(parent))

	var toLeaf func(node int, cp byte, count int) int
	toLeaf = func(node int, cp byte, count int) int {
		if s[node] != cp {
			count++
			if mi := caches[node].maxExcept(0); mi > 0 {
				return mi + count
			}
			max := 0
			for _, c := range reverse[node] {
				m := toLeaf(c, s[node], 0)
				caches[node].addChildPath(c, m)
				if m > max {
					max = m
				}
			}
			return max + count
		}
		return count
	}

	var fn func(node, child int, cp byte, count int) int
	fn = func(node, child int, cp byte, count int) int {
		if s[node] != cp {
			if node == 0 {
				if mi := caches[node].maxExcept(child); mi > 0 {
					return mi + count
				}
				// reach root
				count++
				max := 0
				for _, c := range reverse[node] {
					if c != child && s[c] != s[node] {
						m := toLeaf(c, s[node], 0)
						caches[node].addChildPath(c, m)
						if m > max {
							max = m
						}
					}
				}
				return max + count
			} else {
				count++
				maxParent := caches[node].parent
				if maxParent == 0 {
					maxParent = fn(parent[node], node, s[node], 0)
					caches[node].parent = maxParent
				}
				maxParent += count
				max := caches[node].maxExcept(child)
				if max == 0 {
					for _, c := range reverse[node] {
						if c != child && s[c] != s[node] {
							m := toLeaf(c, s[node], 0)
							caches[node].addChildPath(c, m)
							if m > max {
								max = m
							}
						}
					}

				}
				if max += count; max < maxParent {
					max = maxParent
				}
				return max
			}
		}
		return count
	}

	maxPath := 1
	for last := len(parent) - 1; last > 0; last-- {
		// only leaf is consider the longest
		if len(reverse[last]) == 0 {
			if m := fn(parent[last], last, s[last], 1); m > maxPath {
				maxPath = m
			}
		}
	}
	return maxPath
}
