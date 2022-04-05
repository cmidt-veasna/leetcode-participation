package day7

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type convertTimeCase struct {
	current string
	correct string
	out     int
}

var convertCases = []*convertTimeCase{
	{"02:30", "04:35", 3},
	{"11:00", "11:01", 1},
	{"04:45", "05:40", 5},
	{"04:42", "05:40", 8},
}

func TestConvertTime(t *testing.T) {
	for i, tc := range convertCases {
		t.Logf("TestConvertTime case #%d", i+1)
		result := convertTime(tc.current, tc.correct)
		require.Equal(t, tc.out, result)
	}
}
