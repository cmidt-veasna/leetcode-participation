package day2

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func readTestCase(file string) [][][]int {
	var testCases = [][][]int{}
	buf, _ := ioutil.ReadFile(file)
	vbuf := bufio.NewReader(bytes.NewReader(buf))
	lineBuf := &bytes.Buffer{}
readTestCase:
	for {
		for {
			line, isPrefix, err := vbuf.ReadLine()
			if err == io.EOF {
				break readTestCase
			} else if len(line) > 0 {
				lineBuf.Write(line)
				if !isPrefix {
					break
				}
			}
		}
		str := strings.TrimSpace(lineBuf.String())
		ss := strings.Split(str, ",")
		c, _ := strconv.ParseInt(ss[len(ss)-1], 10, 64)
		k, _ := strconv.ParseInt(ss[len(ss)-2], 10, 64)
		ss = ss[:len(ss)-2]
		is := make([]int, 0, len(ss))
		for _, s := range ss {
			i, _ := strconv.ParseInt(s, 10, 64)
			is = append(is, int(i))
		}
		testCases = append(testCases, [][]int{is, {int(k)}, {int(c)}})
		lineBuf.Reset()
	}
	return testCases
}

var testPairFiles = []string{
	"testdata/countpair_small.txt",
	"testdata/countpair_medium.txt",
	"testdata/countpair_large.txt",
}

func TestCountPair(t *testing.T) {
	for fi, f := range testPairFiles {
		testCases := readTestCase(f)
		for i, tc := range testCases {
			t.Logf("TestCountPairB (%d) #%d :: SIZE(%d)", fi+1, i+1, len(tc[0]))
			var result = countPairs(tc[0], tc[1][0])
			require.Equal(t, int64(tc[2][0]), result)
		}
	}
}

func BenchmarkCountPair(b *testing.B) {
	cases := readTestCase(testPairFiles[len(testPairFiles)-1])
	for i := 0; i < b.N; i++ {
		countPairs(cases[0][0], cases[0][1][0])
	}
}
