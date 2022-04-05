package day7

import (
	"sort"
	"strings"
)

type Encrypter struct {
	keys   [255]string
	values map[string][]byte
	dict   []string
}

func Constructor(keys []byte, values []string, dictionary []string) Encrypter {
	var k [255]string
	for i, c := range keys {
		k[c] = values[i]
	}
	v := make(map[string][]byte)
	for i, w := range values {
		if r, ok := v[w]; ok {
			v[w] = append(r, keys[i])
		} else {
			v[w] = make([]byte, 1)
			v[w][0] = keys[i]
		}
	}
	sort.Strings(dictionary)
	return Encrypter{k, v, dictionary}
}

func (e *Encrypter) Encrypt(word1 string) string {
	sb := &strings.Builder{}
	for i := 0; i < len(word1); i++ {
		if e.keys[word1[i]] != "" {
			sb.WriteString(e.keys[word1[i]])
		} else {
			sb.WriteByte(word1[i])
		}
	}
	return sb.String()
}

func take(a, b []string, sc string, i, j int) ([]string, bool) {
	found := false
	for _, c := range a {
		if j <= len(c) && c[i:j] == sc {
			b = append(b, c)
			found = true
		}
	}
	return b, found
}

func takec(a, b []string, sc byte, i int) ([]string, bool) {
	found := false
	for _, c := range a {
		if i < len(c) && c[i] == sc {
			b = append(b, c)
			found = true
		}
	}
	return b, found
}

func (e *Encrypter) Decrypt(word2 string) int {
	s1 := e.dict
	s2 := make([]string, 0, len(e.dict))
	sblen := 0
	for i, j := 0, 2; j <= len(word2); i, j = j, j+2 {
		if vv, ok := e.values[word2[i:j]]; ok {
			afound, found := false, false
			for _, c := range vv {
				if s2, found = takec(s1, s2, c, sblen); !afound && found {
					afound = true
				}
			}
			if !afound {
				s1 = nil
				goto done
			}
			sblen++
		} else {
			found := false
			n := sblen + 2
			if s2, found = take(s1, s2, word2[i:j], sblen, n); !found {
				s1 = nil
				goto done
			}
			sblen = n
		}
		s1 = s2
		s2 = make([]string, 0, len(e.dict))
	}
done:
	count := 0
	for _, c := range s1 {
		if len(c) == sblen {
			count++
		}
	}
	return count
}
