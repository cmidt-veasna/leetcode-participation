package day8

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var largestCases = [][2]int{
	{1234, 3412},
	{65875, 87655},
	{2011, 2011},
	{51026, 51620},
	{247, 427},
	{197102030, 973120010},
	{303, 303},
	{1000000000, 1000000000},
}

func TestLargestInteger(t *testing.T) {
	for i, tc := range largestCases {
		t.Logf("TestLargestInteger case #%d", i+1)
		result := largestInteger(tc[0])
		require.Equal(t, tc[1], result, "input %d", tc[0])
	}
}
