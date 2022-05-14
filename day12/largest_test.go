package day12

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var largestCases = [][2]string{
	{"6777133339", "777"},
	{"2300019", "000"},
	{"42352338", ""},
}

func TestLargestGoodInteger(t *testing.T) {
	for i, tc := range largestCases {
		t.Logf("TestLargestGoodInteger case #%d", i+1)
		result := largestGoodInteger(tc[0])
		require.Equal(t, tc[1], result)
	}
}
