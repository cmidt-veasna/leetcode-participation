package day8

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var minimizeCases = [][2]string{
	{"247+38", "2(47+38)"},
	{"12+34", "1(2+3)4"},
	{"56+98", "(56+98)"},
	{"5+22", "(5+2)2"},
}

func TestMinimizeResult(t *testing.T) {
	for i, tc := range minimizeCases {
		t.Logf("TestMinimizeResult case #%d", i+1)
		result := minimizeResult(tc[0])
		require.Equal(t, tc[1], result)
	}
}
