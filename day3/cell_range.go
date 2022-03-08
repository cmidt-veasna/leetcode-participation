package day3

import (
	"fmt"
	"strconv"
)

func cellsInRange(s string) []string {
	l1 := s[0]
	ln1, _ := strconv.ParseInt(s[1:2], 10, 32)
	l2 := s[3]
	ln2, _ := strconv.ParseInt(s[4:5], 10, 32)
	result := make([]string, 0, int(l2-l1+1)*int(ln2-ln1+1))
	for i := l1; i <= l2; i++ {
		for j := ln1; j <= ln2; j++ {
			result = append(result, fmt.Sprintf("%c%d", i, j))
		}
	}
	return result
}
