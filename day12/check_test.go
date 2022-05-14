package day12

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type checkCase struct {
	in  [][]byte
	out bool
}

var checkCases = []*checkCase{
	{ // case 1
		[][]byte{{'(', '(', '('}, {')', '(', ')'}, {'(', '(', ')'}, {'(', '(', ')'}},
		true,
	},
	{ // case 2
		[][]byte{{')', ')'}, {'(', '('}},
		false,
	},
	{ // case 3
		[][]byte{{'(', ')', ')', '(', ')', ')', ')'}},
		false,
	},
	{ // case 4
		[][]byte{
			{'(', ')', ')', '(', '(', '(', '(', ')', ')', '(', ')', '(', ')', '(', '(', '(', '(', ')', '(', ')', '('},
			{'(', '(', ')', ')', '(', ')', ')', ')', '(', ')', '(', ')', '(', '(', ')', '(', '(', '(', '(', '(', ')'},
			{')', ')', '(', ')', ')', '(', '(', ')', '(', '(', ')', '(', ')', ')', '(', ')', ')', '(', '(', ')', ')'},
			{'(', '(', ')', '(', ')', '(', ')', ')', ')', '(', ')', '(', '(', ')', '(', ')', ')', '(', ')', ')', ')'},
			{'(', '(', '(', ')', '(', '(', ')', '(', ')', ')', '(', ')', ')', ')', ')', ')', ')', '(', ')', '(', '('},
			{')', ')', '(', '(', ')', ')', ')', ')', ')', '(', ')', ')', ')', '(', '(', ')', '(', '(', '(', '(', ')'},
			{')', ')', ')', ')', '(', ')', '(', ')', '(', '(', ')', '(', '(', ')', '(', '(', ')', ')', '(', ')', '('},
			{'(', ')', '(', '(', '(', ')', ')', ')', ')', '(', '(', ')', '(', '(', ')', ')', '(', ')', ')', ')', '('},
			{'(', ')', '(', ')', '(', '(', '(', '(', ')', '(', '(', '(', '(', '(', '(', ')', '(', ')', '(', ')', ')'},
			{'(', ')', '(', '(', '(', ')', '(', ')', ')', ')', ')', '(', '(', '(', '(', ')', ')', '(', '(', '(', ')'},
			{'(', '(', ')', '(', ')', ')', '(', ')', '(', ')', ')', ')', ')', ')', '(', ')', '(', ')', ')', ')', '('},
			{')', '(', '(', '(', ')', '(', ')', ')', '(', ')', '(', ')', '(', '(', ')', '(', '(', ')', '(', '(', ')'},
			{'(', ')', '(', ')', ')', '(', '(', ')', '(', ')', '(', ')', ')', ')', '(', '(', '(', '(', ')', '(', ')'},
			{'(', '(', ')', '(', ')', ')', '(', '(', '(', ')', '(', ')', '(', '(', ')', ')', '(', '(', '(', ')', ')'},
			{'(', '(', '(', '(', ')', ')', '(', ')', '(', '(', '(', ')', ')', '(', ')', '(', ')', ')', ')', ')', '('},
			{'(', '(', '(', ')', ')', ')', '(', ')', ')', '(', ')', ')', '(', '(', ')', '(', ')', '(', '(', '(', ')'},
			{')', ')', ')', ')', ')', ')', '(', ')', ')', ')', '(', '(', ')', '(', ')', '(', '(', '(', '(', ')', ')'},
		},
		false,
	},
	{ // case 5
		[][]byte{
			{'(', ')', '(', '('},
			{'(', ')', ')', '('},
			{')', '(', ')', ')'},
			{')', '(', '(', '('},
			{'(', ')', '(', '('},
			{'(', ')', '(', '('},
			{')', ')', '(', ')'},
		},
		true,
	},
	{ // case 6
		[][]byte{
			{'(', ')', '(', '('},
			{'(', ')', ')', '('},
			{')', '(', ')', ')'},
			{')', '(', '(', '('},
			{'(', ')', '(', '('},
			{'(', ')', '(', '('},
			{'(', ')', '(', ')'},
		},
		true,
	},
	{ // case 7
		[][]byte{
			{'(', '(', ')', '(', ')', '(', '(', ')', '(', '(', ')', ')', ')', ')', ')', '(', ')', '(', '(', ')', '(', '(', ')', ')', ')', ')', ')', '(', '(', '(', '('},
			{')', '(', '(', '(', ')', '(', ')', '(', '(', ')', ')', ')', ')', '(', ')', ')', '(', '(', ')', ')', '(', ')', '(', ')', '(', '(', ')', '(', ')', '(', '('},
			{')', ')', '(', '(', ')', '(', '(', ')', ')', ')', ')', '(', '(', ')', ')', '(', ')', '(', ')', ')', '(', '(', '(', ')', ')', ')', '(', ')', ')', '(', ')'},
			{'(', '(', ')', '(', ')', '(', '(', ')', '(', '(', '(', ')', ')', '(', ')', '(', ')', ')', ')', ')', ')', ')', '(', '(', ')', '(', ')', '(', ')', '(', '('},
			{')', ')', '(', ')', ')', '(', '(', '(', ')', ')', '(', ')', '(', ')', ')', ')', '(', '(', '(', ')', ')', '(', ')', '(', ')', ')', '(', '(', '(', '(', ')'},
			{')', ')', '(', '(', ')', '(', ')', '(', ')', '(', ')', '(', ')', ')', '(', ')', '(', ')', ')', '(', ')', '(', '(', '(', ')', '(', ')', ')', ')', '(', '('},
			{')', '(', '(', '(', '(', '(', '(', ')', ')', '(', '(', ')', '(', ')', ')', '(', ')', ')', ')', '(', '(', '(', ')', '(', '(', ')', ')', '(', ')', '(', ')'},
			{')', ')', '(', '(', '(', '(', '(', '(', '(', ')', ')', '(', '(', '(', '(', '(', '(', '(', '(', '(', '(', '(', '(', ')', ')', '(', '(', ')', ')', '(', ')'},
			{'(', ')', ')', ')', '(', '(', ')', ')', ')', ')', '(', ')', ')', '(', ')', ')', '(', '(', '(', '(', '(', '(', '(', ')', ')', '(', '(', ')', ')', '(', '('},
			{'(', '(', ')', '(', ')', ')', ')', ')', '(', '(', '(', ')', ')', ')', '(', ')', '(', '(', ')', '(', '(', '(', ')', '(', '(', '(', '(', '(', ')', ')', ')'},
			{'(', ')', '(', '(', '(', '(', ')', '(', '(', ')', ')', '(', '(', ')', '(', '(', '(', ')', '(', '(', '(', ')', ')', '(', ')', ')', '(', ')', '(', '(', ')'},
			{')', ')', '(', '(', '(', '(', ')', '(', '(', ')', ')', '(', ')', ')', '(', ')', '(', '(', '(', '(', '(', '(', '(', ')', '(', '(', ')', ')', '(', '(', '('},
			{'(', ')', ')', ')', '(', ')', '(', '(', '(', ')', ')', ')', '(', ')', '(', ')', ')', '(', '(', '(', '(', ')', '(', ')', ')', ')', ')', ')', ')', '(', '('},
			{'(', '(', '(', '(', '(', '(', ')', ')', '(', ')', '(', '(', '(', ')', ')', '(', ')', '(', ')', '(', ')', '(', '(', '(', ')', ')', ')', '(', ')', '(', '('},
			{'(', ')', ')', ')', ')', '(', '(', ')', ')', ')', ')', ')', ')', '(', '(', ')', '(', ')', ')', '(', ')', '(', ')', ')', ')', '(', '(', ')', '(', '(', '('},
			{'(', ')', ')', '(', '(', ')', ')', '(', ')', ')', '(', '(', '(', ')', ')', ')', ')', '(', '(', '(', ')', ')', '(', ')', '(', '(', '(', '(', ')', ')', ')'},
			{')', '(', '(', ')', '(', '(', ')', ')', ')', '(', '(', '(', '(', ')', '(', ')', ')', ')', '(', ')', '(', ')', '(', '(', ')', '(', '(', '(', '(', '(', '('},
			{'(', ')', '(', ')', '(', '(', ')', ')', ')', ')', ')', '(', '(', ')', ')', '(', ')', '(', ')', ')', ')', ')', '(', '(', '(', ')', '(', ')', '(', ')', ')'},
			{'(', ')', '(', ')', ')', ')', '(', '(', '(', ')', '(', ')', '(', '(', ')', ')', '(', ')', '(', ')', '(', ')', '(', '(', '(', '(', '(', ')', '(', ')', '('},
			{')', ')', ')', ')', ')', '(', ')', ')', '(', '(', ')', '(', ')', ')', '(', ')', ')', '(', '(', '(', '(', ')', '(', '(', '(', '(', ')', ')', ')', ')', '('},
			{')', '(', '(', '(', '(', '(', ')', '(', ')', ')', ')', ')', '(', '(', '(', ')', ')', '(', ')', ')', '(', '(', '(', '(', '(', '(', ')', ')', '(', '(', '('},
			{'(', '(', '(', ')', ')', '(', ')', '(', ')', ')', ')', ')', '(', ')', ')', ')', ')', '(', ')', '(', '(', '(', '(', ')', '(', '(', '(', '(', '(', '(', ')'},
		},
		false,
	},
	{ // case 8
		[][]byte{
			{'(', ')', ')', '(', ')', ')', '(', '(', ')', ')', '(', '(', ')', ')', ')', ')', '(', '(', '(', ')', '(', ')', ')', ')', ')', '(', '(', '(', ')', ')', '(', ')', '(', '(', '(', ')', ')', ')', ')'},
			{')', ')', '(', ')', ')', '(', '(', '(', '(', ')', ')', '(', ')', '(', ')', ')', ')', '(', ')', ')', ')', ')', '(', '(', ')', ')', ')', '(', ')', '(', ')', '(', ')', ')', ')', '(', ')', '(', ')'},
			{')', ')', ')', ')', ')', ')', ')', ')', ')', ')', ')', '(', '(', ')', ')', ')', '(', ')', ')', ')', ')', ')', '(', '(', '(', ')', ')', ')', '(', '(', '(', ')', ')', ')', '(', ')', ')', ')', '('},
			{')', '(', ')', '(', '(', '(', '(', '(', ')', '(', '(', ')', '(', '(', '(', ')', ')', '(', ')', ')', '(', '(', '(', '(', ')', ')', ')', '(', ')', '(', '(', ')', ')', '(', ')', '(', '(', ')', ')'},
			{'(', '(', ')', '(', '(', ')', ')', ')', ')', ')', '(', '(', '(', '(', '(', '(', ')', '(', ')', '(', '(', '(', '(', '(', ')', ')', ')', '(', '(', ')', ')', ')', ')', '(', '(', '(', ')', ')', '('},
			{')', ')', ')', ')', ')', ')', ')', ')', ')', ')', ')', ')', '(', '(', '(', '(', ')', ')', '(', '(', ')', ')', ')', '(', ')', ')', '(', '(', ')', '(', ')', '(', ')', '(', '(', '(', ')', ')', ')'},
			{'(', '(', ')', ')', ')', '(', ')', '(', '(', '(', ')', ')', '(', '(', '(', ')', ')', '(', '(', '(', ')', ')', '(', ')', ')', ')', ')', ')', '(', '(', ')', '(', ')', ')', ')', ')', '(', '(', '('},
			{')', ')', '(', '(', ')', ')', '(', ')', '(', '(', '(', ')', ')', ')', ')', ')', '(', ')', '(', ')', ')', ')', ')', '(', ')', '(', ')', ')', '(', '(', ')', ')', ')', '(', '(', '(', ')', ')', ')'},
			{'(', ')', '(', '(', '(', ')', ')', ')', ')', ')', ')', ')', '(', ')', ')', ')', '(', '(', '(', '(', ')', '(', '(', '(', ')', ')', '(', '(', '(', ')', '(', '(', ')', ')', ')', '(', ')', ')', '('},
			{'(', '(', ')', '(', '(', '(', '(', '(', ')', ')', '(', ')', ')', '(', '(', '(', '(', '(', '(', '(', ')', '(', '(', '(', '(', ')', '(', '(', ')', '(', ')', '(', '(', '(', ')', ')', ')', '(', ')'},
			{'(', ')', '(', '(', '(', '(', ')', ')', ')', '(', ')', ')', '(', '(', ')', ')', ')', ')', '(', ')', '(', '(', ')', ')', '(', ')', '(', ')', ')', ')', '(', '(', '(', '(', ')', ')', ')', '(', ')'},
			{')', ')', ')', '(', ')', '(', '(', ')', ')', '(', '(', ')', ')', ')', '(', '(', '(', ')', '(', '(', '(', ')', '(', '(', ')', '(', ')', '(', ')', ')', ')', '(', ')', '(', '(', ')', '(', '(', '('},
			{')', ')', '(', ')', '(', ')', '(', '(', '(', ')', ')', '(', '(', ')', ')', '(', ')', ')', ')', ')', ')', '(', ')', ')', ')', ')', '(', '(', ')', '(', ')', '(', ')', '(', '(', '(', '(', ')', ')'},
			{'(', ')', ')', '(', '(', '(', ')', ')', '(', '(', ')', '(', ')', '(', '(', ')', '(', '(', '(', ')', '(', '(', ')', '(', ')', ')', '(', ')', '(', ')', '(', ')', '(', '(', ')', '(', ')', ')', '('},
			{')', '(', ')', '(', ')', '(', '(', '(', ')', ')', ')', '(', ')', ')', '(', '(', '(', ')', '(', '(', '(', '(', '(', '(', ')', ')', ')', ')', '(', ')', ')', ')', ')', '(', '(', ')', '(', ')', '('},
			{'(', ')', ')', ')', ')', ')', '(', '(', ')', '(', '(', '(', '(', '(', '(', '(', '(', ')', ')', '(', '(', '(', '(', '(', ')', '(', '(', ')', ')', '(', ')', '(', ')', ')', ')', ')', ')', ')', ')'},
			{'(', ')', '(', ')', '(', '(', ')', ')', ')', '(', '(', '(', '(', '(', ')', ')', ')', '(', ')', ')', ')', ')', ')', '(', ')', ')', ')', ')', ')', '(', ')', '(', ')', '(', ')', ')', '(', '(', ')'},
			{'(', ')', '(', '(', '(', ')', ')', '(', '(', '(', ')', '(', ')', '(', '(', '(', '(', ')', ')', ')', '(', '(', ')', ')', '(', '(', '(', ')', '(', '(', ')', '(', ')', ')', '(', ')', ')', '(', '('},
			{')', '(', ')', ')', ')', ')', ')', '(', '(', '(', '(', '(', ')', '(', '(', ')', '(', ')', '(', '(', '(', '(', '(', ')', ')', '(', '(', ')', ')', ')', '(', '(', '(', ')', '(', '(', ')', '(', ')'},
			{')', ')', ')', ')', ')', ')', '(', '(', ')', ')', '(', '(', ')', ')', '(', ')', '(', ')', ')', '(', '(', '(', ')', '(', ')', '(', ')', ')', '(', ')', ')', ')', '(', '(', ')', ')', '(', '(', ')'},
			{'(', ')', ')', ')', ')', '(', '(', ')', ')', ')', ')', ')', '(', '(', ')', '(', ')', ')', ')', '(', '(', ')', '(', ')', '(', ')', '(', '(', ')', ')', '(', ')', '(', ')', ')', ')', '(', ')', '('},
			{')', '(', ')', '(', '(', ')', ')', ')', '(', ')', ')', ')', '(', ')', '(', '(', ')', '(', ')', ')', '(', '(', ')', '(', '(', '(', '(', ')', '(', ')', ')', ')', ')', ')', ')', ')', ')', '(', '('},
			{')', '(', ')', '(', '(', ')', '(', ')', '(', '(', ')', '(', ')', '(', ')', '(', '(', '(', ')', '(', ')', '(', ')', '(', ')', '(', ')', ')', ')', '(', ')', '(', '(', '(', '(', ')', ')', '(', '('},
			{')', ')', '(', '(', ')', ')', '(', ')', ')', ')', ')', ')', '(', '(', '(', '(', '(', ')', ')', ')', '(', ')', ')', ')', ')', '(', ')', ')', '(', ')', '(', '(', ')', '(', ')', ')', ')', '(', ')'},
			{'(', '(', ')', ')', ')', ')', ')', ')', '(', '(', '(', '(', ')', '(', '(', ')', ')', '(', '(', ')', '(', '(', ')', '(', '(', '(', '(', ')', ')', '(', '(', ')', '(', ')', '(', ')', '(', '(', ')'},
			{')', '(', ')', '(', ')', '(', ')', ')', ')', ')', '(', '(', ')', ')', '(', ')', '(', '(', '(', '(', '(', ')', ')', '(', ')', ')', '(', ')', ')', ')', ')', '(', '(', '(', ')', '(', '(', '(', ')'},
			{'(', '(', '(', '(', ')', ')', ')', '(', ')', ')', ')', '(', '(', '(', ')', ')', ')', ')', '(', '(', ')', ')', '(', ')', '(', '(', ')', '(', '(', ')', '(', ')', ')', ')', '(', '(', '(', ')', '('},
			{')', '(', '(', '(', '(', '(', ')', '(', ')', '(', ')', '(', ')', ')', ')', ')', ')', '(', '(', '(', '(', ')', ')', '(', ')', '(', '(', ')', ')', ')', ')', ')', ')', '(', ')', '(', '(', ')', '('},
			{'(', '(', ')', '(', '(', '(', ')', '(', ')', ')', '(', ')', '(', ')', '(', ')', ')', ')', ')', '(', '(', ')', ')', '(', ')', '(', ')', ')', '(', '(', '(', ')', ')', '(', '(', ')', '(', '(', ')'},
			{')', ')', ')', '(', ')', ')', ')', ')', '(', '(', '(', ')', ')', ')', '(', '(', '(', '(', ')', '(', ')', ')', ')', ')', ')', '(', ')', '(', '(', ')', '(', '(', ')', ')', '(', '(', '(', ')', ')'},
			{'(', ')', '(', '(', ')', '(', '(', ')', ')', '(', ')', '(', ')', ')', ')', '(', '(', ')', ')', ')', ')', ')', '(', ')', '(', '(', '(', ')', ')', ')', '(', ')', '(', '(', '(', '(', '(', '(', ')'},
			{')', '(', ')', ')', ')', '(', ')', '(', '(', ')', ')', ')', '(', ')', '(', ')', '(', '(', '(', ')', ')', ')', '(', ')', ')', ')', ')', '(', '(', ')', ')', '(', ')', ')', '(', ')', '(', ')', '('},
			{'(', '(', '(', ')', ')', '(', ')', '(', '(', '(', '(', '(', '(', '(', ')', '(', ')', '(', '(', ')', ')', ')', '(', '(', '(', ')', '(', ')', ')', '(', ')', ')', ')', '(', ')', ')', ')', '(', '('},
			{')', '(', ')', '(', ')', '(', '(', ')', '(', ')', ')', '(', ')', ')', ')', ')', '(', ')', '(', ')', ')', ')', ')', ')', '(', ')', '(', ')', '(', ')', ')', ')', ')', '(', '(', ')', ')', ')', ')'},
			{')', ')', '(', ')', '(', ')', '(', '(', ')', '(', ')', ')', ')', '(', '(', '(', '(', '(', '(', ')', '(', ')', ')', '(', ')', ')', ')', '(', ')', '(', ')', '(', ')', ')', '(', ')', ')', '(', ')'},
			{'(', '(', '(', ')', ')', ')', '(', ')', '(', '(', '(', '(', ')', '(', ')', ')', '(', '(', ')', ')', '(', ')', ')', '(', ')', ')', ')', '(', '(', ')', ')', ')', ')', ')', ')', '(', '(', '(', '('},
			{')', '(', '(', '(', '(', ')', '(', ')', ')', ')', ')', ')', ')', ')', ')', ')', '(', ')', ')', ')', '(', ')', '(', ')', '(', ')', ')', '(', ')', '(', ')', ')', '(', '(', '(', '(', '(', '(', '('},
			{')', '(', '(', ')', ')', '(', '(', ')', ')', '(', '(', '(', '(', ')', ')', ')', '(', '(', '(', ')', '(', ')', ')', ')', '(', '(', ')', ')', ')', ')', '(', ')', ')', ')', ')', ')', ')', '(', '('},
			{')', ')', ')', '(', '(', '(', ')', ')', ')', ')', '(', '(', ')', ')', ')', ')', '(', ')', ')', ')', ')', '(', '(', '(', ')', ')', ')', '(', ')', ')', '(', '(', ')', '(', '(', ')', '(', '(', ')'},
			{')', '(', ')', ')', ')', ')', ')', '(', ')', ')', ')', '(', ')', '(', ')', '(', '(', ')', '(', ')', '(', '(', ')', '(', '(', '(', ')', '(', '(', ')', '(', '(', ')', '(', ')', ')', ')', '(', '('},
			{'(', '(', '(', '(', ')', ')', ')', '(', '(', ')', '(', '(', ')', '(', ')', '(', '(', ')', ')', '(', ')', '(', ')', ')', ')', ')', '(', ')', '(', '(', ')', ')', ')', ')', '(', ')', ')', ')', ')'},
			{')', '(', '(', '(', '(', '(', '(', ')', '(', ')', '(', '(', ')', '(', '(', '(', ')', '(', '(', ')', ')', '(', '(', '(', ')', '(', '(', ')', ')', ')', '(', '(', '(', ')', '(', '(', '(', '(', ')'},
		},
		false,
	},
	{
		[][]byte{{')', '('}},
		false,
	},
	{
		[][]byte{{')', '(', '(', ')'}},
		false,
	},
}

func init() {
	var grid = make([][]byte, 99)
	for i := 0; i < 99; i++ {
		grid[i] = make([]byte, 100)
		for l := 0; l < 100; l++ {
			if i > 47 && l >= 50 {
				grid[i][l] = ')'
			} else {
				grid[i][l] = '('
			}
		}
	}
	checkCases = append(checkCases, &checkCase{grid, true})
}

func TestHasValidPath(t *testing.T) {
	for i, tc := range checkCases {
		t.Logf("TestHasValidPath case #%d", i+1)
		result := hasValidPath(tc.in)
		require.Equal(t, tc.out, result)
	}
}