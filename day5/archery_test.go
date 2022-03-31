package day5

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var archeryCase = [][][]int{
	{{9}, {1, 1, 0, 1, 0, 0, 2, 1, 0, 1, 2, 0}, {47}},
	{{3}, {0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 2}, {27}},
	{{6}, {0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1}, {30}},
	{{8}, {1, 0, 1, 0, 0, 0, 1, 1, 1, 1, 1, 1}, {39}},
	{{3}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 0}, {26}},
	{{9}, {0, 0, 0, 0, 0, 2, 0, 5, 1, 1, 0, 0}, {51}},
	{{1}, {0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0}, {11}},
	{{6}, {0, 0, 1, 1, 2, 1, 0, 0, 0, 0, 1, 0}, {45}},
	{{7}, {0, 0, 1, 1, 3, 0, 0, 0, 0, 2, 0, 0}, {48}},
}

func TestMaximumBobPoints(t *testing.T) {
	for i, tc := range archeryCase {
		t.Logf("TestMaximumBobPoints case #%d", i+1)
		result := maximumBobPoints(tc[0][0], tc[1])
		bs, bc, as := 0, 0, 0
		for i := 0; i < 12; i++ {
			if tc[1][i] > 0 && tc[1][i] >= result[i] {
				as += i
			}
			if result[i] > 0 {
				bs += i
				bc += result[i]
			}
		}
		require.Equal(t, tc[0][0], bc)
		require.GreaterOrEqual(t, bs, as, "bob arrows:", result)
		require.Equal(t, tc[2][0], bs)
	}
}
