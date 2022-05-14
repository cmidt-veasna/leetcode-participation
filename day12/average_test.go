package day12

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type averageCase struct {
	in  []int
	out int
}

var averageCases = []*averageCase{
	{
		[]int{4, 8, 5, 0, 1, -1, 6},
		5,
	},
	{
		[]int{1},
		1,
	},
}

func arrayToNodeTree(pos int, in []int) (node *TreeNode) {
	if pos < len(in) && in[pos] != -1 {
		node = &TreeNode{Val: in[pos]}
		cp1, cp2 := pos*2+1, pos*2+2
		node.Left = arrayToNodeTree(cp1, in)
		node.Right = arrayToNodeTree(cp2, in)
	}
	return node
}

func TestAverageOfSubtree(t *testing.T) {
	for i, tc := range averageCases {
		t.Logf("TestAverageOfSubtree case #%d", i+1)
		result := averageOfSubtree(arrayToNodeTree(0, tc.in))
		require.Equal(t, tc.out, result)
	}
}
