package day6

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var deleteCases = [][]int{
	{0},
	{1, 1, 2},
	{1, 1},
	{1, 2, 1, 1},
	{1, 1, 1, 3},
	{1, 1, 2, 3, 5, 1},
	{1, 1, 2, 2, 3, 3, 2},
	{1, 3, 4, 5, 3, 4, 5, 1},
	{1, 1, 1, 1, 3, 1, 1, 2, 2, 5},
}

func TestMinDeletion(t *testing.T) {
	for i, tc := range deleteCases {
		t.Logf("TestMinDeletion case #%d", i+1)
		result := minDeletion(tc[:len(tc)-1])
		require.Equal(t, tc[len(tc)-1], result)
	}
}
