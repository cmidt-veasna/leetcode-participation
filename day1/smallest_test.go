package day1

import (
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

type swapInfo struct {
	num    int64
	pos    [2]int
	result int64
}

var swapCase = []*swapInfo{
	/* case 1 */ {num: 2010, pos: [2]int{0, 2}, result: 1020},
	/* case 2 */ {num: 231, pos: [2]int{0, 2}, result: 132},
	/* case 3 */ {num: 23, pos: [2]int{0, 1}, result: 32},
	/* case 4 */ {num: 20104, pos: [2]int{0, 2}, result: 10204},
	/* case 5 */ {num: 221113, pos: [2]int{0, 1}, result: 221113},
	/* case 6 */ {num: 221113, pos: [2]int{4, 1}, result: 211123},
}

func TestSwap(t *testing.T) {
	for i, tc := range swapCase {
		t.Logf("TestSwap case #%d", i+1)
		ns := newNumSort(tc.num)
		ns.Swap(tc.pos[0], tc.pos[1])
		assert.Equal(t, tc.result, ns.num)
	}
}

var smallestTestCases = [][2]int64{
	/* case 1 */ {12, 12},
	/* case 2 */ {2713618, 1123678},
	/* case 3 */ {183, 138},
	/* case 4 */ {103, 103},
	/* case 5 */ {1203, 1023},
	/* case 6 */ {120901, 100129},
	/* case 7 */ {1870901, 1001789},
	/* case 8 */ {71, 17},
	/* case 9 */ {6510, 1056},
	/* case 10 */ {6016510, 1001566},
	/* case 11 */ {6016512, 1012566},
	/* case 12 */ {60262, 20266},
	/* case 13 */ {-12, -21},
	/* case 14 */ {-134012, -432110},
	/* case 15 */ {-1034012, -4321100},
	/* case 15 */ {-402, -420},
	/* case 16 */ {28943955507, 20345557899},
	/* case 17 */ {794235010051, 100012345579},
	/* case 18 */ {10, 10},
	/* case 19 */ {30, 30},
}

func TestSmallestNumber(t *testing.T) {
	for i, tc := range smallestTestCases {
		t.Logf("TestSmallestNumber case #%d", i+1)
		result := smallestNumber(tc[0])
		assert.Equal(t, tc[1], result)
	}
}

func TestSmallestNumberSort(t *testing.T) {
	for i, tc := range smallestTestCases {
		t.Logf("TestSmallestNumber case #%d", i+1)
		result := smallestNumberSort(tc[0])
		assert.Equal(t, tc[1], result)
	}
}

func TestSmallestRandom(t *testing.T) {
	for i := 0; i < 20000; i++ {
		a := rand.Int63n(int64(math.Pow10(2 + rand.Intn(13))))
		exp := smallestNumberArray(a)
		t.Logf("TestSmallestNumber case #%d value: %d expected %d", i+1, a, exp)
		result := smallestNumber(a)
		assert.Equal(t, exp, result)
	}
}

func TestSmallestNumberArray(t *testing.T) {
	for i, tc := range smallestTestCases {
		t.Logf("TestSmallestNumberArray case #%d", i+1)
		result := smallestNumberArray(tc[0])
		assert.Equal(t, tc[1], result)
	}
}

func TestSmallestNumberString(t *testing.T) {
	for i, tc := range smallestTestCases {
		t.Logf("TestSmallestNumberString case #%d", i+1)
		result := smallestNumberString(tc[0])
		assert.Equal(t, tc[1], result)
	}
}

var (
	smallNumI64  = int64(6301)
	mediumNumI64 = int64(82153551)
	largeNumI64  = int64(484611666145821)
)

func BenchmarkNumberSmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		smallestNumber(smallNumI64)
	}
}

func BenchmarkNumberMedium(b *testing.B) {
	for i := 0; i < b.N; i++ {
		smallestNumber(mediumNumI64)
	}
}

func BenchmarkNumberLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		smallestNumber(largeNumI64)
	}
}

func BenchmarkNumberArraySmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		smallestNumberArray(smallNumI64)
	}
}

func BenchmarkNumberArrayMedium(b *testing.B) {
	for i := 0; i < b.N; i++ {
		smallestNumberArray(mediumNumI64)
	}
}

func BenchmarkNumberArrayLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		smallestNumberArray(largeNumI64)
	}
}

func BenchmarkNumberStringSmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		smallestNumberString(smallNumI64)
	}
}

func BenchmarkNumberStringMedium(b *testing.B) {
	for i := 0; i < b.N; i++ {
		smallestNumberString(mediumNumI64)
	}
}

func BenchmarkNumberStringLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		smallestNumberString(largeNumI64)
	}
}

func BenchmarkNumberSortSmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		smallestNumberSort(smallNumI64)
	}
}

func BenchmarkNumberSortMedium(b *testing.B) {
	for i := 0; i < b.N; i++ {
		smallestNumberSort(mediumNumI64)
	}
}

func BenchmarkNumberSortgLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		smallestNumberSort(largeNumI64)
	}
}
