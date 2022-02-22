package day2


func countInteger(num int) int {
	count := 0
	for i := 2; i <= num; i++ {
		if i < 10 {
			if i%2 == 0 {
				count++
			}
			continue
		}
		sum := 0
		for k := i; true; {
			sum += (k % 10)
			k /= 10
			if k < 10 {
				sum += k
				break
			}
		}
		if sum%2 == 0 {
			count++
		}
	}
	return count
}
