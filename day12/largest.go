package day12

func largestGoodInteger(num string) string {
	result := ""
	i, j := 0, 3
	for j <= len(num) {
		if num[i] == num[i+1] && num[i] == num[i+2] {
			if len(result) == 0 || result[0] < num[i] {
				result = num[i:j]
			}
			i, j = i+3, j+3
		} else {
			i, j = i+1, j+1
		}
	}
	return result
}
