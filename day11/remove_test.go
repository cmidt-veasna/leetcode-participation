package day11

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type rmDg struct {
	number string
	digit  byte
	out    string
}

var rmCases = []*rmDg{
	{"123", '3', "12"},
	{"1231", '1', "231"},
	{"551", '5', "51"},
	{"522122", '2', "52212"},
	{"523234", '2', "53234"},
	{"55552", '5', "5552"},
	{"743", '7', "43"},
	{"8331743", '7', "833143"},
	{"9731", '2', "9731"},
	{"15454", '4', "1554"},
	{"5486229336471789895711", '7', "548622933647189895711"},
}

func TestRemoveDigit(t *testing.T) {
	for i, tc := range rmCases {
		t.Logf("TestRemoveDigit case #%d", i+1)
		result := removeDigit(tc.number, tc.digit)
		require.Equal(t, tc.out, result)
	}
}
