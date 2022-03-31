package day4

import (
	"sort"
)

type traverse func(n *node) [][2]int

type node struct {
	from  [][2]int
	to    [][2]int
	n     int
	index int
}

func (n *node) updateFrom(in []int) { n.from = minimal(n.from, in[0], in[2]) }
func (n *node) updateTo(in []int)   { n.to = minimal(n.to, in[1], in[2]) }

func minimal(s [][2]int, in, w int) [][2]int {
	index := sort.Search(len(s), func(i int) bool { return s[i][0] >= in })
	if index >= len(s) {
		s = append(s, [2]int{in, w})
	} else if s[index][0] != in {
		s = append(s, [2]int{-1, -1})
		copy(s[index+1:], s[index:])
		s[index][0], s[index][1] = in, w
	} else if s[index][1] > w {
		s[index][1] = w
	}
	return s
}

func forward(n *node) [][2]int { return n.to }
func reverse(n *node) [][2]int { return n.from }

// func (n *node) reverse(nodes []*node, src int, hw int, ws []int) []int {
// 	if n.index > 0 {
// 		// loop back to node which have been visited
// 		return ws
// 	}

// 	if hw > 0 && (hw < ws[n.n] || ws[n.n] == 0) {
// 		ws[n.n] = hw
// 	}

// 	if n.n == src {
// 		return ws
// 	}

// 	origin := n.index
// 	for i := n.index; i < len(n.from); i++ {
// 		n.index++
// 		ni, w := n.from[i][0], n.from[i][1]
// 		if ni == n.n {
// 			continue
// 		}
// 		ws = nodes[ni].reverse(nodes, src, w+hw, ws)
// 	}
// 	n.index = origin
// 	return ws
// }

// func (n *node) trace(nodes []*node, dest int, hw int, ws []int) []int {
// 	if n.index > 0 {
// 		// loop back to node which have been visited
// 		return ws
// 	}

// 	if hw > 0 && (hw < ws[n.n] || ws[n.n] == 0) {
// 		ws[n.n] = hw
// 	}

// 	if n.n == dest {
// 		return ws
// 	}

// 	origin := n.index
// 	for i := n.index; i < len(n.to); i++ {
// 		n.index++
// 		ni, w := n.to[i][0], n.to[i][1]
// 		if ni == n.n {
// 			continue
// 		}
// 		ws = nodes[ni].trace(nodes, dest, w+hw, ws)
// 	}
// 	n.index = origin
// 	return ws
// }

func (n *node) lookup(nodes []*node, src int, hw int, f traverse, ws []int) []int {
	if n.index > 0 {
		// loop back to node which have been visited
		return ws
	}

	if hw > 0 && (hw < ws[n.n] || ws[n.n] == 0) {
		ws[n.n] = hw
	}

	if n.n == src {
		return ws
	}

	s := f(n)
	origin := n.index
	for i := n.index; i < len(s); i++ {
		n.index++
		ni, w := s[i][0], s[i][1]
		if ni == n.n {
			continue
		}
		ws = nodes[ni].lookup(nodes, src, w+hw, f, ws)
	}
	n.index = origin
	return ws
}

func makeArray(size int, val int, mod ...int) []int {
	a := make([]int, size)
	if mod != nil {
		copy(a, mod)
	} else {
		for i := 0; i < size; i++ {
			a[i] = val
		}
	}
	return a
}

func minimumWeight(n int, edges [][]int, src1 int, src2 int, dest, iii int) int64 {
	if src1 == src2 && dest == src1 {
		return 0
	}

	nodes := make([]*node, n)
	for _, inf := range edges {
		ns := nodes[inf[0]]
		if ns == nil {
			ns = &node{n: inf[0]}
			nodes[inf[0]] = ns
		}
		ns.updateTo(inf)
		if ns = nodes[inf[1]]; ns == nil {
			ns = &node{n: inf[1]}
			nodes[inf[1]] = ns
		}
		ns.updateFrom(inf)
	}

	if nodes[dest] == nil || len(nodes[dest].from) == 0 || nodes[src1] == nil || nodes[src2] == nil {
		return -1
	}

	dws := makeArray(n, 1e13)
	sws1 := makeArray(n, 1e13, dws...)
	sws2 := makeArray(n, 1e13, dws...)

	dws[dest], sws1[src1], sws2[src2] = 0, 0, 0

	dws = nodes[dest].lookup(nodes, src1, 0, reverse, dws)
	sws1 = nodes[src1].lookup(nodes, dest, 0, forward, sws1)
	sws2 = nodes[src2].lookup(nodes, dest, 0, forward, sws2)

	max := int64(1e13)
	weights := max
	for i := 0; i < n; i++ {
		if cc := int64(sws1[i] + sws2[i] + dws[i]); cc < weights {
			weights = cc
		}
	}

	if weights == max {
		return -1
	}
	return weights
}
