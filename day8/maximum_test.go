package day8

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var maximumCases = [][][]int{
	{{0, 4}, {5, 20}},
	{{6, 3, 3, 2}, {2, 216}},
	{{9, 7, 8}, {9, 1331}},
	{{92, 36, 15, 84, 57, 60, 72, 86, 70, 43, 16}, {1, 254290239}},
}

func TestMaximumProduct(t *testing.T) {
	for i, tc := range maximumCases {
		t.Logf("TestMaximumProduct case #%d", i+1)
		result := maximumProduct(tc[0], tc[1][0])
		require.Equal(t, tc[1][1], result)
	}
}
