package day7

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type candyCase struct {
	in  []int
	k   int64
	out int
}

var candyCases = []*candyCase{
	{[]int{6}, 6, 1},
	{[]int{5, 8, 6}, 3, 5},
	{[]int{2, 5}, 11, 0},
	{[]int{3, 5, 2, 7, 1, 9, 6, 8}, 11, 3},
	{[]int{2, 40, 68}, 12, 8},
	{[]int{5, 20, 30}, 3, 15},
	{[]int{10, 31}, 5, 7},
	{[]int{1002, 2384}, 3040, 1},
	{[]int{1505777, 58744111}, 199922394, 0},
	{[]int{2, 341505777, 497842394}, 199922394, 4},
	{[]int{3, 83837674, 78724433}, 7, 20959418},
}

func TestMaximumCandies(t *testing.T) {
	for i, tc := range candyCases {
		t.Logf("TestMaximumCandies case #%d", i+1)
		result := maximumCandies(tc.in, tc.k)
		require.Equal(t, tc.out, result)
	}
}
