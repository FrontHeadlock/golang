package main

import "fmt"

func main() {
	// 배열, 슬라이스, 맵 예제를 순서대로 출력
	printArray()
	fmt.Println()

	printSlice()
	fmt.Println()

	printMap()
	fmt.Println()

	ex()
	fmt.Println()
}

func printArray() {
	fmt.Println("[Arrays]")

	// 길이가 5로 고정된 int 배열
	var numbers [5]int

	// 배열은 인덱스로 접근, 값을 넣지 않은 칸은 0
	numbers[4] = 100
	fmt.Println("numbers =", numbers)
	fmt.Println("numbers[4] =", numbers[4])

	scores := [5]float64{98, 93, 77, 82, 83}
	var total float64

	for _, score := range scores {
		total += score
	}

	fmt.Println("scores length =", len(scores))
	fmt.Println("average =", total/float64(len(scores)))
}

func printSlice() {
	fmt.Println("[Slice]")

	slice1 := []int{1, 2, 3}
	// append는 기존 슬라이스 뒤에 값을 붙인 새 결과를 반환
	slice2 := append(slice1, 4, 5)
	fmt.Println("append slice =", slice1, slice2)

	slice3 := []int{1, 2, 3}
	// make([]int, 2)는 int의 0으로 값을 채움
	slice4 := make([]int, 2)

	// copy 순서: copy(dst, src)이다.
	// 현재 코드는 slice4([0 0])의 값을 slice3으로 복사 -> slice3([0 0 3])
	copy(slice3, slice4)
	fmt.Println("copy slice =", slice3, slice4)
}

func printMap() {
	fmt.Println("[Map]")

	// map은 key를 통해 value를 찾는 자료구조이며, 사용 전 make로 초기화
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x)

	// key 타입과 value 타입은 다르게 지정할 수 있다.
	y := make(map[int]int)
	y[1] = 10
	fmt.Println(y[1])

	// map 안에 또 다른 map을 넣는 중첩 map 예제
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}

	// key가 존재하는지 확인할 때는 value와 ok를 함께 받는 패턴을 자주 사용
	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}

// 연습 문제
func ex() {

	//x := [6]string{"a","b","c","d","e","f"}의 x[2:5] 찾기
	x := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(x[2:5])

	//다음 리스트에서 가장 작은 숫자를 찾는 프로그램을 작성
	k := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	small := k[0]
	// 두 번째 값부터 끝까지 확인하면서 더 작은 값이 나오면 갱신
	for _, value := range k[1:] {
		if value < small {
			small = value
		}
	}

	fmt.Println(small)
}
