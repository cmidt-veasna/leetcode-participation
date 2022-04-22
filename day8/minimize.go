package day8

import (
	"fmt"
	"strconv"
	"strings"
)

func minimizeResult(expression string) string {
	s := strings.Split(expression, "+")
	a, _ := strconv.ParseInt(s[0], 10, 32)
	b, _ := strconv.ParseInt(s[1], 10, 32)
	result := fmt.Sprintf("(%s)", expression)
	if a != b {
		smallest := a + b
		as, bs := len(s[0]), len(s[1])
		for i, ix := int64(10), 0; ix < as; i, ix = i*10, ix+1 {
			a1, a2 := a/i, a%i
			for j, jx := int64(1), 0; jx < bs; j, jx = j*10, jx+1 {
				b1, b2 := b/j, b%j
				n := (a2 + b1)
				if a1 > 0 {
					n *= a1
				}
				if b2 > 0 {
					n *= b2
				}
				if smallest > n {
					smallest = n
					switch {
					case a1 == 0 && b2 == 0:
						result = fmt.Sprintf("%d(%d+%d)", a1, a2, b1)
					case a1 == 0:
						result = fmt.Sprintf("(%d+%d)%d", a2, b1, b2)
					case b2 == 0:
						result = fmt.Sprintf("%d(%d+%d)", a1, a2, b1)
					default:
						result = fmt.Sprintf("%d(%d+%d)%d", a1, a2, b1, b2)
					}
				}
			}
		}
	}
	return result
}
