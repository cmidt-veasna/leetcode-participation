package day7

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type edCase struct {
	key    []byte
	values []string
	dict   []string
	encIn  []string
	decIn  []string
	encOut []string
	decOut []int
}

var edCases = []*edCase{
	{
		[]byte{'a', 'b', 'c', 'd'},
		[]string{"ei", "zf", "ei", "am"},
		[]string{"abcd", "acbd", "adbc", "badc", "dacb", "cadb", "cbda", "abad"},
		[]string{"abcd"},
		[]string{"eizfeiam"},
		[]string{"eizfeiam"},
		[]int{2},
	},
	{
		[]byte{'a', 'b', 'c', 'z'},
		[]string{"aa", "bb", "cc", "zz"},
		[]string{"aa", "aaa", "aaaa", "aaaaa", "aaaaaaa"},
		nil,
		[]string{"aaaa", "aa", "aaaa", "aaaaaaaa", "aaaaaaaaaaaaaa", "aefagafvabfgshdthn"},
		nil,
		[]int{1, 0, 1, 1, 1, 0},
	},
	{
		[]byte{'a', 'b', 'c', 'z'},
		[]string{"aa", "bb", "cc", "zz"},
		[]string{"ab", "ba", "a", "zz", "zzaa", "zabc"},
		[]string{"aabbczz", "zzzzaaaa"},
		[]string{"aabbcczz", "aa", "zzaabbcc", "asdfghjklm"},
		[]string{"aaaabbbbcczzzz", "zzzzzzzzaaaaaaaa"},
		[]int{0, 1, 1, 0},
	},
	{
		[]byte{'a', 'b', 'c', 'd'},
		[]string{"ei","zf","ei","am"},
		[]string{"abcd","acbd","adbc","badc","dacb","cadb","cbda","abad"},
		[]string{"abcd"},
		[]string{"eizfeiam"},
		[]string{"eizfeiam"},
		[]int{2},
	},
}

func TestEncryptDecrypt(t *testing.T) {
	for i, tc := range edCases {
		t.Logf("TestEncryptDecrypt case #%d", i+1)
		encrypter := Constructor(tc.key, tc.values, tc.dict)
		if len(tc.encIn) > 0 {
			result := make([]string, 0, 100)
			for _, in := range tc.encIn {
				result = append(result, encrypter.Encrypt(in))
			}
			require.Equal(t, tc.encOut, result)
		}
		if len(tc.decIn) > 0 {
			result := make([]int, 0, 100)
			for _, in := range tc.decIn {
				result = append(result, encrypter.Decrypt(in))
			}
			require.Equal(t, tc.decOut, result)
		}
	}
}
