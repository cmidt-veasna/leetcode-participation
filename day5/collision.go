package day5

func countCollisions(directions string) int {
	collision := 0
	queue := 0
	if len(directions) > 1 {
		first := directions[0]
		for i := 1; i < len(directions); i++ {
			switch {
			case first == 'R' && directions[i] == 'S':
				first = 'S'
				fallthrough
			case first == 'S' && directions[i] == 'L':
				collision += 1 + queue
				queue = 0
			case first == 'R' && directions[i] == 'L':
				collision += 2 + queue
				queue = 0
				first = 'S'
			case first == 'R' && directions[i] == 'R':
				queue++
			default:
				first = directions[i]
			}
		}
	}
	return collision
}
