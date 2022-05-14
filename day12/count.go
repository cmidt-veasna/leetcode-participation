package day12

var keys = []int{
	2: 3,
	3: 3,
	4: 3,
	5: 3,
	6: 3,
	7: 4,
	8: 3,
	9: 4,
}

const mo = 1000000007

func countTexts(pressedKeys string) int {
	log := make([]int, len(pressedKeys)+1)
	log[0] = 1
	for i, c := range pressedKeys {
		t := 0
		j, k := i, 0
		for j >= 0 && byte(c) == pressedKeys[j] && k < keys[c-'0'] {
			t += log[j]
			t %= mo
			j, k = j-1, k+1
		}
		log[i+1] = t
	}
	return log[len(log)-1]
}
