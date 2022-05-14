package day10

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type intersectionCase struct {
	in  [][]int
	out []int
}

var intersectionCases = []*intersectionCase{
	{
		[][]int{{3, 1, 2, 4, 5}, {1, 2, 3, 4}, {3, 4, 5, 6}},
		[]int{3, 4},
	},
	{
		[][]int{{1, 2, 3}, {4, 5, 6}},
		[]int{},
	},
	{
		[][]int{{7, 34, 45, 10, 12, 27, 13}, {27, 21, 45, 10, 12, 13}},
		[]int{10, 12, 13, 27, 45},
	},
}

func TestIntersection(t *testing.T) {
	for i, tc := range intersectionCases {
		t.Logf("TestIntersection case #%d", i+1)
		result := intersection(tc.in)
		require.Equal(t, tc.out, result)
	}
}
