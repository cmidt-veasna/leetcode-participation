package day11

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type appealCase struct {
	in  string
	out int64
}

var appealCases = []*appealCase{
	{"actnnynssgh", 209},
	{"abbca", 28},
	{"code", 20},
	{"fxfz", 18},
}

func init() {
	a := make([]byte, 100000)
	for i, alb := 0, byte('a'); i < len(a); i++ {
		a[i] = alb
		if alb++; alb > 'z' {
			alb = 'a'
		}
	}
	appealCases = append(appealCases, &appealCase{string(a), 129968802600})
}

func TestAppealSum(t *testing.T) {
	for i, tc := range appealCases {
		t.Logf("TestAppealSum case #%d", i+1)
		result := appealSum(tc.in)
		require.Equal(t, tc.out, result)
	}
}
