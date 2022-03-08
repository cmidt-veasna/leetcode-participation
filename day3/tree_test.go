package day3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type treeInputTest struct {
	input  [][3]int
	result *TreeNode
}

var testTree = []treeInputTest{
	{
		[][3]int{{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1}},
		&TreeNode{
			Val: 50,
			Left: &TreeNode{
				Val:   20,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 17},
			},
			Right: &TreeNode{
				Val:  80,
				Left: &TreeNode{Val: 19},
			},
		},
	},
	{
		[][3]int{{1, 2, 1}, {2, 3, 0}, {3, 4, 1}},
		&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 4},
				},
			},
		},
	},
	{
		[][3]int{{39, 70, 1}, {13, 39, 1}, {85, 74, 1}, {74, 13, 1}, {38, 82, 1}, {82, 85, 1}},
		&TreeNode{
			Val: 38,
			Left: &TreeNode{
				Val: 82,
				Left: &TreeNode{
					Val: 85,
					Left: &TreeNode{
						Val: 74,
						Left: &TreeNode{
							Val: 13,
							Left: &TreeNode{
								Val:  39,
								Left: &TreeNode{Val: 70},
							},
						},
					},
				},
			},
		},
	},
}

func TestTree(t *testing.T) {
	for i, tc := range testTree {
		t.Logf("TestTree case #%d", i+1)
		require.Equal(t, tc.result, createBinaryTree(tc.input))
	}
}
