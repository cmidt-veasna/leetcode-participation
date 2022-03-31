package day5

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type collisionCaseData struct {
	s string
	o int
}

var collisionCase = []*collisionCaseData{
	{"RLRSLL", 5},
	{"LLRR", 0},
	{"LLRLRLLSLRLLSLSSSS", 10},
	{"SSRSSRLLRSLLRSRSSRLRRRRLLRRLSSRR", 20},
	{"SRRLRLRSRLRSSRRLSLRLLRSLSLLSSRRLSRSLSLRRS", 28},
}

func TestCountCollisions(t *testing.T) {
	for i, tc := range collisionCase {
		t.Logf("TestCountCollisions case #%d", i+1)
		result := countCollisions(tc.s)
		require.Equal(t, tc.o, result)
	}
}
