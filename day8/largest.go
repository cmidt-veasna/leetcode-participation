package day8

func largestInteger(num int) int {
	var b = make([]byte, 10)
	var i = 9
	for ; num > 0; i-- {
		b[i] = byte(num % 10)
		num /= 10
	}
	b = b[i+1:]
	for i := 0; i < len(b); i++ {
		n := b[i]
		if n%2 == 0 {
			for j := i + 1; j < len(b); j++ {
				if b[j]%2 == 0 && n < b[j] {
					b[i], b[j] = b[j], b[i]
					n = b[i]
				}
			}
		} else {
			for j := i + 1; j < len(b); j++ {
				if b[j]%2 != 0 && n < b[j] {
					b[i], b[j] = b[j], b[i]
					n = b[i]
				}
			}
		}
	}
	result := 0
	for i, k := len(b)-1, 1; i >= 0; i, k = i-1, k*10 {
		result += int(b[i]) * k
	}
	return result
}
