package day4

func digArtifacts(n int, artifacts [][]int, dig [][]int) int {
	var a = make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
	}
	for _, info := range dig {
		a[info[0]][info[1]] = 1
	}
	count := 0
nextArtifact:
	for _, info := range artifacts {
		for i := info[0]; i <= info[2]; i++ {
			for j := info[1]; j <= info[3]; j++ {
				if a[i][j] != 1 {
					continue nextArtifact
				}
			}
		}
		count++
	}
	return count
}
