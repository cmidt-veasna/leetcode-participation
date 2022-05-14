package day11

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type distinctCase struct {
	nums []int
	k    int
	p    int
	out  int
}

var distinctCases = []*distinctCase{
	{
		[]int{2, 3, 3, 2, 2},
		2,
		2,
		11,
	},
	{
		[]int{1, 2, 3, 4},
		4,
		1,
		10,
	},
	{
		[]int{1, 2, 3, 3, 3, 4, 3, 6, 7},
		3,
		5,
		41,
	},
	{
		[]int{1, 2, 3, 3, 3, 4, 3, 6, 7},
		3,
		3,
		30,
	},
	{
		[]int{13, 4, 14, 13, 15, 4, 8, 13, 4, 12},
		5,
		14,
		50,
	},
	{
		[]int{19, 10, 1, 9, 18, 5, 19, 1, 9},
		4,
		8,
		41,
	},
}

func init() {
	a := make([]int, 200)
	for i := 0; i < 200; i++ {
		a[i] = i + 1
	}
	distinctCases = append(distinctCases, &distinctCase{a, 10, 1, 1955})
	distinctCases = append(distinctCases, &distinctCase{a, 10, 4, 7460})
}

func TestCountDistinct(t *testing.T) {
	for i, tc := range distinctCases {
		t.Logf("TestCountDistinct case #%d", i+1)
		result := countDistinct(tc.nums, tc.k, tc.p)
		require.Equal(t, tc.out, result)
	}
}
