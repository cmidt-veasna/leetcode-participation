package day8

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type mb struct {
	gardens    []int
	newFlowers int64
	target     int
	full       int
	partial    int
	out        int64
}

var maxBeautyCases = []*mb{
	{[]int{1, 3, 1, 1}, 7, 6, 12, 1, 14},
	{[]int{2, 4, 5, 3}, 10, 5, 2, 6, 30},
	{[]int{18, 16, 10, 10, 5}, 10, 3, 15, 4, 75},
	{[]int{8, 20, 13}, 12, 12, 4, 3, 41},
	{[]int{13}, 18, 15, 9, 2, 28},
	{
		[]int{36131, 31254, 5607, 11553, 82824, 59230, 40967, 69571, 36874, 38700, 81107, 28500, 61796, 54371, 23405, 51780, 75265, 37257, 86314, 32258, 47254, 76690, 18014, 21538, 96700, 15094, 57253, 57073, 7284, 24501, 21412, 69582, 15724, 43829, 81444, 78281, 88953, 6671, 94646, 31037, 89465, 86033, 27431, 30774, 48592, 26067},
		2304903454,
		48476,
		5937,
		15214,
		737765815,
	},
	{[]int{1, 1}, 2, 2, 1, 2, 3},
	{[]int{20, 1, 15, 17, 10, 2, 4, 16, 15, 11}, 2, 20, 10, 2, 14},
	{[]int{5, 19, 1, 1, 6, 10, 18, 12, 20, 10, 11}, 6, 20, 3, 11, 47},
	{[]int{
		22323, 64818, 97718, 14354, 27837, 6347, 43299, 23010, 18590, 12706, 1579, 52401, 23473, 82978, 1012, 2248, 50247, 755, 12672, 57870, 90646, 87848, 71069, 82723, 83385, 66909, 39266, 97768, 62453, 30454, 68978, 88590, 11213, 74010, 65683, 75460, 3118, 98281, 28128, 84992, 52206, 92770, 74240, 33266, 41603, 19267, 36991, 86658, 43834, 84533, 75187, 31502, 61181, 31333, 37324, 13352, 51735, 89812, 56745, 44204, 85482, 70358, 48830, 91989, 82778, 34042, 3293, 63626, 41301, 43763, 39681, 91917, 40165, 57944, 34357, 22395, 26084, 21666, 40781, 28998, 12385, 10720, 66853, 42324, 28327, 30125, 15522, 12223,
	},
		997843, 100000, 4880, 45790, 2080606020,
	},
	{[]int{8, 2}, 24, 18, 6, 3, 54},
	{[]int{20, 1, 15, 17, 10, 2, 4, 16, 15, 11}, 2, 20, 10, 2, 14},
	{[]int{
		89287, 5538, 37141, 72522, 84502, 44451, 24432, 2324, 72779, 57060, 99178, 6, 29440, 53664, 76197, 46742, 17384, 22072, 33067, 66274, 19323, 72943, 12914, 91475, 96826, 84847, 28100, 89590, 34977, 74052, 4813, 24563, 97491, 71687, 8533, 49262, 2265, 10553, 63902, 19647, 27006, 64548, 89892, 64046, 72766, 34623, 17297, 21417, 70630, 93469, 83379, 19483, 93842, 65968, 28401, 1889, 24441, 99401, 37907, 13794, 3640, 95432, 36875, 10200, 95360, 10829, 96763, 15900, 8490, 68972, 52537, 72458, 95269,
	},
		42, 4534, 32415, 11040, 2734140,
	},
}

func TestMaximumBeauty(t *testing.T) {
	for i, tc := range maxBeautyCases {
		t.Logf("TestMaximumBeauty case #%d", i+1)
		result := maximumBeauty(tc.gardens, tc.newFlowers, tc.target, tc.full, tc.partial)
		require.Equal(t, tc.out, result)
	}
}
