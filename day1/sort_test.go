package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testCases = [][][]int{
/*case 1 */	{{1}, {1}},
/*case 2 */	{{2, 1}, {2, 1}},
/*case 3 */	{{2, 3, 1}, {1, 3, 2}},
/*case 4 */	{{5, 1, 2, 3}, {2, 3, 5, 1}},
/*case 5 */	{{5, 1, 2, 3, 6}, {2, 3, 5, 1, 6}},
/*case 6 */	{{5, 1, 2, 3, 2}, {2, 3, 2, 1, 5}},
/*case 7 */	{{5, 1, 2, 3, 1}, {1, 3, 2, 1, 5}},
}

func TestSortEventOdd(t *testing.T) {
	for i, tc := range testCases {
		t.Logf("Test case #%d", i+1)
		result := sortEvenOdd(tc[0])
		assert.Equal(t, tc[1], result)
	}
}
