package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCellRange = [][][]string{
	{[]string{"A1:A1"}, []string{"A1"}},
	{[]string{"K1:L2"}, []string{"K1", "K2", "L1", "L2"}},
	{[]string{"A1:F1"}, []string{"A1", "B1", "C1", "D1", "E1", "F1"}},
}

func TestCellRange(t *testing.T) {
	for i, tc := range testCellRange {
		t.Logf("TestCellRange case #%d", i+1)
		assert.Equal(t, tc[1], cellsInRange(tc[0][0]))
	}
}
