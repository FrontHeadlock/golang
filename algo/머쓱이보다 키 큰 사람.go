package algo

func solution(array []int, height int) int {
	// 동적 할당
	count := 0

	//  v 값 array 삽입
	for _, v := range array {
		if v > height {
			count++
		}
	}

	return count
}
