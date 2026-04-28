package algo

func solution1000(message string, spoiler_ranges [][]int) int {
	type Word struct {
		start int
		end   int
		text  string
	}

	// 1. 단어 범위 파싱
	words := make([]Word, 0)
	n := len(message)

	for i := 0; i < n; {
		if message[i] == ' ' {
			i++
			continue
		}

		start := i
		for i < n && message[i] != ' ' {
			i++
		}
		end := i - 1

		words = append(words, Word{
			start: start,
			end:   end,
			text:  message[start : end+1],
		})
	}

	m := len(spoiler_ranges)

	// revealEvents[i] = i번째 스포 구간 클릭 후 완전히 공개되는 단어들
	revealEvents := make([][]int, m)

	// 일반 구간에 완전히 노출된 단어
	normalWords := make(map[string]bool)

	rangeIdx := 0

	for wordIdx, word := range words {
		// 현재 단어보다 앞에서 끝난 스포 구간은 제외
		for rangeIdx < m && spoiler_ranges[rangeIdx][1] < word.start {
			rangeIdx++
		}

		revealAt := -1

		// 현재 단어와 겹치는 스포 구간 확인
		for j := rangeIdx; j < m && spoiler_ranges[j][0] <= word.end; j++ {
			if spoiler_ranges[j][1] >= word.start {
				revealAt = j
			}
		}

		if revealAt == -1 {
			// 어떤 스포 구간과도 겹치지 않은 일반 단어
			normalWords[word.text] = true
		} else {
			// 이 단어는 revealAt번째 구간을 클릭했을 때 완전히 공개됨
			revealEvents[revealAt] = append(revealEvents[revealAt], wordIdx)
		}
	}

	answer := 0

	// 이미 공개된 스포 방지 단어
	revealedSpoilerWords := make(map[string]bool)

	// 2. 스포 구간을 왼쪽부터 클릭
	for i := 0; i < m; i++ {
		for _, wordIdx := range revealEvents[i] {
			word := words[wordIdx].text

			if !normalWords[word] && !revealedSpoilerWords[word] {
				answer++
			}

			// 중요 단어 여부와 무관하게, 공개된 스포 단어로 기록
			revealedSpoilerWords[word] = true
		}
	}

	return answer
}
