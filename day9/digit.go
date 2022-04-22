package day9

import (
	"fmt"
)

func digitSum(s string, k int) string {
	for len(s) > k {
		a := ""
		for i := 0; i < len(s); i += k {
			sum := 0
			for j, m := i, i+k; j < m && j < len(s); j++ {
				sum += int(s[j] - '0')
			}
			a += fmt.Sprintf("%d", sum)
		}
		s = a
	}
	return s
}
