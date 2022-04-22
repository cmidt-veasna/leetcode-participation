package day9

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type digitCase struct {
	s   string
	k   int
	out string
}

var digitCases = []*digitCase{
	{"11111222223", 3, "135"},
	{"00000000", 3, "000"},
	{"1234", 2, "37"},
}

func TestDigitSum(t *testing.T) {
	for i, tc := range digitCases {
		t.Logf("TestDigitSum case #%d", i+1)
		result := digitSum(tc.s, tc.k)
		require.Equal(t, tc.out, result)
	}
}
