package day7

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var winnerCases = [][][][]int{
	{
		{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}},
		{{1, 2, 10}, {4, 5, 7, 8}},
	},
	{
		{{2, 3}, {1, 3}, {5, 4}, {6, 4}},
		{{1, 2, 5, 6}, {}},
	},
}

func TestFindWinners(t *testing.T) {
	for i, tc := range winnerCases {
		t.Logf("TestFindWinners case #%d", i+1)
		result := findWinners(tc[0])
		require.Equal(t, tc[1], result)
	}
}
