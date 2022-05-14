package day10

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type latticCase struct {
	in  [][]int
	out int
}

var latticCases = []*latticCase{
	{
		[][]int{{2, 2, 1}},
		5,
	},
	{
		[][]int{{2, 2, 2}, {3, 4, 1}},
		16,
	},
	{
		[][]int{{34, 70, 10}, {28, 89, 6}, {36, 95, 18}, {51, 93, 48}, {13, 50, 8}, {52, 20, 14}, {64, 1, 1}, {45, 90, 17}, {94, 47, 3}, {12, 20, 4}, {46, 1, 1}, {60, 59, 54}, {32, 86, 25}, {41, 51, 15}, {26, 66, 20}, {76, 60, 31}, {95, 56, 27}, {67, 51, 13}, {18, 70, 9}, {87, 63, 32}, {84, 6, 6}, {25, 55, 14}, {11, 74, 2}, {47, 21, 7}, {57, 88, 56}, {60, 4, 1}, {34, 14, 14}, {51, 61, 16}, {39, 38, 1}, {23, 69, 14}, {79, 75, 70}, {95, 80, 10}, {14, 66, 4}, {69, 91, 67}, {95, 18, 11}, {35, 84, 7}, {9, 48, 6}, {13, 72, 3}, {76, 11, 5}, {14, 75, 11}, {8, 94, 1}, {34, 80, 21}, {53, 86, 49}, {32, 15, 3}, {64, 61, 58}, {62, 28, 10}, {67, 22, 20}, {87, 48, 46}, {3, 5, 2}, {7, 48, 4}, {75, 88, 68}, {81, 55, 52}, {95, 74, 63}, {73, 18, 17}, {17, 7, 4}, {57, 38, 23}, {81, 22, 10}, {47, 33, 29}, {58, 33, 31}, {48, 77, 15}, {56, 55, 36}, {45, 77, 11}, {82, 93, 22}, {61, 8, 1}, {97, 58, 38}, {97, 93, 92}, {53, 39, 24}, {65, 25, 23}, {32, 90, 4}, {17, 22, 15}, {40, 98, 32}, {65, 17, 4}, {10, 28, 10}, {10, 65, 6}, {23, 94, 3}, {43, 19, 10}, {33, 90, 5}, {2, 28, 2}, {96, 26, 9}, {99, 98, 32}},
		0,
	},
}

func TestCountLatticePoints(t *testing.T) {
	for i, tc := range latticCases {
		t.Logf("TestCountLatticePoints case #%d", i+1)
		result := countLatticePoints(tc.in)
		require.Equal(t, tc.out, result)
	}
}