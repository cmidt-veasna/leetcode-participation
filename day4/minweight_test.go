package day4

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type weightCase struct {
	edges  [][]int
	n      int
	src1   int
	src2   int
	dest   int
	output int64
}

var minweightCase = []weightCase{
	{
		[][]int{{0, 1, 1}, {2, 1, 1}}, 3, 0, 1, 2, -1,
	},
	{
		[][]int{{0, 2, 2}, {0, 5, 6}, {1, 0, 3}, {1, 4, 5}, {2, 1, 1}, {2, 3, 3}, {2, 3, 4}, {3, 4, 2}, {4, 5, 1}}, 6, 0, 1, 5, 9,
	},
	{
		[][]int{{4, 7, 24}, {1, 3, 30}, {4, 0, 31}, {1, 2, 31}, {1, 5, 18}, {1, 6, 19}, {4, 6, 25}, {5, 6, 32}, {0, 6, 50}}, 8, 4, 1, 6, 44,
	},
	{
		[][]int{{0, 1, 22}, {0, 5, 11}, {0, 2, 7}, {4, 3, 4}, {5, 3, 48}, {5, 3, 17}, {4, 3, 45}}, 6, 4, 0, 3, 32,
	},
	{
		[][]int{{0, 2, 1}, {0, 3, 1}, {2, 4, 1}, {3, 4, 1}, {1, 2, 1}, {1, 3, 10}}, 5, 0, 1, 4, 3,
	},
	{
		[][]int{{7, 0, 19}, {7, 2, 48}, {7, 8, 31}, {9, 3, 10}, {7, 1, 17}, {9, 6, 40}, {7, 5, 35}, {7, 4, 13}, {7, 4, 46}, {7, 4, 36}, {0, 4, 38}}, 10, 9, 7, 4, -1,
	},
	{
		[][]int{{6, 4, 6}, {6, 1, 26}, {6, 5, 9}, {6, 2, 17}, {6, 3, 36}, {4, 3, 36}, {4, 3, 28}, {4, 3, 19}}, 7, 0, 6, 3, -1,
	},
	{
		[][]int{{9, 3, 133}, {8, 9, 9}, {4, 9, 37}, {12, 16, 32}, {12, 4, 148}, {12, 16, 59}, {17, 15, 67}, {12, 15, 75}, {7, 8, 101}, {11, 16, 10}, {3, 17, 37}, {17, 10, 111}, {1, 17, 2}, {12, 16, 68}, {10, 16, 140}, {8, 13, 110}, {9, 12, 49}, {9, 1, 124}, {1, 13, 51}, {6, 7, 8}, {13, 9, 50}, {11, 12, 33}, {0, 12, 127}, {1, 2, 137}, {6, 1, 130}, {8, 17, 131}, {12, 10, 60}, {11, 9, 110}, {2, 7, 40}, {2, 10, 79}, {11, 5, 140}, {2, 1, 48}, {17, 0, 33}, {15, 16, 101}, {12, 10, 25}, {14, 10, 120}, {6, 2, 97}, {5, 11, 42}, {9, 2, 40}, {9, 6, 126}, {17, 9, 35}, {9, 13, 64}, {1, 12, 129}, {10, 16, 109}, {15, 11, 139}, {7, 5, 112}, {8, 0, 18}, {3, 14, 92}, {3, 7, 21}, {13, 11, 27}, {0, 3, 39}, {4, 11, 73}, {9, 10, 67}, {14, 1, 122}, {11, 2, 60}, {5, 9, 54}, {3, 4, 53}, {15, 10, 125}, {16, 6, 58}, {3, 6, 5}, {3, 9, 6}, {17, 16, 48}, {16, 4, 35}, {3, 11, 88}, {9, 1, 100}, {4, 16, 141}, {5, 3, 23}, {13, 14, 35}, {4, 16, 90}, {5, 12, 85}, {12, 16, 24}, {5, 10, 126}, {5, 14, 47}, {1, 10, 84}, {15, 2, 43}, {9, 15, 30}, {7, 0, 50}, {11, 6, 78}, {2, 1, 118}, {3, 8, 79}, {8, 4, 140}, {0, 8, 116}, {8, 10, 106}, {8, 4, 6}, {3, 16, 62}, {16, 1, 118}, {7, 17, 103}, {15, 3, 31}, {17, 2, 142}, {15, 10, 5}, {12, 0, 95}, {4, 9, 72}, {10, 17, 88}, {6, 14, 102}, {17, 0, 8}, {11, 7, 62}, {0, 5, 48}, {2, 15, 107}, {17, 11, 55}, {8, 12, 129}, {0, 5, 68}, {3, 13, 137}, {6, 16, 19}, {13, 17, 4}, {6, 14, 119}, {9, 11, 48}, {15, 13, 128}, {5, 3, 103}, {16, 0, 48}, {1, 15, 87}, {12, 9, 84}, {17, 6, 109}, {12, 7, 148}, {7, 16, 127}, {16, 12, 129}, {14, 2, 90}, {9, 16, 131}, {1, 17, 24}, {13, 0, 72}, {3, 9, 82}, {5, 4, 141}, {3, 9, 4}, {4, 5, 119}, {6, 2, 10}, {12, 2, 126}, {0, 15, 69}, {16, 8, 22}, {14, 7, 147}, {10, 3, 53}, {11, 15, 12}, {6, 7, 52}},
		18, 15, 3, 17, 68,
	},
}

func TestMinimumWeight(t *testing.T) {
	for i, tc := range minweightCase {
		t.Logf("TestMinimumWeight case #%d", i+1)
		result := minimumWeight(tc.n, tc.edges, tc.src1, tc.src2, tc.dest, i)
		require.Equal(t, tc.output, result)
	}
}