package day12

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type countCase struct {
	in  string
	out int
}

var countCases = []*countCase{
	{"22233", 8},
	{"222222222222222222222222222222222222", 82876089},
}

func TestCountTexts(t *testing.T) {
	for i, tc := range countCases {
		t.Logf("TestCountTexts case #%d", i+1)
		result := countTexts(tc.in)
		require.Equal(t, tc.out, result)
	}
}
