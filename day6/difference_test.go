package day6

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

type differentCase struct {
	n1  []int
	n2  []int
	out [][]int
}

var diffCases = []*differentCase{
	{[]int{1, 2, 3}, []int{2, 4, 6}, [][]int{{1, 3}, {4, 6}}},
	{[]int{1, 2, 3, 3}, []int{1, 1, 2, 2}, [][]int{{3}, {}}},
	{
		[]int{26, 48, -78, -25, 42, -8, 94, -68, 26}, []int{61, -17},
		[][]int{{26, 48, -78, -25, 42, -8, 94, -68}, {-17, 61}},
	},
	{
		[]int{-80, -15, -81, -28, -61, 63, 14, -45, -35, -10},
		[]int{-1, -40, -44, 41, 10, -43, 69, 10, 2},
		[][]int{{-81, -15, -35, -10, -28, -61, -45, 14, -80, 63}, {-44, -43, -40, -1, 2, 10, 41, 69}},
	},
}

func TestFindDifference(t *testing.T) {
	for i, tc := range diffCases {
		t.Logf("TestFindDifference case #%d", i+1)
		result := findDifference(tc.n1, tc.n2)
		sort.Ints(result[0])
		sort.Ints(result[1])
		sort.Ints(tc.out[0])
		sort.Ints(tc.out[1])
		require.Equal(t, tc.out, result)
	}
}
