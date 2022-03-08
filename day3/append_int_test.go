package day3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testAppendInt = [][][]int{
	{
		[]int{5, 6}, []int{6}, []int{25},
	},
	{
		[]int{1, 4, 25, 10, 25}, []int{2}, []int{5},
	},
	{
		[]int{96, 44, 99, 25, 61, 84, 88, 18, 19, 33, 60, 86, 52, 19, 32, 47, 35, 50, 94, 17, 29, 98, 22, 21, 72, 100, 40, 84},
		[]int{35}, []int{794},
	},
}

func init() {
	a := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		a[i] = i + 1
	}
	testAppendInt = append(testAppendInt, [][]int{
		a, {10}, {1000055},
	})
}

func TestAppendInt(t *testing.T) {
	for i, tc := range testAppendInt {
		t.Logf("TestAppendInt case #%d", i+1)
		require.Equal(t, int64(tc[2][0]), minimalKSum(tc[0], tc[1][0]))
	}
}
