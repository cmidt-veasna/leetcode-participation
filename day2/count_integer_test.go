package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountInteger(t *testing.T) {
	var testCases = [][]int{
		// {2, 1},
		// {4, 2}, 
		// {5, 2},
		// {30, 14},
		{896, 447},
	}
	for i, tc := range testCases {
		t.Logf("TestCountInteger case #%d :: %d", i+1, tc[0])
		result := countInteger(tc[0])
		assert.Equal(t, tc[1], result)
	}
}
