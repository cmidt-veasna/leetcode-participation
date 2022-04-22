package day9

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var minimumCases = [][][]int{
	{{2, 2, 3, 3, 2, 4, 4, 4, 4, 4}, {4}},
	{{2, 3, 3}, {-1}},
	{{5, 5, 5, 5}, {2}},
	{
		{66, 66, 63, 61, 63, 63, 64, 66, 66, 65, 66, 65, 61, 67, 68, 66, 62, 67, 61, 64, 66, 60, 69, 66, 65, 68, 63, 60, 67, 62, 68, 60, 66, 64, 60, 60, 60, 62, 66, 64, 63, 65, 60, 69, 63, 68, 68, 69, 68, 61},
		{20},
	},
}

func TestMinimumRounds(t *testing.T) {
	for i, tc := range minimumCases {
		t.Logf("TestMinimumRounds case #%d", i+1)
		result := minimumRounds(tc[0])
		require.Equal(t, tc[1][0], result)
	}
}
