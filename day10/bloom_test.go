package day10

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type bloomCase struct {
	flowers [][]int
	persons []int
	out     []int
}

var bloomCases = []*bloomCase{
	{
		[][]int{{1, 6}, {3, 7}, {9, 12}, {4, 13}},
		[]int{2, 3, 7, 11},
		[]int{1, 2, 2, 2},
	},
	{
		[][]int{{1, 10}, {3, 3}},
		[]int{3, 3, 2},
		[]int{2, 2, 1},
	},
}

func init() {
	flowers1 := make([][]int, 50000)
	flowers2 := make([][]int, 50000)
	flowers3 := make([][]int, 50000)
	for s, e, i := 0, 0, 0; i < 50000; i++ {
		flowers1[i] = []int{1000000000, 1000000000}
		flowers2[i] = []int{1, 1000000000}
		flowers3[i] = []int{1 + s, 1000000000 - e}
		s = min(s+1, 1000)
		e = min(e+1, 1000)
	}
	persons1 := make([]int, 50000)
	persons2 := make([]int, 50000)
	persons3 := make([]int, 50000)
	answers1 := make([]int, 50000)
	answers2 := make([]int, 50000)
	answers3 := make([]int, 50000)
	answers4 := make([]int, 50000)
	for i := 0; i < 50000; i++ {
		persons1[i] = i + 1
		answers1[i] = 0
		persons2[i] = 1000000000
		answers2[i] = 50000
		persons3[i] = i + 1000
		answers3[i] = 50000
		answers4[i] = 50000
	}
	answers4[0] = 1000
	bloomCases = append(bloomCases, &bloomCase{
		flowers: flowers1,
		persons: persons1,
		out:     answers1,
	})
	bloomCases = append(bloomCases, &bloomCase{
		flowers: flowers1,
		persons: persons2,
		out:     answers2,
	})
	bloomCases = append(bloomCases, &bloomCase{
		flowers: flowers2,
		persons: persons3,
		out:     answers3,
	})
	bloomCases = append(bloomCases, &bloomCase{
		flowers: flowers3,
		persons: persons3,
		out:     answers4,
	})
}

func TestFullBloomFlowers(t *testing.T) {
	for i, tc := range bloomCases {
		s := time.Now()
		t.Logf("TestFullBloomFlowers case #%d", i+1)
		result := fullBloomFlowers(tc.flowers, tc.persons)
		require.Equal(t, tc.out, result)
		fmt.Println("case", i+1, time.Since(s))
	}
}
