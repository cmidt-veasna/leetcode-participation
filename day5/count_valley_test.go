package day5

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var countHillValleyCase = [][][]int{
	{{2, 4, 1, 1, 6, 5}, {3}},
	{{6, 6, 5, 5, 4, 1}, {0}},
	{{5, 7, 7, 1, 7}, {2}},
	{{35, 03, 32, 9, 57, 78, 47, 99, 59, 49, 88, 1, 27, 51, 73, 92}, {9}},
}

func TestCountHillValley(t *testing.T) {
	for i, tc := range countHillValleyCase {
		t.Logf("TestCountHillValley case #%d", i+1)
		result := countHillValley(tc[0])
		require.Equal(t, tc[1][0], result)
	}
}
